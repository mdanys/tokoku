package repository

import (
	"tokoku/features/cart"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID    uint
	Qty       uint
	ProductID uint
	Product   Product `gorm:"foreignKey:ProductID"`
}

type Product struct {
	gorm.Model
	Name   string
	Price  uint
	Qty    uint
	Detail string
	Image  string
}

func FromCore(dc cart.Core) Cart {
	return Cart{
		Model:     gorm.Model{ID: dc.ID},
		UserID:    dc.UserID,
		Qty:       dc.Qty,
		ProductID: dc.ProductID,
		Product:   Product{Name: dc.Product.Name, Price: dc.Product.Price, Qty: dc.Product.Qty, Detail: dc.Product.Detail, Image: dc.Product.Image},
	}
}

func ToCore(c Cart) cart.Core {
	return cart.Core{
		ID:        c.ID,
		UserID:    c.UserID,
		Qty:       c.Qty,
		ProductID: c.ProductID,
		Product:   cart.ProductCore{Name: c.Product.Name, Price: c.Product.Price, Qty: c.Product.Qty, Detail: c.Product.Detail, Image: c.Product.Image},
	}
}

func ToCoreArray(ca []Cart) []cart.Core {
	var arr []cart.Core
	for _, val := range ca {
		arr = append(arr, cart.Core{
			ID:        val.ID,
			UserID:    val.UserID,
			Qty:       val.Qty,
			ProductID: val.ProductID,
			Product:   cart.ProductCore{Name: val.Product.Name, Price: val.Product.Price, Qty: val.Product.Qty, Detail: val.Product.Detail, Image: val.Product.Image},
		})
	}

	return arr
}
