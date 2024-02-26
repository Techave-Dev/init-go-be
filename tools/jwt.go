package tools

import (
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	jwt.RegisteredClaims
	CredentialID uuid.UUID `json:"credentialId"`
}
