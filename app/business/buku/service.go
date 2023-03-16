package buku

import "github.com/afrizal423/go-gin-newbie/app/models"

type bukuService struct {
	repository IBukuRepository
}

func NewBukuService(repository IBukuRepository) IBukuService {
	return &bukuService{
		repository,
	}
}

func (b *bukuService) CreateBuku(data models.Buku) bool {
	_, status := b.repository.CreateBukus(data)
	return status
}

func (b *bukuService) ShowAllBuku() []models.Buku {
	return b.repository.ShowAllBukus()
}
