package deploymanager

import (
	"context"

	"github.com/gavril-s/borzoi/internal/pkg/domain"
)

type Storage interface {
	UpsertDeploy(ctx context.Context, deploy domain.Deploy) error
}
