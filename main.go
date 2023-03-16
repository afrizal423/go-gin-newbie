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

	userHandler := bukuController.NewBukuController(bukuService.NewBukuService(
		bukuRepository.NewBukuRepository(data)))

	router.GET("/books", userHandler.ShowAllBuku)
	router.GET("/books/:bookId", userHandler.GetBukuById)
	router.POST("/books", userHandler.TambahBuku)

	router.Run(":8000")
}
