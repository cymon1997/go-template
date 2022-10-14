package ping

import "context"

type Pingable interface {
	Name() string
	Ping(ctx context.Context) error
}
