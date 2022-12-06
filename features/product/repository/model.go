package repository

import (
	"tokoku/features/product"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name   string
	Price  uint
	Qty    uint
	Detail string
	Image  string
}

func FromCore(pc product.Core) Product {
	return Product{
		Model:  gorm.Model{ID: pc.ID, CreatedAt: pc.CreatedAt, UpdatedAt: pc.UpdatedAt},
		Name:   pc.Name,
		Price:  pc.Price,
		Qty:    pc.Qty,
		Detail: pc.Detail,
		Image:  pc.Image,
	}
}

func ToCore(p Product) product.Core {
	return product.Core{
		ID:        p.ID,
		Name:      p.Name,
		Price:     p.Price,
		Qty:       p.Qty,
		Detail:    p.Detail,
		Image:     p.Image,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func ToCoreArray(pa []Product) []product.Core {
	var arr []product.Core
	for _, val := range pa {
		arr = append(arr, product.Core{
			ID:        val.ID,
			Name:      val.Name,
			Price:     val.Price,
			Qty:       val.Qty,
			Detail:    val.Detail,
			Image:     val.Image,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
		})
	}

	return arr
}
