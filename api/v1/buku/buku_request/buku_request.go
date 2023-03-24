package buku_request

import (
	"github.com/afrizal423/go-gin-newbie/app/models"
)

// struct ini untuk menampung data dari request user
type BukuRequest struct {
	Title  string `json:"name_book"`
	Author string `json:"author"`
}

// function untuk mengkonversi dari struct request ke models buku
func (request *BukuRequest) CreateJenisBuku() *models.Buku {
	var book models.Buku
	book.Title = request.Title
	book.Author = request.Author
	return &book
}
