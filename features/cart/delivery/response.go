package delivery

import "tokoku/features/cart"

type CartResponse struct {
	ID        uint             `json:"id_cart" form:"id_cart"`
	UserID    uint             `json:"id_user" form:"id_user"`
	Qty       uint             `json:"qty" form:"qty"`
	ProductID uint             `json:"id_product" form:"id_product"`
	Product   cart.ProductCore `json:"-" form:"-"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "cart":
		cnv := core.(cart.Core)
		res = CartResponse{ID: cnv.ID, UserID: cnv.UserID, Qty: cnv.Qty, ProductID: cnv.ProductID}
	case "all":
		var arr []CartResponse
		cnv := core.([]cart.Core)
		for _, val := range cnv {
			arr = append(arr, CartResponse{
				ID:        val.ID,
				UserID:    val.UserID,
				Qty:       val.Qty,
				ProductID: val.ProductID,
			})
		}
		res = arr
	}

	return res
}

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func FailResponse(msg interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}
