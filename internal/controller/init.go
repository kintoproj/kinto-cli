package controller

import (
	"fmt"
	"github.com/kintoproj/kinto-cli/internal/config"
	"github.com/kintoproj/kinto-cli/internal/utils"
)

//Set kintoCoreHost for CLI or reset it to default production host.
func (c *Controller) Init(kintoCoreHost string) {

	if kintoCoreHost == "Reset" || kintoCoreHost == "reset" {
		config.SetKintoCoreHost(config.DefaultkintoCoreHost)
		config.Save()
		utils.SuccessMessage("CoreHost deleted")
	} else {
		config.SetKintoCoreHost(kintoCoreHost)
		config.Save()
		utils.SuccessMessage(fmt.Sprintf("New CoreHost set as => %s", kintoCoreHost))
	}
}
