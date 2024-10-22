package portmanager

import (
	"context"
)

type Storage interface {
	IsPortBusy(ctx context.Context, port int) (bool, error)
	MarkPortAsBusy(ctx context.Context, port int) error
}
