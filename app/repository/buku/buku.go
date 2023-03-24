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

// func (bk *BukuRepository) GetBukus(id int, db *sql.DB) (*models.Buku, error) {
// 	// kita tutup koneksinya di akhir proses
// 	defer configs.PostgresClose(db)
// 	// inisialisasi variable buat nyimpan data hasil db
// 	var data models.Buku

// 	// buat sql query
// 	sqlStatement := `SELECT * FROM book WHERE id=$1`

// 	// eksekusi sql statement
// 	row := db.QueryRow(sqlStatement, id)

// 	err := row.Scan(&data.Id, &data.Title, &data.Author, &data.Desc)

// 	switch err {
// 	case sql.ErrNoRows:
// 		fmt.Println("Tidak ada data yang dicari!")
// 		return &data, errors.New("tidak ada data yang dicari")
// 	case nil:
// 		return &data, nil
// 	default:
// 		log.Fatalf("tidak bisa mengambil data. %v", err)
// 	}
// 	return &data, err
// }

// func (bk *BukuRepository) UpdateBukus(id int, data models.Buku, db *sql.DB) error {
// 	// kita tutup koneksinya di akhir proses
// 	defer configs.PostgresClose(db)

// 	// kita buat sql query update
// 	sqlStatement := `UPDATE "book" SET "title"=$2, "author"=$3, "desc"=$4 WHERE "id"=$1`

// 	// eksekusi sql statement
// 	res, err := db.Exec(sqlStatement, id, data.Title, data.Author, data.Desc)

// 	if err != nil {
// 		return err
// 	}

// 	// cek berapa banyak row/data yang diupdate
// 	rowsAffected, err := res.RowsAffected()

// 	//kita cek
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Printf("Total rows/record yang diupdate %v\n", rowsAffected)
// 	return err
// }

// func (bk *BukuRepository) DeleteBukus(id int, db *sql.DB) error {
// 	// kita tutup koneksinya di akhir proses
// 	defer configs.PostgresClose(db)
// 	// buat sql query
// 	sqlStatement := `DELETE FROM book WHERE id=$1`

// 	// eksekusi sql statement
// 	res, err := db.Exec(sqlStatement, id)

// 	if err != nil {
// 		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
// 	}

// 	// cek berapa jumlah data/row yang di hapus
// 	rowsAffected, err := res.RowsAffected()

// 	if err != nil {
// 		log.Fatalf("tidak bisa mencari data. %v", err)
// 	}

// 	fmt.Printf("Total data yang terhapus %v", rowsAffected)

// 	return err
// }

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
