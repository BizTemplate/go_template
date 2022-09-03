package core

import (
	"errors"
	"viper_contract/src/infrastructure"
)

type Starter interface {
	Start()
	Reload()
	Ping()
}

type StarterType string

const (
	HTTP StarterType = "HTTP"
)

func NewStarter(config infrastructure.Config) (Starter, error) {
	switch config.StarterConfig.StarterType {
	case string(HTTP):
		return &RoutingStarter{}, nil
	default:
		return nil, errors.New("starter Type Not support")
	}
}
