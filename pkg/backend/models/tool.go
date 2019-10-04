package models

import (
	"github.com/jinzhu/gorm"
	"github.com/rfornea/tool-inventory/pkg/backend/utils"
	"time"
)

type ToolData struct {
	Name         string `json:"name" binding:"required" gorm:"unique_index:idx_name_model_manufacturer"`
	Model        string `json:"model" binding:"required" gorm:"unique_index:idx_name_model_manufacturer"`
	Manufacturer string `json:"manufacturer" binding:"required" gorm:"unique_index:idx_name_model_manufacturer"`
	Description  string `json:"description" binding:"required" gorm:"Type:mediumtext"`
}

type Tool struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	ToolData
}

/*BeforeCreate - callback called before the row is created*/
func (t *Tool) BeforeCreate(scope *gorm.Scope) error {
	return utils.Validator.Struct(t)
}

/*BeforeUpdate - callback called before the row is updated*/
func (t *Tool) BeforeUpdate(scope *gorm.Scope) error {
	return utils.Validator.Struct(t)
}
