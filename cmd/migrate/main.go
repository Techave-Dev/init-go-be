package main

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
	"github.com/techave-dev/init-go-be/tools"
)

func main() {
	cfg, err := tools.LoadConfig()
	if err != nil {
		logrus.Fatal("cannot load config")
	}

	mgr, err := migrate.New(cfg.MigrationPath, cfg.PostgresURL)
	if err != nil {
		logrus.Fatal("failed to run new migration instance")
	}

	if err := mgr.Up(); err != nil && err != migrate.ErrNoChange {
		logrus.Fatal("failed to run migrate up")
	}

	logrus.Info("db migrated successfully")
}
