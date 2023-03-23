package database

import (
	"github.com/afrizal423/go-gin-newbie/app/models"
	"gorm.io/gorm"
)

func DbMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.Buku{})
}
