package buku

import (
	"fmt"
	"html"
	"net/http"
	"strconv"
	"strings"

	"github.com/afrizal423/go-gin-newbie/api/v1/buku/buku_request"
	"github.com/afrizal423/go-gin-newbie/api/v1/buku/buku_response"
	"github.com/afrizal423/go-gin-newbie/app/business/buku"
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

// Controller tambah buku
func (c *Controller) TambahBuku(ctx *gin.Context) {
	// alokasikan memori
	// https://dev.to/bhanu011/how-is-new-different-in-go-41lk
	// https://go.dev/doc/effective_go
	buku := new(buku_request.BukuRequest)
	// pastikan data harus berupa json
	if err := ctx.ShouldBindJSON(&buku); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.ErrStatus, restErr)
		return
	}

	// proses ke bagian services
	data, err := c.service.CreateBuku(*buku.CreateUpdateJenisBuku())
	if err != nil {
		restErr := errors.NewBadRequestError(fmt.Sprintf("%v", err.Error()))
		ctx.JSON(restErr.ErrStatus, restErr)
	}
	// tinggal di return aja
	ctx.JSON(http.StatusCreated, buku_response.SingleDataResponse(&data))
}

// Controller menampilkan semua buku
func (c *Controller) ShowAllBuku(ctx *gin.Context) {
	// memanggil service
	data, err := c.service.ShowAllBuku()
	// cek apakah ada yang error atau tidak
	if err != nil {
		restErr := errors.NewBadRequestError("something went wrong")
		ctx.JSON(restErr.ErrStatus, restErr)
	}
	// return json
	ctx.JSON(http.StatusOK, data)
}

// Controller menampilkan data detail buku
func (c *Controller) GetBukuById(ctx *gin.Context) {
	// ambil param id
	bId := ctx.Param("bookId")
	// sanitasi data dari user
	bukuId, err := strconv.Atoi(html.EscapeString(strings.TrimSpace(bId)))
	if err != nil {
		fmt.Println("Error during conversion")
		return
	}
	// proses ke services
	data, err := c.service.GetBuku(bukuId)
	if err != nil {
		// jika datanya error
		restErr := errors.NewBadRequestError(fmt.Sprintf("gagal mendapatkan data: %v", err.Error()))
		ctx.JSON(restErr.ErrStatus, restErr)
		return
	}
	ctx.JSON(http.StatusOK, buku_response.SingleDataResponse(data))

}

func (c *Controller) UpdateBuku(ctx *gin.Context) {
	// alokasikan memori
	// https://dev.to/bhanu011/how-is-new-different-in-go-41lk
	// https://go.dev/doc/effective_go
	buku := new(buku_request.BukuRequest)

	bId := ctx.Param("bookId")
	// sanitasi data dari user
	bukuId, err := strconv.Atoi(html.EscapeString(strings.TrimSpace(bId)))
	if err != nil {
		fmt.Println("Error during conversion")
		return
	}
	// cek json
	if err := ctx.ShouldBindJSON(&buku); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.ErrStatus, restErr)
		return
	}
	// proses ke services update buku
	data, err := c.service.UpdateBuku(bukuId, *buku.CreateUpdateJenisBuku())
	// cek jika eror
	if err != nil {
		restErr := errors.NewBadRequestError(fmt.Sprintf("invalid update data: %v", err.Error()))
		ctx.JSON(restErr.ErrStatus, restErr)
		return
	}

	// tinggal di return aja
	ctx.JSON(http.StatusCreated, buku_response.SingleDataResponse(&data))
}

func (c *Controller) HapusBuku(ctx *gin.Context) {
	bId := ctx.Param("bookId")
	// sanitasi data dari user
	bukuId, err := strconv.Atoi(html.EscapeString(strings.TrimSpace(bId)))
	if err != nil {
		fmt.Println("Error during conversion")
		return
	}
	// proses delete di services
	if err := c.service.DeleteBuku(bukuId); err != nil {
		// jika datanya error
		restErr := errors.NewBadRequestError(err.Error())
		ctx.JSON(restErr.ErrStatus, restErr)
		return
	}
	// return object any gin.H{}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Book deleted successfully",
	})
}
