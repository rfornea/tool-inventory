package routes

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rfornea/tool-inventory/pkg/backend/utils"
	"io"
	"net/http"
	"reflect"
	"strings"
)

const (
	postFormTag     = "form"
	postFormFileTag = "formFile"
	bindingTag      = "binding"
)

func verifyAndParseFormRequest(dest interface{}, c *gin.Context) error {
	defer c.Request.Body.Close()
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxRequestSize)
	err := c.Request.ParseMultipartForm(MaxRequestSize)
	if err != nil {
		return BadRequestResponse(c, err)
	}

	t := reflect.ValueOf(dest).Elem().Type()
	s := reflect.ValueOf(dest).Elem()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i) // Get the field, returns https://golang.org/pkg/reflect/#StructField
		if field.Anonymous {
			// Only go 1 level down
			for j := 0; j < field.Type.NumField(); j++ {
				nestField := field.Type.Field(j)
				strV, err := getValueFromPostForm(nestField, c)
				if err != nil {
					return err
				}
				if strV == "" {
					continue
				}
				if !s.Field(i).Field(j).CanSet() {
					return InternalErrorResponse(c, fmt.Errorf("Field is not settable, It should be upper case but has this: %v", nestField))
				}
				s.Field(i).Field(j).SetString(strV)
			}
		}

		strV, err := getValueFromPostForm(field, c)
		if err != nil {
			return err
		}
		if strV == "" {
			continue
		}
		if !s.Field(i).CanSet() {
			return InternalErrorResponse(c, fmt.Errorf("Field is not settable, It should be upper case but has this: %v", field))
		}
		s.Field(i).SetString(strV)
	}

	if ii, ok := dest.(parsableObjectInterface); ok {
		if err := utils.ParseStringifiedRequest(ii.getObjectAsString(), ii.getObjectRef()); err != nil {
			return BadRequestResponse(c, fmt.Errorf("bad request, unable to parse request body: %v", err))
		}
	}

	return nil
}

func getValueFromPostForm(field reflect.StructField, c *gin.Context) (string, error) {
	strV := ""
	formTag := field.Tag.Get(postFormTag)
	fileTag := field.Tag.Get(postFormFileTag)
	binding := field.Tag.Get(bindingTag)
	required := strings.Contains(binding, "required")

	if formTag == "" && fileTag == "" {
		return "", nil
	}
	if formTag != "" {
		strV = c.Request.FormValue(formTag)
	}

	if fileTag != "" {
		v, err := readFileFromForm(fileTag, c.Request)
		if err != nil && required {
			return "", BadRequestResponse(c, err)
		}
		// otherwise, just ignore the error
		strV = v
	}
	return strV, nil
}

func readFileFromForm(fileTag string, r *http.Request) (string, error) {
	multiFile, _, err := r.FormFile(fileTag)
	defer func() {
		if multiFile != nil {
			multiFile.Close()
		}
	}()

	if err != nil {
		return "", fmt.Errorf("Unable to get file %v from POST form", fileTag)
	}
	var fileBytes bytes.Buffer
	if _, err := io.Copy(&fileBytes, multiFile); err != nil {
		return "", err
	}
	return fileBytes.String(), nil
}
