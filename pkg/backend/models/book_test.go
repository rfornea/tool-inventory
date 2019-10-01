package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_BookCreation_Success(t *testing.T) {
	setupTests(t)

	bookData := ReturnValidBookDataForTest(t)

	book := Book{
		BookData: bookData[0],
	}

	err := DB.Create(&book).Error
	assert.Nil(t, err)
}

func Test_BookCreation_Missing_Title(t *testing.T) {
	setupTests(t)

	bookData := ReturnValidBookDataForTest(t)

	book := Book{BookData: bookData[0]}
	book.Title = ""

	err := DB.Create(&book).Error
	assert.NotNil(t, err)
}

func Test_BookCreation_Missing_AuthorLastName(t *testing.T) {
	setupTests(t)

	bookData := ReturnValidBookDataForTest(t)

	book := Book{BookData: bookData[0]}
	book.AuthorLastName = ""

	err := DB.Create(&book).Error
	assert.NotNil(t, err)
}

func Test_BookCreation_Missing_AuthorFirstName(t *testing.T) {
	setupTests(t)

	bookData := ReturnValidBookDataForTest(t)

	book := Book{BookData: bookData[0]}
	book.AuthorFirstName = ""

	err := DB.Create(&book).Error
	assert.NotNil(t, err)
}

func Test_BookCreation_Missing_Isbn(t *testing.T) {
	setupTests(t)

	bookData := ReturnValidBookDataForTest(t)

	book := Book{BookData: bookData[0]}
	book.Isbn = ""

	err := DB.Create(&book).Error
	assert.NotNil(t, err)
}

func Test_BookCreation_Missing_Description(t *testing.T) {
	setupTests(t)

	bookData := ReturnValidBookDataForTest(t)

	book := Book{BookData: bookData[0]}
	book.Description = ""

	err := DB.Create(&book).Error
	assert.NotNil(t, err)
}

func Test_BookCreation_Invalid_Isbn(t *testing.T) {
	setupTests(t)

	bookData := ReturnValidBookDataForTest(t)

	book := Book{BookData: bookData[0]}
	book.Isbn = "3"

	err := DB.Create(&book).Error
	assert.NotNil(t, err)
}
