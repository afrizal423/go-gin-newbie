package configs

import (
	"context"

	"github.com/afrizal423/go-gin-newbie/app/models"
)

var ctx context.Context

func InitData() []models.Buku {
	var buku []models.Buku
	return buku
}

func InitContext() context.Context {
	ctx = context.Background()
	return ctx
}
