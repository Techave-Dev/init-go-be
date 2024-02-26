package middlewares

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/techave-dev/init-go-be/internal/repo"
	"github.com/techave-dev/init-go-be/tools"
)

type Services struct {
	config  *tools.Config
	pool    *pgxpool.Pool
	queries repo.Querier
}

func NewServices(config *tools.Config, pool *pgxpool.Pool) Services {
	return Services{
		config:  config,
		pool:    pool,
		queries: repo.New(pool),
	}
}

func (s *Services) FindCredentialAbilities(c context.Context, id uuid.UUID) ([]repo.AbilityEnum, error) {
	return s.queries.FindCredentialAbilities(c, id)
}
