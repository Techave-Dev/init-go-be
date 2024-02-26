package middlewares

import (
	"github.com/techave-dev/init-go-be/tools"
)

type MiddlewaresManager struct {
	config   *tools.Config
	services Services
}

func NewMiddlewaresManager(config *tools.Config, services Services) MiddlewaresManager {
	return MiddlewaresManager{
		config, services,
	}
}
