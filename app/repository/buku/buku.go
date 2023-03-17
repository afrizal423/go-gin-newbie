package buku

import (
	"github.com/afrizal423/go-gin-newbie/app/models"
)

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

func (bk *BukuRepository) GetBukus(id int) (models.Buku, bool) {
	var data models.Buku
	for _, m := range bk.Data_buku {
		if m.Id == id {
			data = m
			return data, true

		}
	}
	return data, false
}

func (bk *BukuRepository) UpdateBukus(id int, data models.Buku) (models.Buku, bool) {
	var datas models.Buku
	for i, m := range bk.Data_buku {
		if m.Id == id {
			bk.Data_buku[i].Id = m.Id
			bk.Data_buku[i].Author = data.Author
			bk.Data_buku[i].Title = data.Title
			bk.Data_buku[i].Desc = data.Desc
			return bk.Data_buku[i], true
		}
	}
	return datas, false
}
