package routes

import (
	"github.com/rfornea/library/pkg/backend/models"
	"net/http"
	"strconv"
	"testing"
)

func Test_DeleteBookHandler_success(t *testing.T) {
	setupTests(t)

	books := models.MakeBooksForTest(t)

	w := httpRequestHelperForTest(t, BooksPath+"/"+strconv.Itoa(int(books[0].ID)), nil, http.MethodDelete)

	confirmExpectedResponse(t, w.Code, http.StatusOK)
}

func Test_DeleteBookHandler_book_does_not_exist(t *testing.T) {
	setupTests(t)

	w := httpRequestHelperForTest(t, BooksPath+"/1", nil, http.MethodDelete)

	confirmExpectedResponse(t, w.Code, http.StatusNotFound)
}
