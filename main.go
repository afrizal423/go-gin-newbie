package main

import (
	bukuController "github.com/afrizal423/go-gin-newbie/api/v1/buku"
	bukuService "github.com/afrizal423/go-gin-newbie/app/business/buku"
	bukuRepository "github.com/afrizal423/go-gin-newbie/app/repository/buku"
	"github.com/afrizal423/go-gin-newbie/configs"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func main() {
	data := configs.InitData()

	bukuHandler := bukuController.NewBukuController(bukuService.NewBukuService(
		bukuRepository.NewBukuRepository(data)))

	router.GET("/books", bukuHandler.ShowAllBuku)
	router.GET("/books/:bookId", bukuHandler.GetBukuById)
	router.PUT("/books/:bookId", bukuHandler.UpdateBuku)
	router.POST("/books", bukuHandler.TambahBuku)
	router.DELETE("/books/:bookId", bukuHandler.HapusBuku)

	router.Run(":8000")
}
