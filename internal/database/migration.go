package database

import (
	"github.com/Sairan-ds/golang-rest-api-course/internal/comment"
	"gorm.io/gorm"
)

func MigrateDb(db *gorm.DB) error {
	if result := db.AutoMigrate(&comment.Comment{}); result != nil {
		return result
	}
	return nil
}
