package routes

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rfornea/tool-inventory/pkg/backend/models"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func UpdateToolHandler(c *gin.Context) {
	id := c.Param("id")

	c.Request.ParseMultipartForm(MaxRequestSize)

	req := models.ToolData{
		Name:         c.Request.Form["name"][0],
		Model:        c.Request.Form["model"][0],
		Manufacturer: c.Request.Form["manufacturer"][0],
		Description:  c.Request.Form["description"][0],
	}

	var tool models.Tool

	if err := models.DB.Where("id = ?", id).Find(&tool).Error; err != nil {
		NotFoundResponse(c, err)
		return
	}

	tool.Name = req.Name
	tool.Model = req.Model
	tool.Manufacturer = req.Manufacturer
	tool.Description = req.Description

	if err := models.DB.Save(&tool).Error; err != nil {
		if strings.Contains(err.Error(), models.DuplicateEntryStr) {
			ConflictResponse(c, err)
			return
		}
		BadRequestResponse(c, err)
		return
	}

	file, handler, err := c.Request.FormFile("file")
	if err == nil {
		defer file.Close()
		if handler.Size > 250000 {
			BadRequestResponse(c, errors.New("file too large"))
			return
		}
		os.Remove("assets/images/" + strconv.FormatUint(uint64(tool.ID), 10) + ".png")
		f, err := os.Create("assets/images/" + strconv.FormatUint(uint64(tool.ID), 10) + ".png")
		if err != nil {
			InternalErrorResponse(c, err)
			return
		}
		defer f.Close()

		// read all of the contents of our uploaded file into a
		// byte array
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			InternalErrorResponse(c, err)
			return
		}
		// write this byte array to our temporary file
		f.Write(fileBytes)
	}

	OkStatusOk(c, tool)
}
