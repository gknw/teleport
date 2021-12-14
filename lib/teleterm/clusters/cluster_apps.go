/*
Copyright 2015 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package clusters

import (
	"context"

	"github.com/gravitational/teleport/api/types"
	"github.com/gravitational/teleport/lib/teleterm/api/uri"

	"github.com/gravitational/trace"
)

// Database describes database
type App struct {
	// URI is the database URI
	URI string `json:"uri"`
	types.Application
}

// GetDatabase returns a database
func (c *Cluster) GetApps(ctx context.Context) ([]App, error) {
	// Get a list of all applications.
	apps, err := c.clusterClient.ListApps(ctx)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	results := []App{}
	for _, app := range apps {
		results = append(results, App{
			URI:         uri.Cluster(c.status.Name).App(app.GetName()).String(),
			Application: app,
		})
	}

	return results, nil
}
