package routes

import (
	"github.com/rfornea/library/pkg/backend/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_GetAllBooksHandler_success(t *testing.T) {
	setupTests(t)

	books := models.MakeBooksForTest(t)

	w := httpRequestHelperForTest(t, BooksPath, nil, http.MethodGet)

	confirmExpectedResponse(t, w.Code, http.StatusOK)

	assert.Contains(t, w.Body.String(), books[0].Title)
	assert.Contains(t, w.Body.String(), books[1].Title)
}

func Test_GetAllBooksHandler_no_books(t *testing.T) {
	setupTests(t)

	w := httpRequestHelperForTest(t, BooksPath, nil, http.MethodGet)

	confirmExpectedResponse(t, w.Code, http.StatusOK)
	assert.Contains(t, w.Body.String(), `"books":[]`)
}
