package buku

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/afrizal423/go-gin-newbie/app/models"
	"github.com/afrizal423/go-gin-newbie/configs"
)

type BukuRepository struct {
	Ctx context.Context
}

func NewBukuRepository(ctx context.Context) *BukuRepository {
	return &BukuRepository{
		ctx,
	}
}

func (bk *BukuRepository) CreateBukus(data models.Buku, db *sql.DB) error {
	// kita tutup koneksinya di akhir proses
	defer configs.PostgresClose(db)

	// kita buat insert query
	// mengembalikan nilai id akan mengembalikan id dari buku yang dimasukkan ke db
	sqlStatement := `INSERT INTO "book" ("title", "author", "desc") VALUES ($1, $2, $3)`
	_, e := db.Exec(sqlStatement, data.Title, data.Author, data.Desc)

	if e != nil {
		log.Fatalf("Tidak Bisa mengeksekusi query. %v", e)
	}

	return e
}

func (bk *BukuRepository) ShowAllBukus(db *sql.DB) ([]models.Buku, error) {
	// kita tutup koneksinya di akhir proses
	defer configs.PostgresClose(db)

	// variable menampung data
	var bukus []models.Buku

	// kita buat select query
	sqlStatement := `SELECT * FROM book`

	// mengeksekusi sql query
	rows, err := db.Query(sqlStatement)

	if err != nil {
		return nil, err
	}

	// kita tutup eksekusi proses sql qeurynya
	defer rows.Close()
	// kita iterasi mengambil datanya
	for rows.Next() {
		// temporary data
		var booktemp models.Buku
		// kita ambil datanya dan unmarshal ke structnya
		err = rows.Scan(&booktemp.Id, &booktemp.Title, &booktemp.Author, &booktemp.Desc)

		if err != nil {
			return nil, err
		}

		// masukkan kedalam slice bukus
		bukus = append(bukus, booktemp)
	}
	return bukus, nil
}

func (bk *BukuRepository) GetBukus(id int, db *sql.DB) (*models.Buku, error) {
	// kita tutup koneksinya di akhir proses
	defer configs.PostgresClose(db)
	// inisialisasi variable buat nyimpan data hasil db
	var data models.Buku

	// buat sql query
	sqlStatement := `SELECT * FROM book WHERE id=$1`

	// eksekusi sql statement
	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&data.Id, &data.Title, &data.Author, &data.Desc)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Tidak ada data yang dicari!")
		return &data, nil
	case nil:
		return &data, nil
	default:
		log.Fatalf("tidak bisa mengambil data. %v", err)
	}
	return &data, err
}

// func (bk *BukuRepository) UpdateBukus(id int, data models.Buku) (models.Buku, bool) {
// 	var datas models.Buku
// 	for i, m := range bk.Data_buku {
// 		if m.Id == id {
// 			bk.Data_buku[i].Id = m.Id
// 			bk.Data_buku[i].Author = data.Author
// 			bk.Data_buku[i].Title = data.Title
// 			bk.Data_buku[i].Desc = data.Desc
// 			return bk.Data_buku[i], true
// 		}
// 	}
// 	return datas, false
// }

// func (bk *BukuRepository) DeleteBukus(id int) bool {
// 	for i, m := range bk.Data_buku {
// 		if m.Id == id {
// 			// Menghapus data buku dengan id n dari array
// 			bk.Data_buku = append(bk.Data_buku[:i], bk.Data_buku[i+1:]...)
// 			return true
// 		}
// 	}
// 	return false
// }
