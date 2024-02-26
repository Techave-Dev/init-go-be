package tools

import (
	"fmt"

	"github.com/caarlos0/env/v10"
	"github.com/invopop/validation"
	"github.com/invopop/validation/is"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type AppEnv string

const (
	AppEnvDevelopment AppEnv = "development"
	AppEnvStaging     AppEnv = "staging"
	AppEnvProduction  AppEnv = "production"
)

func (a AppEnv) Valid() bool {
	switch a {
	case AppEnvDevelopment,
		AppEnvStaging,
		AppEnvProduction:
		return true
	}
	return false
}

type Config struct {
	AppEnv        AppEnv `env:"APP_ENV"`
	PostgresURL   string `env:"POSTGRES_URL"`
	MigrationPath string `env:"MIGRATION_PATH"`
	AppPort       string `env:"APP_PORT"`
}

func (c *Config) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.AppEnv, validation.Required, validation.By(func(value any) error {
			var err error = fmt.Errorf("invalid APP_ENV: %s ", c.AppEnv)
			switch enum := value.(type) {
			case AppEnv:
				if valid := enum.Valid(); valid {
					err = nil
				}
			}
			return err
		})),
		validation.Field(&c.PostgresURL, validation.Required, is.URL),
		validation.Field(&c.MigrationPath, validation.Required),
		validation.Field(&c.AppPort, validation.Required),
	)
}

func LoadConfig() (Config, error) {
	const path = ".env"

	if err := godotenv.Load(path); err != nil {
		logrus.Warn(err.Error())
	}

	var config Config
	if err := env.Parse(&config); err != nil {
		return Config{}, err
	}

	if err := config.Validate(); err != nil {
		return Config{}, err
	}

	return config, nil
}
