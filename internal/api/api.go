package api

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/techave-dev/init-go-be/tools"
)

type Server struct {
	fiber *fiber.App
	cfg   *tools.Config
	pool  *pgxpool.Pool
}

func NewServer(cfg *tools.Config, pool *pgxpool.Pool) *Server {
	return &Server{fiber: fiber.New(), cfg: cfg, pool: pool}
}

func (s *Server) Run() error {

	go func() {
		if err := s.fiber.Listen(s.cfg.AppPort); err != nil {
			logrus.Fatal("Error starting Server: ", err.Error())
		}
	}()

	s.MapHandlers(s.fiber)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	_, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return s.fiber.Shutdown()
}
