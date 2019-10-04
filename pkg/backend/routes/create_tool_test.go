package routes

import (
	"github.com/rfornea/tool-inventory/pkg/backend/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_CreateToolHandler_success(t *testing.T) {
	setupTests(t)

	newTool := models.ReturnValidToolDataForTest(t)

	form := map[string]string{
		"name":         newTool[0].Name,
		"model":        newTool[0].Model,
		"manufacturer": newTool[0].Manufacturer,
		"description":  newTool[0].Description,
	}

	w := httpFormRequestHelperForTest(t, ToolsPath, form, http.MethodPost)

	confirmExpectedResponse(t, w.Code, http.StatusCreated)

	assert.Contains(t, w.Body.String(), newTool[0].Name)
	assert.Contains(t, w.Body.String(), newTool[0].Description)
	assert.Contains(t, w.Body.String(), newTool[0].Model)
	assert.Contains(t, w.Body.String(), newTool[0].Manufacturer)
}

func Test_CreateToolHandler_missing_required_field(t *testing.T) {
	setupTests(t)

	newTool := models.ReturnValidToolDataForTest(t)
	newTool[0].Name = ""

	form := map[string]string{
		"name":         newTool[0].Name,
		"model":        newTool[0].Model,
		"manufacturer": newTool[0].Manufacturer,
		"description":  newTool[0].Description,
	}

	w := httpFormRequestHelperForTest(t, ToolsPath, form, http.MethodPost)

	confirmExpectedResponse(t, w.Code, http.StatusBadRequest)
}
