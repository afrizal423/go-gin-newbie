package buku

import (
	"github.com/afrizal423/go-gin-newbie/api/v1/buku/buku_response"
	"github.com/afrizal423/go-gin-newbie/app/models"
)

type bukuService struct {
	repository IBukuRepository
}

func NewBukuService(repository IBukuRepository) *bukuService {
	return &bukuService{
		repository,
	}
}

// services proses membuat data buku baru
func (b *bukuService) CreateBuku(data models.Buku) (models.Buku, error) {
	var dt2 models.Buku
	data, err := b.repository.CreateBukus(data)
	if err != nil {
		return dt2, err
	}
	return data, nil
}

// services proses menampilkan data buku semua
func (b *bukuService) ShowAllBuku() ([]buku_response.BukuResponse, error) {
	data, err := b.repository.ShowAllBukus()
	if err != nil {
		return []buku_response.BukuResponse{}, nil
	}
	return data, nil
}

// func (b *bukuService) GetBuku(id int) (*models.Buku, error) {
// 	db := configs.PostgresConn()
// 	bukuById, err := b.repository.GetBukus(id, db)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return bukuById, nil
// }

// func (b *bukuService) UpdateBuku(id int, data models.Buku) error {
// 	db := configs.PostgresConn()
// 	if err := b.repository.HitungBukus(id, db); err != nil {
// 		return err
// 	}
// 	if err := b.repository.UpdateBukus(id, data, db); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (b *bukuService) DeleteBuku(id int) error {
// 	db := configs.PostgresConn()
// 	if err := b.repository.HitungBukus(id, db); err != nil {
// 		return err
// 	}
// 	if err := b.repository.DeleteBukus(id, db); err != nil {
// 		return err
// 	}
// 	return nil
// }
