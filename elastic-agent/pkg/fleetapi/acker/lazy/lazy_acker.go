// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package lazy

import (
	"context"

	"github.com/elastic/elastic-agent-poc/elastic-agent/pkg/core/logger"
	"github.com/elastic/elastic-agent-poc/elastic-agent/pkg/fleetapi"
)

type batchAcker interface {
	AckBatch(ctx context.Context, actions []fleetapi.Action) error
}

type ackForcer interface {
	ForceAck()
}

// Acker is a lazy acker which performs HTTP communication on commit.
type Acker struct {
	log   *logger.Logger
	acker batchAcker
	queue []fleetapi.Action
}

// NewAcker creates a new lazy acker.
func NewAcker(baseAcker batchAcker, log *logger.Logger) *Acker {
	return &Acker{
		acker: baseAcker,
		queue: make([]fleetapi.Action, 0),
		log:   log,
	}
}

// Ack acknowledges action.
func (f *Acker) Ack(ctx context.Context, action fleetapi.Action) error {
	f.enqueue(action)

	if _, isAckForced := action.(ackForcer); isAckForced {
		return f.Commit(ctx)
	}

	return nil
}

// Commit commits ack actions.
func (f *Acker) Commit(ctx context.Context) error {
	err := f.acker.AckBatch(ctx, f.queue)
	if err != nil {
		// do not cleanup on error
		return err
	}

	f.queue = make([]fleetapi.Action, 0)
	return nil
}

func (f *Acker) enqueue(action fleetapi.Action) {
	for _, a := range f.queue {
		if a.ID() == action.ID() {
			f.log.Debugf("action with id '%s' has already been queued", action.ID())
			return
		}
	}
	f.queue = append(f.queue, action)
	f.log.Debugf("appending action with id '%s' to the queue", action.ID())
}