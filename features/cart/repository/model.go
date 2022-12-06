package repository

import (
	"tokoku/features/cart"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID    uint
	ProductID uint
	Name      string `gorm:"-:migration;<-:false"`
	Price     uint   `gorm:"-:migration;<-:false"`
	Qty       uint
	Detail    string `gorm:"-:migration;<-:false"`
	Image     string `gorm:"-:migration;<-:false"`
}

type Product struct {
	gorm.Model
	Name   string
	Price  uint
	Qty    uint
	Detail string
	Image  string
	Cart   []Cart `gorm:"foreignKey:ProductID"`
}

func FromCore(dc cart.Core) Cart {
	return Cart{
		Model:     gorm.Model{ID: dc.ID},
		UserID:    dc.UserID,
		ProductID: dc.ProductID,
		Name:      dc.Name,
		Price:     dc.Price,
		Qty:       dc.Qty,
		Detail:    dc.Detail,
		Image:     dc.Image,
	}
}

func ToCore(c Cart) cart.Core {
	return cart.Core{
		ID:        c.ID,
		UserID:    c.UserID,
		ProductID: c.ProductID,
		Name:      c.Name,
		Price:     c.Price,
		Qty:       c.Qty,
		Detail:    c.Detail,
		Image:     c.Image,
	}
}

func ToCoreArray(ca []Cart) []cart.Core {
	var arr []cart.Core
	for _, val := range ca {
		arr = append(arr, cart.Core{
			ID:        val.ID,
			UserID:    val.UserID,
			ProductID: val.ProductID,
			Name:      val.Name,
			Price:     val.Price,
			Qty:       val.Qty,
			Detail:    val.Detail,
			Image:     val.Image,
		})
	}

	return arr
}
