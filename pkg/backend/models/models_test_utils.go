package models

import (
	"github.com/gin-gonic/gin"
	"github.com/rfornea/tool-inventory/pkg/backend/utils"
	"testing"
)

const testDatabaseErrorString = "should only be calling this method on test database"

func setupTests(t *testing.T) {
	utils.SetTesting("../.env")
	Connect(utils.Env.TestDatabaseURL)
	DeleteToolsForTest(t)
	gin.SetMode(gin.TestMode)
}

func ReturnValidToolDataForTest(t *testing.T) []ToolData {
	utils.AbortIfNotTesting(t)
	return []ToolData{
		{
			Name:         "Hammer",
			Model:        "00201",
			Manufacturer: "Black and Decker",
			Description:  "Use this for nails",
		},
		{
			Name:         "Screwdriver",
			Model:        "917",
			Manufacturer: "Ryobi",
			Description:  "Use this for screws",
		},
	}
}

func MakeToolsForTest(t *testing.T) []Tool {
	utils.AbortIfNotTesting(t)

	var tools []Tool

	toolData := ReturnValidToolDataForTest(t)

	for _, val := range toolData {
		tool := Tool{
			ToolData: val,
		}
		DB.Create(&tool)
		tools = append(tools, tool)
	}
	return tools
}

func DeleteToolsForTest(t *testing.T) {
	deleteForTest(t, "tools")
}

func deleteForTest(t *testing.T, table string) {
	if utils.Env.DatabaseURL != utils.Env.TestDatabaseURL {
		t.Fatalf(testDatabaseErrorString)
	} else {
		cmd := "DELETE from " + table + ";"
		DB.Exec(cmd)
	}
}
