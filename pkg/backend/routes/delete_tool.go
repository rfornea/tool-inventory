package routes

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rfornea/tool-inventory/pkg/backend/models"
)

func DeleteToolHandler(c *gin.Context) {
	id := c.Param("id")

	if rowsAffected := models.DB.Where("id = ?", id).Delete(&models.Tool{}).RowsAffected; rowsAffected == int64(0) {
		NotFoundResponse(c, errors.New("cannot delete resource that does not exist"))
		return
	}

	OkStatusOk(c, StatusRes{
		Status: "resource deleted",
	})
}
