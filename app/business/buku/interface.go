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
	// services proses menampilkan data detail buku
	GetBuku(id int) (*models.Buku, error)
	// services proses mengubah data buku
	UpdateBuku(id int, data models.Buku) (models.Buku, error)
	// services proses menghapus data buku
	DeleteBuku(id int) error
}

type IBukuRepository interface {
	// repository menambahkan buku baru
	CreateBukus(data models.Buku) (models.Buku, error)
	// repository menampilkan semua buku
	ShowAllBukus() ([]buku_response.BukuResponse, error)
	// repository untuk menampilkan detail buku
	GetBukus(id int) (*models.Buku, error)
	// repository untuk mengupdate buku
	UpdateBukus(id int, data models.Buku) (models.Buku, error)
	// repository untuk menghapus buku
	DeleteBukus(id int) error
	// HitungBukus(id int, db *sql.DB) error
}
