package routes

import (
	"github.com/rfornea/tool-inventory/pkg/backend/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_GetAllToolsHandler_success(t *testing.T) {
	setupTests(t)

	tools := models.MakeToolsForTest(t)

	w := httpRequestHelperForTest(t, ToolsPath, nil, http.MethodGet)

	confirmExpectedResponse(t, w.Code, http.StatusOK)

	assert.Contains(t, w.Body.String(), tools[0].Name)
	assert.Contains(t, w.Body.String(), tools[1].Name)
}

func Test_GetAllToolsHandler_no_tools(t *testing.T) {
	setupTests(t)

	w := httpRequestHelperForTest(t, ToolsPath, nil, http.MethodGet)

	confirmExpectedResponse(t, w.Code, http.StatusOK)
	assert.Contains(t, w.Body.String(), `"tools":[]`)
}
