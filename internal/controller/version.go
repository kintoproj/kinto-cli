package controller

import (
	"fmt"
	"github.com/kintohub/kinto-cli/internal/config"
	"github.com/kintohub/kinto-cli/internal/utils"
)

func (c Controller) Version() {
	utils.NoteMessage(fmt.Sprintf("Kinto Command Line Interface (CLI) %s", config.Version))
}
