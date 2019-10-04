package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rfornea/tool-inventory/pkg/backend/models"
)

type getAllToolsRes struct {
	Tools []models.Tool `json:"tools"`
}

func GetAllToolsHandler(c *gin.Context) {
	var tools []models.Tool

	if err := models.DB.Find(&tools).Error; err != nil {
		InternalErrorResponse(c, err)
		return
	}

	OkStatusOk(c, getAllToolsRes{
		Tools: tools,
	})
}
