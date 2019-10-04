package routes

import (
	"github.com/rfornea/tool-inventory/pkg/backend/models"
	"net/http"
	"strconv"
	"testing"
)

func Test_UpdateToolHandler_success(t *testing.T) {
	setupTests(t)

	tools := models.MakeToolsForTest(t)

	form := map[string]string{
		"name":         tools[0].Name + " Version 2",
		"model":        tools[0].Model,
		"manufacturer": tools[0].Manufacturer,
		"description":  tools[0].Description,
	}

	w := httpFormRequestHelperForTest(t, ToolsPath+"/"+strconv.Itoa(int(tools[0].ID)), form, http.MethodPut)

	confirmExpectedResponse(t, w.Code, http.StatusOK)
}

func Test_UpdateToolHandler_tool_does_not_exist(t *testing.T) {
	setupTests(t)

	tools := models.MakeToolsForTest(t)

	form := map[string]string{
		"name":         tools[0].Name + " Version 2",
		"model":        tools[0].Model,
		"manufacturer": tools[0].Manufacturer,
		"description":  tools[0].Description,
	}
	// delete the tool that we are about to try to update
	models.DB.Where("id = ?", tools[0].ID).Delete(&models.Tool{})

	w := httpFormRequestHelperForTest(t, ToolsPath+"/"+strconv.Itoa(int(tools[0].ID)), form, http.MethodPut)

	confirmExpectedResponse(t, w.Code, http.StatusNotFound)
}

func Test_UpdateToolHandler_name_model_and_manufacturer_conflict(t *testing.T) {
	setupTests(t)

	tools := models.MakeToolsForTest(t)

	// cause conflict with another entry in DB by using same name, model, and manufacturer
	form := map[string]string{
		"name":         tools[1].Name,
		"model":        tools[1].Model,
		"manufacturer": tools[1].Manufacturer,
		"description":  tools[0].Description,
	}

	w := httpFormRequestHelperForTest(t, ToolsPath+"/"+strconv.Itoa(int(tools[0].ID)), form, http.MethodPut)

	confirmExpectedResponse(t, w.Code, http.StatusConflict)
}

func Test_UpdateToolHandler_bad_request(t *testing.T) {
	setupTests(t)

	tools := models.MakeToolsForTest(t)

	// cause bad request by making name empty
	form := map[string]string{
		"name":         "",
		"model":        tools[0].Model,
		"manufacturer": tools[0].Manufacturer,
		"description":  tools[0].Description,
	}

	w := httpFormRequestHelperForTest(t, ToolsPath+"/"+strconv.Itoa(int(tools[0].ID)), form, http.MethodPut)

	confirmExpectedResponse(t, w.Code, http.StatusBadRequest)
}
