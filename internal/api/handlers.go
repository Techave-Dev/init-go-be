package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/techave-dev/init-go-be/internal/api/credentials"
	"github.com/techave-dev/init-go-be/internal/api/middlewares"
	"github.com/techave-dev/init-go-be/internal/repo"
)

func (s *Server) MapHandlers(f *fiber.App) {
	// init services
	credServices := credentials.NewServices(s.cfg, s.pool)
	midServices := middlewares.NewServices(s.cfg, s.pool)

	// init middlewares
	mid := middlewares.NewMiddlewaresManager(s.cfg, midServices)

	// init handlers
	credHandlers := credentials.NewHandlers(s.cfg, credServices)

	credentials := f.Group("/credentials")
	credentials.Post("/login", credHandlers.Login())
	credentials.Post("/register", credHandlers.Register())
	credentials.Post("/me", mid.Verify(repo.AbilityEnumPrivate), credHandlers.Me())
}
