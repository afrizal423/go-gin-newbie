package main

import (
	bukuController "github.com/afrizal423/go-gin-newbie/api/v1/buku"
	bukuService "github.com/afrizal423/go-gin-newbie/app/business/buku"
	bukuRepository "github.com/afrizal423/go-gin-newbie/app/repository/buku"
	"github.com/afrizal423/go-gin-newbie/configs"
	"github.com/afrizal423/go-gin-newbie/database"
	_ "github.com/afrizal423/go-gin-newbie/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	router = gin.Default()
)

func main() {
	conn := configs.GormPostgresConn()
	// migrate db
	database.DbMigrate(conn)

	bukuHandler := bukuController.NewBukuController(bukuService.NewBukuService(
		bukuRepository.NewBukuRepository(conn)))

	router.GET("/books", bukuHandler.ShowAllBuku)
	router.GET("/books/:bookId", bukuHandler.GetBukuById)
	router.PUT("/books/:bookId", bukuHandler.UpdateBuku)
	router.POST("/books", bukuHandler.TambahBuku)
	router.DELETE("/books/:bookId", bukuHandler.HapusBuku)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}
