// Copyright 2020 The PipeCD Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ecs

import (
	"context"
	"fmt"
	"io"
	"time"

	"go.uber.org/zap"

	provider "github.com/pipe-cd/pipecd/pkg/app/piped/cloudprovider/ecs"
	"github.com/pipe-cd/pipecd/pkg/app/piped/planner"
	"github.com/pipe-cd/pipecd/pkg/model"
)

// Planner plans the deployment pipeline for ECS application.
type Planner struct {
}

type registerer interface {
	Register(k model.ApplicationKind, p planner.Planner) error
}

// Register registers this planner into the given registerer.
func Register(r registerer) {
	r.Register(model.ApplicationKind_ECS, &Planner{})
}

// Plan decides which pipeline should be used for the given input.
func (p *Planner) Plan(ctx context.Context, in planner.Input) (out planner.Output, err error) {
	ds, err := in.TargetDSP.Get(ctx, io.Discard)
	if err != nil {
		err = fmt.Errorf("error while preparing deploy source data (%v)", err)
		return
	}

	cfg := ds.ApplicationConfig.ECSApplicationSpec
	if cfg == nil {
		err = fmt.Errorf("missing ECSApplicationSpec in application configuration")
		return
	}

	// Determine application version from the task definition
	if version, err := determineVersion(ds.AppDir, cfg.Input.TaskDefinitionFile); err == nil {
		out.Version = version
	} else {
		out.Version = "unknown"
		in.Logger.Warn("unable to determine target version", zap.Error(err))
	}

	out.Versions, err = determineVersions(ds.AppDir, cfg.Input.TaskDefinitionFile)
	if err != nil {
		in.Logger.Warn("unable to determine target versions", zap.Error(err))
		out.Versions = []*model.ArtifactVersion{
			{
				Kind:    model.ArtifactVersion_UNKNOWN,
				Version: "unknown",
			},
		}
	}

	autoRollback := *cfg.Input.AutoRollback

	// In case the strategy has been decided by trigger.
	// For example: user triggered the deployment via web console.
	switch in.Trigger.SyncStrategy {
	case model.SyncStrategy_QUICK_SYNC:
		out.SyncStrategy = model.SyncStrategy_QUICK_SYNC
		out.Stages = buildQuickSyncPipeline(autoRollback, time.Now())
		out.Summary = in.Trigger.StrategySummary
		return
	case model.SyncStrategy_PIPELINE:
		if cfg.Pipeline == nil {
			err = fmt.Errorf("unable to force sync with pipeline because no pipeline was specified")
			return
		}
		out.SyncStrategy = model.SyncStrategy_PIPELINE
		out.Stages = buildProgressivePipeline(cfg.Pipeline, autoRollback, time.Now())
		out.Summary = in.Trigger.StrategySummary
		return
	}

	// When no pipeline was configured, perform the quick sync.
	if cfg.Pipeline == nil || len(cfg.Pipeline.Stages) == 0 {
		out.SyncStrategy = model.SyncStrategy_QUICK_SYNC
		out.Stages = buildQuickSyncPipeline(autoRollback, time.Now())
		out.Summary = fmt.Sprintf("Quick sync to deploy image %s and configure all traffic to it (pipeline was not configured)", out.Version)
		return
	}

	// Force to use pipeline when the alwaysUsePipeline field was configured.
	if cfg.Planner.AlwaysUsePipeline {
		out.SyncStrategy = model.SyncStrategy_PIPELINE
		out.Stages = buildProgressivePipeline(cfg.Pipeline, autoRollback, time.Now())
		out.Summary = "Sync with the specified pipeline (alwaysUsePipeline was set)"
		return
	}

	// If this is the first time to deploy this application or it was unable to retrieve last successful commit,
	// we perform the quick sync strategy.
	if in.MostRecentSuccessfulCommitHash == "" {
		out.SyncStrategy = model.SyncStrategy_QUICK_SYNC
		out.Stages = buildQuickSyncPipeline(autoRollback, time.Now())
		out.Summary = fmt.Sprintf("Quick sync to deploy image %s and configure all traffic to it (it seems this is the first deployment)", out.Version)
		return
	}

	// Load service manifest at the last deployed commit to decide running version.
	ds, err = in.RunningDSP.Get(ctx, io.Discard)
	if err == nil {
		if lastVersion, e := determineVersion(ds.AppDir, cfg.Input.TaskDefinitionFile); e == nil {
			out.SyncStrategy = model.SyncStrategy_PIPELINE
			out.Stages = buildProgressivePipeline(cfg.Pipeline, autoRollback, time.Now())
			out.Summary = fmt.Sprintf("Sync with pipeline to update image from %s to %s", lastVersion, out.Version)
			return
		}
	}

	out.SyncStrategy = model.SyncStrategy_PIPELINE
	out.Stages = buildProgressivePipeline(cfg.Pipeline, autoRollback, time.Now())
	out.Summary = "Sync with the specified pipeline"
	return
}

func determineVersion(appDir, taskDefinitonFile string) (string, error) {
	taskDefinition, err := provider.LoadTaskDefinition(appDir, taskDefinitonFile)
	if err != nil {
		return "", err
	}

	return provider.FindImageTag(taskDefinition)
}

func determineVersions(appDir, taskDefinitonFile string) ([]*model.ArtifactVersion, error) {
	taskDefinition, err := provider.LoadTaskDefinition(appDir, taskDefinitonFile)
	if err != nil {
		return nil, err
	}

	return provider.FindArtifactVersions(taskDefinition)
}
