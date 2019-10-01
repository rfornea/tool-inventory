package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rfornea/library/pkg/backend/models"
)

type getAllBooksRes struct {
	Books []models.Book `json:"books"`
}

func GetAllBooksHandler(c *gin.Context) {
	var books []models.Book

	if err := models.DB.Find(&books).Error; err != nil {
		InternalErrorResponse(c, err)
		return
	}

	OkStatusOk(c, getAllBooksRes{
		Books: books,
	})
}
