package routes

import (
	"github.com/rfornea/tool-inventory/pkg/backend/models"
	"net/http"
	"strconv"
	"testing"
)

func Test_DeleteToolHandler_success(t *testing.T) {
	setupTests(t)

	tools := models.MakeToolsForTest(t)

	w := httpRequestHelperForTest(t, ToolsPath+"/"+strconv.Itoa(int(tools[0].ID)), nil, http.MethodDelete)

	confirmExpectedResponse(t, w.Code, http.StatusOK)
}

func Test_DeleteToolHandler_tool_does_not_exist(t *testing.T) {
	setupTests(t)

	w := httpRequestHelperForTest(t, ToolsPath+"/1", nil, http.MethodDelete)

	confirmExpectedResponse(t, w.Code, http.StatusNotFound)
}
