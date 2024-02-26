package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/techave-dev/init-go-be/internal/api"
	"github.com/techave-dev/init-go-be/tools"
)

func main() {
	var formatter logrus.Formatter = &logrus.TextFormatter{DisableTimestamp: true}

	if env, exist := os.LookupEnv("APP_ENV"); exist && env != string(tools.AppEnvDevelopment) {
		fmt.Println("mashok")
		formatter = &logrus.JSONFormatter{}
	}

	logrus.SetFormatter(formatter)

	config, err := tools.LoadConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	pool, err := pgxpool.New(context.Background(), config.PostgresURL)
	if err != nil {
		logrus.Fatal("cannot connect to db")
	}

	server := api.NewServer(&config, pool)

	if err := server.Run(); err != nil {
		logrus.Fatal(err)
	}
}
