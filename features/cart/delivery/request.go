package delivery

import "tokoku/features/cart"

type AddFormat struct {
	UserID    uint `json:"id_user" form:"id_user"`
	Qty       uint `json:"qty" form:"qty"`
	ProductID uint `json:"id_product" form:"id_product"`
}

type UpdateFormat struct {
	UserID    uint `json:"id_user" form:"id_user"`
	Qty       uint `json:"qty" form:"qty"`
	ProductID uint `json:"id_product" form:"id_product"`
}

func ToCore(i interface{}) cart.Core {
	switch i.(type) {
	case AddFormat:
		cnv := i.(AddFormat)
		return cart.Core{UserID: cnv.UserID, Qty: cnv.Qty, ProductID: cnv.ProductID}
	case UpdateFormat:
		cnv := i.(UpdateFormat)
		return cart.Core{UserID: cnv.UserID, Qty: cnv.Qty, ProductID: cnv.ProductID}
	}

	return cart.Core{}
}
