package database

import (
	"be-golang-chapter-46-implem/model"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
	)
}
