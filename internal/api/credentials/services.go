package credentials

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/techave-dev/init-go-be/internal/repo"
	"github.com/techave-dev/init-go-be/tools"
	"golang.org/x/crypto/bcrypt"
)

type Services struct {
	config  *tools.Config
	pool    *pgxpool.Pool
	queries repo.Querier
}

func NewServices(config *tools.Config, pool *pgxpool.Pool) Services {
	return Services{
		config,
		pool,
		repo.New(pool),
	}
}

type RegisterParams = repo.InsertCredentialParams

func (s *Services) Register(c context.Context, params RegisterParams) (repo.Credential, error) {
	safePassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return repo.Credential{}, err
	}

	credential, err := s.queries.InsertCredential(c, repo.InsertCredentialParams{
		Email:    params.Email,
		Password: string(safePassword),
		Role:     params.Role,
	})

	return credential, err
}

type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:""`
}

type LoginReturn struct {
	Token      string          `json:"token"`
	Credential repo.Credential `json:"credential"`
}

func (s *Services) Login(c context.Context, params LoginParams) (LoginReturn, error) {
	credential, err := s.queries.FindCredentialByEmail(c, params.Email)
	if err != nil {
		return LoginReturn{}, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(credential.Password), []byte(params.Password)); err != nil {
		return LoginReturn{}, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tools.JwtClaims{
		CredentialID: credential.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(1 * time.Hour),
			},
		},
	})

	tokenString, err := token.SignedString([]byte("handiism"))
	if err != nil {
		return LoginReturn{}, err
	}

	return LoginReturn{
		Token:      tokenString,
		Credential: credential,
	}, nil
}

func (s *Services) CredentialById(c context.Context, id uuid.UUID) (repo.Credential, error) {
	return s.queries.FindCredentialById(c, id)
}
