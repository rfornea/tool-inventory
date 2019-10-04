package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ToolCreation_Success(t *testing.T) {
	setupTests(t)

	toolData := ReturnValidToolDataForTest(t)

	tool := Tool{
		ToolData: toolData[0],
	}

	err := DB.Create(&tool).Error
	assert.Nil(t, err)
}

func Test_ToolCreation_Missing_Name(t *testing.T) {
	setupTests(t)

	toolData := ReturnValidToolDataForTest(t)

	tool := Tool{ToolData: toolData[0]}
	tool.Name = ""

	err := DB.Create(&tool).Error
	assert.NotNil(t, err)
}

func Test_ToolCreation_Missing_Model(t *testing.T) {
	setupTests(t)

	toolData := ReturnValidToolDataForTest(t)

	tool := Tool{ToolData: toolData[0]}
	tool.Model = ""

	err := DB.Create(&tool).Error
	assert.NotNil(t, err)
}

func Test_ToolCreation_Missing_Manufacturer(t *testing.T) {
	setupTests(t)

	toolData := ReturnValidToolDataForTest(t)

	tool := Tool{ToolData: toolData[0]}
	tool.Manufacturer = ""

	err := DB.Create(&tool).Error
	assert.NotNil(t, err)
}

func Test_ToolCreation_Missing_Description(t *testing.T) {
	setupTests(t)

	toolData := ReturnValidToolDataForTest(t)

	tool := Tool{ToolData: toolData[0]}
	tool.Description = ""

	err := DB.Create(&tool).Error
	assert.NotNil(t, err)
}
