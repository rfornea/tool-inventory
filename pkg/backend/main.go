package main

import (
	"github.com/rfornea/library/pkg/backend/models"
	"github.com/rfornea/library/pkg/backend/routes"
	"github.com/rfornea/library/pkg/backend/utils"
)

func main() {
	utils.SetLive()

	models.Connect(utils.Env.DatabaseURL)

	routes.CreateRoutes()
}
