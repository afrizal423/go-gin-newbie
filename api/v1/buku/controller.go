package buku

import (
	"fmt"
	"html"
	"net/http"
	"strconv"
	"strings"

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
	data, err := c.service.ShowAllBuku()
	if err != nil {
		restErr := errors.NewBadRequestError("something went wrong")
		ctx.JSON(restErr.ErrStatus, restErr)
	}
	ctx.JSON(http.StatusOK, data)
}

func (c *Controller) GetBukuById(ctx *gin.Context) {
	bId := ctx.Param("bookId")
	bukuId, err := strconv.Atoi(html.EscapeString(strings.TrimSpace(bId)))
	if err != nil {
		fmt.Println("Error during conversion")
		return
	}

	data, err := c.service.GetBuku(bukuId)
	if err != nil {
		// jika datanya error
		restErr := errors.NewBadRequestError(fmt.Sprintf("gagal mendapatkan data: %v", err.Error()))
		ctx.JSON(restErr.ErrStatus, restErr)
	} else {
		ctx.JSON(http.StatusOK, data)
	}

}

func (c *Controller) UpdateBuku(ctx *gin.Context) {
	var buku models.Buku
	bId := ctx.Param("bookId")
	bukuId, err := strconv.Atoi(html.EscapeString(strings.TrimSpace(bId)))
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

	if err := c.service.UpdateBuku(bukuId, buku); err != nil {
		restErr := errors.NewBadRequestError(fmt.Sprintf("invalid update data: %v", err.Error()))
		ctx.JSON(restErr.ErrStatus, restErr)
		return
	}

	ctx.JSON(http.StatusCreated, "Updated")
}

func (c *Controller) HapusBuku(ctx *gin.Context) {
	bId := ctx.Param("bookId")
	bukuId, err := strconv.Atoi(bId)
	if err != nil {
		fmt.Println("Error during conversion")
		return
	}
	if err := c.service.DeleteBuku(bukuId); err != nil {
		// jika datanya error
		restErr := errors.NewBadRequestError(err.Error())
		ctx.JSON(restErr.ErrStatus, restErr)
		return
	}
	ctx.JSON(http.StatusCreated, "Deleted")
}
