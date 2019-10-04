package main

import (
	"github.com/rfornea/tool-inventory/pkg/backend/models"
	"github.com/rfornea/tool-inventory/pkg/backend/routes"
	"github.com/rfornea/tool-inventory/pkg/backend/utils"
)

func main() {
	utils.SetLive()

	models.Connect(utils.Env.DatabaseURL)

	routes.CreateRoutes()
}
