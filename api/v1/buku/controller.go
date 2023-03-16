package buku

import (
	"fmt"
	"net/http"

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
