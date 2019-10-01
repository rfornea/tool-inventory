package routes

import (
	"github.com/rfornea/library/pkg/backend/models"
	"net/http"
	"strconv"
	"testing"
)

func Test_UpdateBookHandler_success(t *testing.T) {
	setupTests(t)

	books := models.MakeBooksForTest(t)

	updatedBook := models.BookData{
		Title:           books[0].Title + " Volume 2",
		Isbn:            books[0].Isbn,
		AuthorLastName:  books[0].AuthorLastName,
		AuthorFirstName: books[0].AuthorFirstName,
		Description:     books[0].Description,
	}

	w := httpRequestHelperForTest(t, BooksPath+"/"+strconv.Itoa(int(books[0].ID)), updatedBook, http.MethodPut)

	confirmExpectedResponse(t, w.Code, http.StatusOK)
}

func Test_UpdateBookHandler_book_does_not_exist(t *testing.T) {
	setupTests(t)

	books := models.MakeBooksForTest(t)

	updatedBook := models.BookData{
		Title:           books[0].Title + " Volume 2",
		Isbn:            books[0].Isbn,
		AuthorLastName:  books[0].AuthorLastName,
		AuthorFirstName: books[0].AuthorFirstName,
		Description:     books[0].Description,
	}

	// delete the book that we are about to try to update
	models.DB.Where("id = ?", books[0].ID).Delete(&models.Book{})

	w := httpRequestHelperForTest(t, BooksPath+"/"+strconv.Itoa(int(books[0].ID)), updatedBook, http.MethodPut)

	confirmExpectedResponse(t, w.Code, http.StatusNotFound)
}

func Test_UpdateBookHandler_isbn_conflict(t *testing.T) {
	setupTests(t)

	books := models.MakeBooksForTest(t)

	// cause conflict with another entry in DB by using same isbn
	updatedBook := models.BookData{
		Title:           books[0].Title,
		Isbn:            books[1].Isbn,
		AuthorLastName:  books[0].AuthorLastName,
		AuthorFirstName: books[0].AuthorFirstName,
		Description:     books[0].Description,
	}

	w := httpRequestHelperForTest(t, BooksPath+"/"+strconv.Itoa(int(books[0].ID)), updatedBook, http.MethodPut)

	confirmExpectedResponse(t, w.Code, http.StatusConflict)
}

func Test_UpdateBookHandler_author_name_and_title_conflict(t *testing.T) {
	setupTests(t)

	books := models.MakeBooksForTest(t)

	// cause conflict with another entry in DB by using same title and author name
	updatedBook := models.BookData{
		Title:           books[1].Title,
		Isbn:            books[0].Isbn,
		AuthorLastName:  books[1].AuthorLastName,
		AuthorFirstName: books[1].AuthorFirstName,
		Description:     books[0].Description,
	}

	w := httpRequestHelperForTest(t, BooksPath+"/"+strconv.Itoa(int(books[0].ID)), updatedBook, http.MethodPut)

	confirmExpectedResponse(t, w.Code, http.StatusConflict)
}

func Test_UpdateBookHandler_bad_request(t *testing.T) {
	setupTests(t)

	books := models.MakeBooksForTest(t)

	// cause bad request by making title empty
	updatedBook := models.BookData{
		Title:           "",
		Isbn:            books[0].Isbn,
		AuthorLastName:  books[0].AuthorLastName,
		AuthorFirstName: books[0].AuthorFirstName,
		Description:     books[0].Description,
	}

	w := httpRequestHelperForTest(t, BooksPath+"/"+strconv.Itoa(int(books[0].ID)), updatedBook, http.MethodPut)

	confirmExpectedResponse(t, w.Code, http.StatusBadRequest)
}
