package buku

import (
	"github.com/afrizal423/go-gin-newbie/api/v1/buku/buku_response"
	"github.com/afrizal423/go-gin-newbie/app/models"
)

type IBukuService interface {
	// services proses membuat data buku baru
	CreateBuku(data models.Buku) (models.Buku, error)
	// services proses menampilkan data buku semua
	ShowAllBuku() ([]buku_response.BukuResponse, error)
	// GetBuku(id int) (*models.Buku, error)
	// UpdateBuku(id int, data models.Buku) error
	// DeleteBuku(id int) error
}

type IBukuRepository interface {
	// repository menambahkan buku baru
	CreateBukus(data models.Buku) (models.Buku, error)
	// repository menampilkan semua buku
	ShowAllBukus() ([]buku_response.BukuResponse, error)
	// GetBukus(id int, db *sql.DB) (*models.Buku, error)
	// UpdateBukus(id int, data models.Buku, db *sql.DB) error
	// DeleteBukus(id int, db *sql.DB) error
	// HitungBukus(id int, db *sql.DB) error
}
