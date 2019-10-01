package routes

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/rfornea/library/pkg/backend/models"
	"github.com/rfornea/library/pkg/backend/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTests(t *testing.T) {
	utils.SetTesting("../.env")
	models.Connect(utils.Env.TestDatabaseURL)
	models.DeleteTripsForTest(t)
	gin.SetMode(gin.TestMode)
}

func httpRequestHelperForTest(t *testing.T, path string, request interface{}, method string) *httptest.ResponseRecorder {
	utils.AbortIfNotTesting(t)

	router := returnEngine()
	v1 := returnV1Group(router)
	setupV1Paths(v1)

	marshalledReq, _ := json.Marshal(request)
	reqBody := bytes.NewBuffer(marshalledReq)

	// Create the mock request you'd like to test. Make sure the second argument
	// here is the same as one of the routes you defined in the router setup
	// block!
	req, err := http.NewRequest(method, v1.BasePath()+path, reqBody)

	if err != nil {
		assert.Fail(t, "Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	return w
}

func confirmExpectedResponse(t *testing.T, receivedCode, expectedCode int) {
	// Check to see if the response was what you expected
	if receivedCode != expectedCode {
		t.Fatalf("Expected to get status %d but instead got %d\n", expectedCode, receivedCode)
	}
}
