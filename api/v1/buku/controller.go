package buku

import (
	"fmt"
	"net/http"

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
	data, err := c.service.CreateBuku(*buku.CreateJenisBuku())
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

// func (c *Controller) GetBukuById(ctx *gin.Context) {
// 	bId := ctx.Param("bookId")
// 	bukuId, err := strconv.Atoi(html.EscapeString(strings.TrimSpace(bId)))
// 	if err != nil {
// 		fmt.Println("Error during conversion")
// 		return
// 	}

// 	data, err := c.service.GetBuku(bukuId)
// 	if err != nil {
// 		// jika datanya error
// 		restErr := errors.NewBadRequestError(fmt.Sprintf("gagal mendapatkan data: %v", err.Error()))
// 		ctx.JSON(restErr.ErrStatus, restErr)
// 	} else {
// 		ctx.JSON(http.StatusOK, data)
// 	}

// }

// func (c *Controller) UpdateBuku(ctx *gin.Context) {
// 	var buku models.Buku
// 	bId := ctx.Param("bookId")
// 	bukuId, err := strconv.Atoi(html.EscapeString(strings.TrimSpace(bId)))
// 	if err != nil {
// 		fmt.Println("Error during conversion")
// 		return
// 	}
// 	if err := ctx.ShouldBindJSON(&buku); err != nil {
// 		restErr := errors.NewBadRequestError("invalid json body")
// 		ctx.JSON(restErr.ErrStatus, restErr)
// 		return
// 	}
// 	fmt.Println(buku, bukuId)

// 	if err := c.service.UpdateBuku(bukuId, buku); err != nil {
// 		restErr := errors.NewBadRequestError(fmt.Sprintf("invalid update data: %v", err.Error()))
// 		ctx.JSON(restErr.ErrStatus, restErr)
// 		return
// 	}

// 	ctx.JSON(http.StatusCreated, "Updated")
// }

// func (c *Controller) HapusBuku(ctx *gin.Context) {
// 	bId := ctx.Param("bookId")
// 	bukuId, err := strconv.Atoi(bId)
// 	if err != nil {
// 		fmt.Println("Error during conversion")
// 		return
// 	}
// 	if err := c.service.DeleteBuku(bukuId); err != nil {
// 		// jika datanya error
// 		restErr := errors.NewBadRequestError(err.Error())
// 		ctx.JSON(restErr.ErrStatus, restErr)
// 		return
// 	}
// 	ctx.JSON(http.StatusCreated, "Deleted")
// }
