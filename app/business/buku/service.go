package buku

import "github.com/afrizal423/go-gin-newbie/app/models"

type bukuService struct {
	repository IBukuRepository
}

func NewBukuService(repository IBukuRepository) *bukuService {
	return &bukuService{
		repository,
	}
}

func (b *bukuService) CreateBuku(data models.Buku) error {
	if err := b.repository.CreateBukus(data); err != nil {
		return err
	}
	return nil
}

func (b *bukuService) ShowAllBuku() ([]models.Buku, error) {
	data, err := b.repository.ShowAllBukus()
	if err != nil {
		return []models.Buku{}, nil
	}
	return data, nil
}

func (b *bukuService) GetBuku(id int) (*models.Buku, error) {
	bukuById, err := b.repository.GetBukus(id)
	if err != nil {
		return nil, err
	}
	return bukuById, nil
}

// func (b *bukuService) UpdateBuku(id int, data models.Buku) bool {
// 	_, status := b.repository.UpdateBukus(id, data)
// 	return status
// }

// func (b *bukuService) DeleteBuku(id int) bool {
// 	return b.repository.DeleteBukus(id)
// }
