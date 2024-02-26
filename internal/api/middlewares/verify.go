package middlewares

import (
	"fmt"
	"slices"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/techave-dev/init-go-be/internal/repo"
	"github.com/techave-dev/init-go-be/tools"
)

func (m *MiddlewaresManager) Verify(ability repo.AbilityEnum) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if ability == repo.AbilityEnumPublic {
			return c.Next()
		}

		headers := c.GetReqHeaders()
		auths := headers["Authorization"]
		tokenString := ""
		if len(auths) > 0 {
			splitedAuths := strings.Split(auths[0], " ")
			fmt.Printf("splitedAuths: %v\n", splitedAuths)
		}

		token, err := jwt.ParseWithClaims(tokenString, &tools.JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte("handiism"), nil
		})
		if err != nil {
			return c.Status(400).JSON(map[string]any{"error": err.Error()})
		}
		claims, ok := token.Claims.(*tools.JwtClaims)
		if !ok {
			return c.Status(400).JSON(map[string]any{"error": err.Error()})
		}

		if ability == repo.AbilityEnumPrivate {
			c.Locals("CredentialID", claims.CredentialID.String())
			return c.Next()
		}

		abilities, err := m.services.FindCredentialAbilities(c.Context(), claims.CredentialID)
		if err != nil {
			return c.Status(400).JSON(map[string]any{"error": err.Error()})
		}

		if authorized := slices.Contains(abilities, ability); !authorized {
			return c.Status(400).JSON(map[string]any{"error": "ability did not permitted"})
		}

		c.Locals("CredentialID", claims.CredentialID.String())
		return c.Next()
	}
}
