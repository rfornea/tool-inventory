package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rfornea/tool-inventory/pkg/backend/utils"
	"net/http"
)

const REQUEST_UUID = "request_uuid"

type handlerFunc func(*gin.Context) error

func InternalErrorResponse(c *gin.Context, err error) error {
	c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	return err
}

func BadRequestResponse(c *gin.Context, err error) error {
	c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	return err
}

func ServiceUnavailableResponse(c *gin.Context, err error) error {
	c.AbortWithStatusJSON(http.StatusServiceUnavailable, err.Error())
	return err
}

func ForbiddenResponse(c *gin.Context, err error) error {
	c.AbortWithStatusJSON(http.StatusForbidden, err.Error())
	return err
}

func ConflictResponse(c *gin.Context, err error) error {
	c.AbortWithStatusJSON(http.StatusConflict, err.Error())
	return err
}

func NotFoundResponse(c *gin.Context, err error) error {
	c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
	return err
}

func OkStatusCreated(c *gin.Context, response interface{}) error {
	if err := utils.Validator.Struct(response); err != nil {
		err = fmt.Errorf("could not create a valid response:  %v", err)
		return BadRequestResponse(c, err)
	}
	c.JSON(http.StatusCreated, response)
	return nil
}

func OkStatusOk(c *gin.Context, response interface{}) error {
	if err := utils.Validator.Struct(response); err != nil {
		err = fmt.Errorf("could not create a valid response:  %v", err)
		return BadRequestResponse(c, err)
	}
	c.JSON(http.StatusOK, response)
	return nil
}

func OkNoContent(c *gin.Context, response interface{}) error {
	if err := utils.Validator.Struct(response); err != nil {
		err = fmt.Errorf("could not create a valid response:  %v", err)
		return BadRequestResponse(c, err)
	}
	c.JSON(http.StatusNoContent, response)
	return nil
}
