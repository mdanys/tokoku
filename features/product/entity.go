package product

import "time"

type Core struct {
	ID        uint
	Name      string
	Price     string
	Qty       string
	Detail    string
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Repository interface {
	GetAll() ([]Core, error)
	GetByID(id uint) (Core, error)
}

type Service interface {
	ShowAll() ([]Core, error)
	ShowByID(id uint) (Core, error)
}
