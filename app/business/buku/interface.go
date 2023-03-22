package buku

import (
	"database/sql"

	"github.com/afrizal423/go-gin-newbie/app/models"
)

type IBukuService interface {
	CreateBuku(data models.Buku) error
	ShowAllBuku() ([]models.Buku, error)
	GetBuku(id int) (*models.Buku, error)
	// UpdateBuku(id int, data models.Buku) bool
	// DeleteBuku(id int) bool
}

type IBukuRepository interface {
	CreateBukus(data models.Buku, db *sql.DB) error
	ShowAllBukus(db *sql.DB) ([]models.Buku, error)
	GetBukus(id int, db *sql.DB) (*models.Buku, error)
	// UpdateBukus(id int, data models.Buku) (models.Buku, bool)
	// DeleteBukus(id int) bool
}
