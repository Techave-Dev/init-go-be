package middlewares

import (
	"slices"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/techave-dev/init-go-be/internal/repo/psql"
	"github.com/techave-dev/init-go-be/tools"
)

func (m *MiddlewaresManager) Verify(ability psql.AbilityEnum) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if ability == psql.AbilityEnumPublic {
			return c.Next()
		}

		headers := c.GetReqHeaders()
		auths := headers["Authorization"]
		query := c.Query("token", "")
		tokenString := ""
		if len(auths) > 0 {
			splitedAuths := strings.Split(auths[0], " ")
			if len(splitedAuths) == 2 {
				tokenString = splitedAuths[1]
			}
		} else if query != "" {
			tokenString = query
		}

		if tokenString == "" {
			return tools.Fail(c, 400, tools.String("Permintaan gagal, tidak ada token permintaan"))
		}

		token, err := jwt.ParseWithClaims(tokenString, &tools.JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(m.config.JwtKey), nil
		})
		if err != nil {
			return tools.Fail(c, 400, tools.String("Format token tidak valid"))
		}

		claims, ok := token.Claims.(*tools.JwtClaims)
		if !ok {
			return tools.Fail(c, 400, tools.String("Format token tidak valid"))
		}

		if ability == psql.AbilityEnumPrivate {
			c.Locals("CredentialID", claims.CredentialID.String())
			return c.Next()
		}

		abilities, err := m.services.FindCredentialAbilities(c.Context(), claims.CredentialID)
		if err != nil {
			return tools.Fail(c, 404, tools.String("Akun tidak dapat ditemukan"))
		}

		if authorized := slices.Contains(abilities, ability); !authorized {
			return tools.Fail(c, 401, tools.String("Akun tidak dapat mengakses endpoint yang dituju"))
		}

		c.Locals("CredentialID", claims.CredentialID.String())
		return c.Next()
	}
}
