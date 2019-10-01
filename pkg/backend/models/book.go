package models

import (
	"github.com/jinzhu/gorm"
	"github.com/rfornea/library/pkg/backend/utils"
	"time"
)

type BookData struct {
	Title           string `json:"title" binding:"required" gorm:"unique_index:idx_title_author_last_name_author_first_name"`
	AuthorLastName  string `json:"authorLastName" binding:"required" gorm:"unique_index:idx_title_author_last_name_author_first_name"`
	AuthorFirstName string `json:"authorFirstName" binding:"required" gorm:"unique_index:idx_title_author_last_name_author_first_name"`
	Isbn            string `json:"isbn" binding:"required,isbn" gorm:"unique"`
	Description     string `json:"description" binding:"required" gorm:"Type:mediumtext"`
	Available       bool   `json:"available"`
}

type Book struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	BookData
}

/*BeforeCreate - callback called before the row is created*/
func (b *Book) BeforeCreate(scope *gorm.Scope) error {
	return utils.Validator.Struct(b)
}

/*BeforeUpdate - callback called before the row is updated*/
func (b *Book) BeforeUpdate(scope *gorm.Scope) error {
	return utils.Validator.Struct(b)
}
