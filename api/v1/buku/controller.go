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

// AddBook godoc
// @Summary Add book details
// @Description Add book details
// @Tags books
// @Accept json
// @Produce json
// @Param buku_request.BukuRequest body buku_request.BukuRequest true "Add the book"
// @Success 200 {object} buku_response.BukuResponse
// @Failure 400 {object} errors.RestErr
// @Failure 500 {string} string "Internal Server Error"
// @Router /books [post]
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

// GetAllBook godoc
// @Summary Get details
// @Description Get details of all book
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {array} buku_response.BukuResponse
// @Failure 404 {object} errors.RestErr
// @Router /books [get]
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

// GetBook godoc
// @Summary Get details by id
// @Description Get details of book by id
// @Tags books
// @Accept json
// @Produce json
// @Param bookId path int true "ID of the book"
// @Success 200 {object} buku_response.BukuResponse
// @Failure 404 {object} errors.RestErr
// @Router /books/{bookId} [get]
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
		// jika datanya tidak ada
		restErr := errors.NewNotFoundError(fmt.Sprintf("gagal mendapatkan data: %v", err.Error()))
		ctx.JSON(restErr.ErrStatus, restErr)
		return
	}
	ctx.JSON(http.StatusOK, buku_response.SingleDataResponse(data))

}

// UpdateBook godoc
// @Summary Update of the book by id
// @Description Update of the book by id
// @Tags books
// @Accept json
// @Produce json
// @Param bookId path int true "ID of the book"
// @Param buku_request.BukuRequest body buku_request.BukuRequest true "Update the book"
// @Success 200 {object} buku_response.BukuResponse
// @Failure 400 {object} errors.RestErr
// @Failure 404 {object} errors.RestErr
// @Failure 500 {string} string "Internal Server Error"
// @Router /books/{bookId} [put]
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
		restErr := errors.NewNotFoundError(fmt.Sprintf("invalid update data: %v", err.Error()))
		ctx.JSON(restErr.ErrStatus, restErr)
		return
	}

	// tinggal di return aja
	ctx.JSON(http.StatusCreated, buku_response.SingleDataResponse(&data))
}

// DeleteBook godoc
// @Summary Delete by id
// @Description Delete of book by id
// @Tags books
// @Accept json
// @Produce json
// @Param bookId path int true "ID of the book"
// @Success 200 {object} buku_response.Deleted
// @Failure 404 {object} buku_response.NotFound
// @Router /books/{bookId} [delete]
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
