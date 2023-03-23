package buku

type IBukuService interface {
	// CreateBuku(data models.Buku) error
	// ShowAllBuku() ([]models.Buku, error)
	// GetBuku(id int) (*models.Buku, error)
	// UpdateBuku(id int, data models.Buku) error
	// DeleteBuku(id int) error
}

type IBukuRepository interface {
	// CreateBukus(data models.Buku, db *sql.DB) error
	// ShowAllBukus(db *sql.DB) ([]models.Buku, error)
	// GetBukus(id int, db *sql.DB) (*models.Buku, error)
	// UpdateBukus(id int, data models.Buku, db *sql.DB) error
	// DeleteBukus(id int, db *sql.DB) error
	// HitungBukus(id int, db *sql.DB) error
}
