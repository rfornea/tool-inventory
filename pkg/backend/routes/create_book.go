package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rfornea/library/pkg/backend/models"
	"github.com/rfornea/library/pkg/backend/utils"
)

func CreateBookHandler(c *gin.Context) {
	req := models.BookData{}

	if err := utils.ParseRequestBody(c.Request, &req); err != nil {
		BadRequestResponse(c, err)
		return
	}

	newBook := models.Book{
		BookData: req,
	}

	if err := models.DB.Create(&newBook).Error; err != nil {
		InternalErrorResponse(c, err)
		return
	}

	OkStatusCreated(c, newBook)
}
