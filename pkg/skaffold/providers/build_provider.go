package providers

import "context"

// BuildProvider implements a common interface that custom builders can use to implement a similar API to the docker client.
type BuildProvider interface {
	ImageID(ctx context.Context, ref string) (string, error)
}
