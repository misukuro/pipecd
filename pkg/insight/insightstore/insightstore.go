package insightstore

import "github.com/pipe-cd/pipecd/pkg/filestore"

type Store interface {
	ApplicationCountStore
	DeploymentStore
	MileStoneStore
}

type store struct {
	filestore filestore.Store
}

func NewStore(fs filestore.Store) Store {
	return &store{
		filestore: fs,
	}
}
