package insightstore

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/pipe-cd/pipecd/pkg/filestore"
	"github.com/pipe-cd/pipecd/pkg/model"
)

var (
	errInvalidArg    = errors.New("invalid arg")
	errUnimplemented = errors.New("unimplemented")
	errNotFound      = errors.New("not found")
)

const (
	maxChunkByteSize = 1 * 1024 * 1024
)

type DeploymentStore interface {
	GetDailyDeployments(ctx context.Context, projectID string, date_rage *model.ChunkDateRange) ([]*model.DailyDeployment, error)

	PutDailyDeployment(ctx context.Context, projectID string, chunk *model.DailyDeployment) error
}

func (s *store) GetDailyDeployments(ctx context.Context, projectID string, date_range *model.ChunkDateRange) ([]*model.DailyDeployment, error) {
	fromYear := time.Unix(date_range.From, 0).Year()
	toYear := time.Unix(date_range.To, 0).Year()

	if fromYear != toYear {
		return nil, errUnimplemented
	}
	if fromYear > toYear {
		return nil, errInvalidArg
	}

	dirPath := determineDeploymentDirPath(fromYear, projectID)
	meta := model.DeploymentChunkMetaData{}
	err := s.loadProtoMessage(ctx, fmt.Sprintf("%s/meta.json", dirPath), &meta)
	if err != nil {
		return nil, err
	}

	keys := findPathFromMeta(&meta, date_range)

	var result []*model.DailyDeployment
	for _, key := range keys {
		rawData, err := s.filestore.Get(ctx, fmt.Sprintf("%s/%s", dirPath, key))
		if err != nil {
			return nil, err
		}

		data := model.DeploymentChunk{}
		err = proto.Unmarshal(rawData, &data)
		if err != nil {
			return nil, err
		}

		dailyDeployments := extractDailyDeploymentFromChunk(&data, date_range)
		result = append(result, dailyDeployments...)
	}
	return result, nil
}

func (s *store) PutDailyDeployment(ctx context.Context, projectID string, dailyDeployments *model.DailyDeployment) error {
	// Load chunk
	year := time.Unix(dailyDeployments.Date, 0).Year()
	dirPath := determineDeploymentDirPath(year, projectID)
	meta := model.DeploymentChunkMetaData{}
	metaPath := fmt.Sprintf("%s/meta.json", dirPath)
	err := s.loadProtoMessage(ctx, metaPath, &meta)
	if err != nil && !errors.Is(err, filestore.ErrNotFound) {
		return err
	}

	// If there is no chunk, simply create new metadata and new chunk
	if len(meta.Data) == 0 {
		// Create chunk
		chunkKey := determineDeploymentChunkKey(1)
		chunkData := model.DeploymentChunk{
			DateRange: &model.ChunkDateRange{
				From: dailyDeployments.Date,
				To:   dailyDeployments.Date,
			},
			Deployments: []*model.DailyDeployment{dailyDeployments},
		}
		size, err := s.storeProtoMessage(ctx, fmt.Sprintf("%s/%s", dirPath, chunkKey), &chunkData)
		if err != nil {
			return err
		}

		// Create meta
		newMetaData := model.DeploymentChunkMetaData{
			Data: []*model.DeploymentChunkMetaData_ChunkData{
				{
					DateRange: &model.ChunkDateRange{
						From: dailyDeployments.Date,
						To:   dailyDeployments.Date,
					},
					ChunkKey:  chunkKey,
					ChunkSize: size,
				},
			},
		}

		_, err = s.storeProtoMessage(ctx, metaPath, &newMetaData)
		return err
	}

	latestChunkMeta := meta.Data[len(meta.Data)-1]
	if dailyDeployments.Date < latestChunkMeta.DateRange.To {
		return errors.New("cannot overwrite past deployment")
	}

	// To know the size, we marshal here instead of using storeProtoMessage.
	rawData, err := proto.Marshal(dailyDeployments)
	if err != nil {
		return err
	}

	// append to current chunk
	if latestChunkMeta.ChunkSize+int64(len(rawData)) < maxChunkByteSize {
		// Update chunk
		chunkData := model.DeploymentChunk{}
		chunkPath := fmt.Sprintf("%s/%s", dirPath, latestChunkMeta.ChunkKey)
		err := s.loadProtoMessage(ctx, chunkPath, &chunkData)
		if err != nil {
			return err
		}

		if dailyDeployments.Date < chunkData.DateRange.To {
			panic("should not come here")
		}

		chunkData.Deployments = append(chunkData.Deployments, dailyDeployments)
		chunkData.DateRange.To = dailyDeployments.Date

		size, err := s.storeProtoMessage(ctx, chunkPath, &chunkData)
		if err != nil {
			return err
		}

		// Update meta
		latestChunkMeta.ChunkSize = size
		latestChunkMeta.DateRange.To = dailyDeployments.Date
		_, err = s.storeProtoMessage(ctx, chunkPath, &meta)
		if err != nil {
			return err
		}
	}

	// Create new meta and chunk
	newChunkKey := determineDeploymentChunkKey(len(meta.Data) + 1)
	chunkData := model.DeploymentChunk{
		DateRange: &model.ChunkDateRange{
			From: dailyDeployments.Date,
			To:   dailyDeployments.Date,
		},
		Deployments: []*model.DailyDeployment{dailyDeployments},
	}
	size, err := s.storeProtoMessage(ctx, fmt.Sprintf("%s/%s", dirPath, newChunkKey), &chunkData)
	if err != nil {
		return err
	}

	// Create meta
	newChunkMetaData := model.DeploymentChunkMetaData_ChunkData{
		DateRange: &model.ChunkDateRange{
			From: dailyDeployments.Date,
			To:   dailyDeployments.Date,
		},
		ChunkKey:  newChunkKey,
		ChunkSize: size,
	}
	meta.Data = append(meta.Data, &newChunkMetaData)

	_, err = s.storeProtoMessage(ctx, metaPath, &meta)
	return err
}

// Load proto message stored in path. If path does not exists, dest does not modified.
func (s *store) loadProtoMessage(ctx context.Context, path string, dest proto.Message) error {
	raw, err := s.filestore.Get(ctx, path)
	if err != nil {
		return err
	}

	err = proto.Unmarshal(raw, dest)
	if err != nil {
		return err
	}
	return nil
}

func (s *store) storeProtoMessage(ctx context.Context, path string, data proto.Message) (dataSize int64, err error) {
	raw, err := proto.Marshal(data)
	if err != nil {
		return dataSize, err
	}

	err = s.filestore.Put(ctx, path, raw)
	if err != nil {
		return dataSize, err
	}
	return int64(len(raw)), nil
}

func findPathFromMeta(meta *model.DeploymentChunkMetaData, date_range *model.ChunkDateRange) []string {
	var paths []string
	for _, m := range meta.Data {
		if overlap(date_range, m.DateRange) {
			paths = append(paths, m.ChunkKey)
		}
	}
	return paths
}

func overlap(lhs *model.ChunkDateRange, rhs *model.ChunkDateRange) bool {
	return (rhs.From < lhs.To && lhs.From < rhs.To)
}

func extractDailyDeploymentFromChunk(chunk *model.DeploymentChunk, date_range *model.ChunkDateRange) []*model.DailyDeployment {
	var result []*model.DailyDeployment
	for _, d := range chunk.Deployments {
		if date_range.From < d.Date && d.Date < date_range.To {
			result = append(result, d)
		}
	}
	return result
}

func determineDeploymentDirPath(year int, projectID string) string {
	return fmt.Sprintf("insights/deployments/%s/%d", projectID, year)
}

func determineDeploymentChunkKey(n int) string {
	return fmt.Sprintf("%04d.proto.bin", n)
}
