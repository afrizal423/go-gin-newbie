package buku

import (
	"github.com/afrizal423/go-gin-newbie/app/models"
	"github.com/afrizal423/go-gin-newbie/configs"
)

type bukuService struct {
	repository IBukuRepository
}

func NewBukuService(repository IBukuRepository) *bukuService {
	return &bukuService{
		repository,
	}
}

func (b *bukuService) CreateBuku(data models.Buku) error {
	db := configs.PostgresConn()
	if err := b.repository.CreateBukus(data, db); err != nil {
		return err
	}
	return nil
}

func (b *bukuService) ShowAllBuku() ([]models.Buku, error) {
	db := configs.PostgresConn()
	data, err := b.repository.ShowAllBukus(db)
	if err != nil {
		return []models.Buku{}, nil
	}
	return data, nil
}

func (b *bukuService) GetBuku(id int) (*models.Buku, error) {
	db := configs.PostgresConn()
	bukuById, err := b.repository.GetBukus(id, db)
	if err != nil {
		return nil, err
	}
	return bukuById, nil
}

func (b *bukuService) UpdateBuku(id int, data models.Buku) error {
	db := configs.PostgresConn()
	if err := b.repository.UpdateBukus(id, data, db); err != nil {
		return err
	}
	return nil
}

func (b *bukuService) DeleteBuku(id int) error {
	db := configs.PostgresConn()
	if err := b.repository.DeleteBukus(id, db); err != nil {
		return err
	}
	return nil
}
