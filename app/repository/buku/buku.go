package buku

import (
	"context"
	"database/sql"
	"errors"
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
		return &data, errors.New("tidak ada data yang dicari")
	case nil:
		return &data, nil
	default:
		log.Fatalf("tidak bisa mengambil data. %v", err)
	}
	return &data, err
}

func (bk *BukuRepository) UpdateBukus(id int, data models.Buku, db *sql.DB) error {
	// kita tutup koneksinya di akhir proses
	defer configs.PostgresClose(db)

	// kita buat sql query update
	sqlStatement := `UPDATE "book" SET "title"=$2, "author"=$3, "desc"=$4 WHERE "id"=$1`

	// eksekusi sql statement
	res, err := db.Exec(sqlStatement, id, data.Title, data.Author, data.Desc)

	if err != nil {
		return err
	}

	// cek berapa banyak row/data yang diupdate
	rowsAffected, err := res.RowsAffected()

	//kita cek
	if err != nil {
		return err
	}

	fmt.Printf("Total rows/record yang diupdate %v\n", rowsAffected)
	return err
}

func (bk *BukuRepository) DeleteBukus(id int, db *sql.DB) error {
	// kita tutup koneksinya di akhir proses
	defer configs.PostgresClose(db)
	// buat sql query
	sqlStatement := `DELETE FROM book WHERE id=$1`

	// eksekusi sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	// cek berapa jumlah data/row yang di hapus
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("tidak bisa mencari data. %v", err)
	}

	fmt.Printf("Total data yang terhapus %v", rowsAffected)

	return err
}

func (bk *BukuRepository) HitungBukus(id int, db *sql.DB) error {
	// kita tutup koneksinya di akhir proses
	//defer configs.PostgresClose(db)

	rows, err := db.Query("SELECT COUNT(*) FROM book WHERE id=$1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var count int

	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("Number of rows are %d\n", count)

	if count == 0 {
		return errors.New("tidak ada data yang dicari")
	}
	return nil
}
