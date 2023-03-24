package buku_response

import (
	"time"

	"github.com/afrizal423/go-gin-newbie/app/models"
)

// struct ini untuk menginisialisai response data ke user
type BukuResponse struct {
	Id        int       `gorm:"primary_key;" json:"id"`
	Title     string    `json:"name_book"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// function ini untuk menampilkan single data
func SingleDataResponse(data *models.Buku) BukuResponse {
	var response BukuResponse
	response.Id = data.Id
	response.Title = data.Title
	response.Author = data.Author
	response.CreatedAt = data.CreatedAt
	response.UpdatedAt = data.UpdatedAt

	return response
}
