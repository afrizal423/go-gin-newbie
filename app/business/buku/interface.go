package buku

import "github.com/afrizal423/go-gin-newbie/app/models"

type IBukuService interface {
	CreateBuku(data models.Buku) bool
	ShowAllBuku() []models.Buku
	GetBuku(id int) (models.Buku, bool)
}

type IBukuRepository interface {
	CreateBukus(data models.Buku) ([]models.Buku, bool)
	ShowAllBukus() []models.Buku
	GetBukus(id int) (models.Buku, bool)
}
