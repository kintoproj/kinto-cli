package controller

import (
	"github.com/kintohub/kinto-cli/internal/api"
)

type ControllerInterface interface {
	Init(kintoCoreHost string)
	EnvironmentAccess(envId ...string)
	Services(envId ...string)
	ServiceAccess(envId string, blockId string)
	Deploy(envId string, blockName string)
	Environment()
	Version()
	Access()
	Teleport()
	Status()
}

type Controller struct {
	api api.ApiInterface
}

func NewControllerOrDie(api api.ApiInterface) ControllerInterface {
	return &Controller{
		api: api,
	}
}
