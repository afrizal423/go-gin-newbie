package buku

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/afrizal423/go-gin-newbie/app/business/buku"
	"github.com/afrizal423/go-gin-newbie/app/models"
	"github.com/afrizal423/go-gin-newbie/utils/errors"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service buku.IBukuService
}

func NewBukuController(service buku.IBukuService) *Controller {
	return &Controller{
		service,
	}
}

func (c *Controller) TambahBuku(ctx *gin.Context) {
	var buku models.Buku
	if err := ctx.ShouldBindJSON(&buku); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.ErrStatus, restErr)
	}
	fmt.Println(buku)
	c.service.CreateBuku(buku)
	ctx.JSON(http.StatusCreated, "Created")
}

func (c *Controller) ShowAllBuku(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.service.ShowAllBuku())
}

func (c *Controller) GetBukuById(ctx *gin.Context) {
	bId := ctx.Param("bookId")
	bukuId, err := strconv.Atoi(bId)
	if err != nil {
		fmt.Println("Error during conversion")
		return
	}

	data, status := c.service.GetBuku(bukuId)
	if status {
		// jika datanya ada
		ctx.JSON(http.StatusOK, data)
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "data not found",
		})
	}

}

func (c *Controller) UpdateBuku(ctx *gin.Context) {
	var buku models.Buku
	bId := ctx.Param("bookId")
	bukuId, err := strconv.Atoi(bId)
	if err != nil {
		fmt.Println("Error during conversion")
		return
	}
	if err := ctx.ShouldBindJSON(&buku); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.ErrStatus, restErr)
		return
	}
	fmt.Println(buku, bukuId)
	status := c.service.UpdateBuku(bukuId, buku)
	if status {
		ctx.JSON(http.StatusCreated, "Updated")
		return
	} else {
		restErr := errors.NewBadRequestError("invalid update data")
		ctx.JSON(restErr.ErrStatus, restErr)
		return
	}
}

func (c *Controller) HapusBuku(ctx *gin.Context) {
	bId := ctx.Param("bookId")
	bukuId, err := strconv.Atoi(bId)
	if err != nil {
		fmt.Println("Error during conversion")
		return
	}
	status := c.service.DeleteBuku(bukuId)
	if status {
		ctx.JSON(http.StatusCreated, "Deleted")
		return
	} else {
		restErr := errors.NewBadRequestError("invalid delete data")
		ctx.JSON(restErr.ErrStatus, restErr)
		return
	}

}
