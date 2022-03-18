// Copyright 2021 The PipeCD Authors.
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

package insightcollector

import (
	"context"
	"sort"
	"time"

	"github.com/pipe-cd/pipecd/pkg/datastore"
	"github.com/pipe-cd/pipecd/pkg/model"
)

const limit = 50

// func (c *Collector) collectDeploymentChangeFailureRate(ctx context.Context, ds []*model.Deployment, target time.Time) error {
// 	apps, projects := groupDeployments(ds)

// 	var updateErr error
// 	for id, ds := range apps {
// 		if err := c.updateApplicationChunks(ctx, ds[0].ProjectId, id, ds, model.InsightMetricsKind_CHANGE_FAILURE_RATE, target); err != nil {
// 			c.logger.Error("failed to update application chunks", zap.Error(err))
// 			updateErr = err
// 		}
// 	}
// 	for id, ds := range projects {
// 		if err := c.updateApplicationChunks(ctx, id, ds[0].ApplicationId, ds, model.InsightMetricsKind_CHANGE_FAILURE_RATE, target); err != nil {
// 			c.logger.Error("failed to update application chunks", zap.Error(err))
// 			updateErr = err
// 		}
// 	}

// 	return updateErr
// }

func (c *Collector) collectDevelopmentFrequency(ctx context.Context, ds []*model.Deployment, target time.Time) error {
	dailyDeployments := groupDeployments(ds)

	var updateErr error
	for _, d := range dailyDeployments {

	}

	return updateErr
}

func (c *Collector) findDeploymentsCreatedInRange(ctx context.Context, from, to int64) ([]*model.Deployment, error) {
	filters := []datastore.ListFilter{
		{
			Field:    "CreatedAt",
			Operator: datastore.OperatorGreaterThanOrEqual,
			Value:    from,
		},
	}

	var deployments []*model.Deployment
	maxCreatedAt := to
	for {
		d, _, err := c.deploymentStore.List(ctx, datastore.ListOptions{
			Limit: limit,
			Filters: append(filters, datastore.ListFilter{
				Field:    "CreatedAt",
				Operator: datastore.OperatorLessThan,
				Value:    maxCreatedAt,
			}),
			Orders: []datastore.Order{
				{
					Field:     "CreatedAt",
					Direction: datastore.Desc,
				},
				{
					Field:     "Id",
					Direction: datastore.Asc,
				},
			},
		})
		if err != nil {
			return nil, err
		}
		if len(d) == 0 {
			// get all deployments in range
			break
		}

		deployments = append(deployments, d...)
		maxCreatedAt = d[len(d)-1].CreatedAt
	}
	return deployments, nil
}

// groupDeployments groups deployments by applicationID and projectID
// deployments must be sorted by createdAt ASC.
// Result is sorted by DailyDeployment.Date ASC.
func groupDeployments(deployments []*model.Deployment) []*model.DailyDeployment {
	dailyDeployments := make(map[int64]*model.DailyDeployment)

	for _, d := range deployments {
		t := time.Unix(d.CreatedAt, 0)
		tt := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC).Unix()
		if _, ok := dailyDeployments[tt]; !ok {
			dailyDeployments[tt] = &model.DailyDeployment{
				Date: tt,
				DailyDeployments: []*model.DeploymentSubset{
					{
						CreatedAt: d.CreatedAt,
						UpdatedAt: d.UpdatedAt,
					},
				},
			}
			continue
		}

		dailyDeployments[tt].DailyDeployments = append(dailyDeployments[tt].DailyDeployments, &model.DeploymentSubset{
			CreatedAt: d.CreatedAt,
			UpdatedAt: d.UpdatedAt,
		})
	}

	result := make([]*model.DailyDeployment, len(dailyDeployments))
	index := 0
	for key := range dailyDeployments {
		result[index] = dailyDeployments[key]
		index++
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Date < result[j].Date
	})

	return result
}
