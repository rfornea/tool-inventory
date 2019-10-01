package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rfornea/library/pkg/backend/models"
	"github.com/rfornea/library/pkg/backend/utils"
	"strings"
)

func UpdateBookHandler(c *gin.Context) {
	id := c.Param("id")

	req := models.BookData{}

	if err := utils.ParseRequestBody(c.Request, &req); err != nil {
		BadRequestResponse(c, err)
		return
	}

	var book models.Book

	if err := models.DB.Where("id = ?", id).Find(&book).Error; err != nil {
		NotFoundResponse(c, err)
		return
	}

	book.Title = req.Title
	book.AuthorFirstName = req.AuthorFirstName
	book.AuthorLastName = req.AuthorLastName
	book.Description = req.Description
	book.Isbn = req.Isbn
	book.Available = req.Available

	if err := models.DB.Save(&book).Error; err != nil {
		if strings.Contains(err.Error(), models.DuplicateEntryStr) {
			ConflictResponse(c, err)
			return
		}
		InternalErrorResponse(c, err)
		return
	}

	OkStatusOk(c, book)
}
