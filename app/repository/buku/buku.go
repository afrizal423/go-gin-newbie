package buku

import "github.com/afrizal423/go-gin-newbie/app/models"

type BukuRepository struct {
	Data_buku []models.Buku
}

func NewBukuRepository(databuku []models.Buku) *BukuRepository {
	return &BukuRepository{
		databuku,
	}
}

func (bk *BukuRepository) CreateBukus(data models.Buku) ([]models.Buku, bool) {
	data.Id = len(bk.Data_buku) + 1
	bk.Data_buku = append(bk.Data_buku, data)
	return bk.Data_buku, true
}

func (bk *BukuRepository) ShowAllBukus() []models.Buku {
	return bk.Data_buku
}
