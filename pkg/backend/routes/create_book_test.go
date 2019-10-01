package routes

import (
	"github.com/rfornea/library/pkg/backend/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_CreateBookHandler_success(t *testing.T) {
	setupTests(t)

	newBook := models.ReturnValidBookDataForTest(t)

	w := httpRequestHelperForTest(t, BooksPath, newBook[0], http.MethodPost)

	confirmExpectedResponse(t, w.Code, http.StatusCreated)

	assert.Contains(t, w.Body.String(), newBook[0].Title)
	assert.Contains(t, w.Body.String(), newBook[0].Description)
	assert.Contains(t, w.Body.String(), newBook[0].AuthorLastName)
	assert.Contains(t, w.Body.String(), newBook[0].AuthorFirstName)
}

func Test_CreateBookHandler_missing_required_field(t *testing.T) {
	setupTests(t)

	newBook := models.ReturnValidBookDataForTest(t)
	newBook[0].Title = ""

	w := httpRequestHelperForTest(t, BooksPath, newBook[0], http.MethodPost)

	confirmExpectedResponse(t, w.Code, http.StatusBadRequest)
}

func Test_CreateBookHandler_invalid_isbn(t *testing.T) {
	setupTests(t)

	newBook := models.ReturnValidBookDataForTest(t)
	newBook[0].Isbn = "333"

	w := httpRequestHelperForTest(t, BooksPath, newBook[0], http.MethodPost)

	confirmExpectedResponse(t, w.Code, http.StatusBadRequest)
}
