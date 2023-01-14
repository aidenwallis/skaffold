/*
Copyright 2019 The Skaffold Authors

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

package finch

import (
	"context"

	"github.com/GoogleContainerTools/skaffold/v2/pkg/skaffold/providers"
)

// LocalDaemon is an interface for interacting with the local finch daemon on the host.
//
// Finch does not seem to have a published API spec out right now, or engine API, so this uses
// CLI commands to interact with the daemon.
//
// See: https://github.com/runfinch/finch
type LocalDaemon struct {
}

// Ensure that localDaemon implements LocalDaemon.
var _ providers.BuildProvider = (*LocalDaemon)(nil)

// NewLocalDaemon returns a new LocalDaemon.
func NewLocalDaemon() *LocalDaemon {
	return &LocalDaemon{}
}

// ImageID returns the image ID for the given image reference.
func (d *LocalDaemon) ImageID(ctx context.Context, ref string) (string, error) {
	return "", nil
}
