package buku

import (
	"github.com/afrizal423/go-gin-newbie/api/v1/buku/buku_response"
	"github.com/afrizal423/go-gin-newbie/app/models"
	"gorm.io/gorm"
)

type BukuRepository struct {
	db *gorm.DB
}

func NewBukuRepository(db *gorm.DB) *BukuRepository {
	return &BukuRepository{
		db,
	}
}

// repository menambahkan buku baru
func (bk *BukuRepository) CreateBukus(data models.Buku) (models.Buku, error) {
	var dt models.Buku
	if err := bk.db.Create(&data).Error; err != nil {
		return dt, err
	}
	return data, nil
}

// repository menampilkan semua buku
func (bk *BukuRepository) ShowAllBukus() ([]buku_response.BukuResponse, error) {
	var data []models.Buku
	var result []buku_response.BukuResponse
	if err := bk.db.Find(&data).Error; err != nil {
		return nil, err
	}
	for _, v := range data {
		var tmp buku_response.BukuResponse
		tmp.Id = v.Id
		tmp.Title = v.Title
		tmp.Author = v.Author
		tmp.CreatedAt = v.CreatedAt
		tmp.UpdatedAt = v.UpdatedAt
		result = append(result, tmp)
	}

	return result, nil
}

// repository untuk menampilkan detail buku
func (bk *BukuRepository) GetBukus(id int) (*models.Buku, error) {
	var data models.Buku
	if err := bk.db.First(&data, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

// repository untuk mengupdate buku
func (bk *BukuRepository) UpdateBukus(id int, data models.Buku) (models.Buku, error) {
	var dt models.Buku
	if err := bk.db.Where("id = ?", id).Updates(&data).Error; err != nil {
		return dt, err
	}
	if err := bk.db.First(&data, "id = ?", id).Error; err != nil {
		return dt, err
	}
	return data, nil
}

func (bk *BukuRepository) DeleteBukus(id int) error {
	var data models.Buku
	if err := bk.db.Where("id = ?", id).Delete(&data).Error; err != nil {
		return err
	}
	return nil
}

// func (bk *BukuRepository) HitungBukus(id int, db *sql.DB) error {
// 	// kita tutup koneksinya di akhir proses
// 	//defer configs.PostgresClose(db)

// 	rows, err := db.Query("SELECT COUNT(*) FROM book WHERE id=$1", id)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	var count int

// 	for rows.Next() {
// 		if err := rows.Scan(&count); err != nil {
// 			log.Fatal(err)
// 		}
// 	}

// 	fmt.Printf("Number of rows are %d\n", count)

// 	if count == 0 {
// 		return errors.New("tidak ada data yang dicari")
// 	}
// 	return nil
// }
