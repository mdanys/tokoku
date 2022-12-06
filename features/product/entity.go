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
