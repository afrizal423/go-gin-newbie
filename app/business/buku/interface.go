package buku

import "github.com/afrizal423/go-gin-newbie/app/models"

type IBukuService interface {
	CreateBuku(data models.Buku) error
	ShowAllBuku() ([]models.Buku, error)
	GetBuku(id int) (*models.Buku, error)
	// UpdateBuku(id int, data models.Buku) bool
	// DeleteBuku(id int) bool
}

type IBukuRepository interface {
	CreateBukus(data models.Buku) error
	ShowAllBukus() ([]models.Buku, error)
	GetBukus(id int) (*models.Buku, error)
	// UpdateBukus(id int, data models.Buku) (models.Buku, bool)
	// DeleteBukus(id int) bool
}
