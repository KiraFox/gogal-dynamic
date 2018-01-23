package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model // Has attributes id, created/updated/deleted_at
	Name       string
	Email      string `gorm:"not null;unique_index"` // Tell GORM this field is
	// required (not an empty string) and no two users can have same email
}
