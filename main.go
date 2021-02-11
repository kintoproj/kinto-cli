package main

import (
	"github.com/kintohub/kinto-cli/internal/api"
	"github.com/kintohub/kinto-cli/internal/cli"
	"github.com/kintohub/kinto-cli/internal/controller"
)

func main() {
	cli := cli.NewCliOrDie()
	api := api.NewApiOrDie(cli.GetHostFlag())
	controller := controller.NewControllerOrDie(api)
	cli.Execute(controller)
}
