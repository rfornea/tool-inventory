package routes

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rfornea/tool-inventory/pkg/backend/models"
	"io/ioutil"
	//"fmt"
	"os"
	"strconv"
)

func CreateToolHandler(c *gin.Context) {
	c.Request.ParseMultipartForm(MaxRequestSize)

	tool := models.Tool{
		ToolData: models.ToolData{
			Name:         c.Request.Form["name"][0],
			Model:        c.Request.Form["model"][0],
			Manufacturer: c.Request.Form["manufacturer"][0],
			Description:  c.Request.Form["description"][0],
		},
	}

	if err := models.DB.Create(&tool).Error; err != nil {
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

	OkStatusCreated(c, tool)
}
