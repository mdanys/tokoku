package delivery

import "tokoku/features/product"

type ProductResponse struct {
	ID    uint   `json:"id_product" form:"id_product"`
	Name  string `json:"name" form:"name"`
	Price uint   `json:"price" form:"price"`
	Qty   uint   `json:"qty" form:"qty"`
	Image string `json:"image" form:"image"`
}

type DetailResponse struct {
	ID     uint   `json:"id_product" form:"id_product"`
	Name   string `json:"name" form:"name"`
	Price  uint   `json:"price" form:"price"`
	Qty    uint   `json:"qty" form:"qty"`
	Detail string `json:"detail" form:"detail"`
	Image  string `json:"image" form:"image"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "all":
		var arr []ProductResponse
		cnv := core.([]product.Core)
		for _, val := range cnv {
			arr = append(arr, ProductResponse{
				ID:    val.ID,
				Name:  val.Name,
				Price: val.Price,
				Qty:   val.Qty,
				Image: val.Image,
			})
		}
		res = arr
	case "detail":
		cnv := core.(product.Core)
		res = DetailResponse{ID: cnv.ID, Name: cnv.Name, Price: cnv.Price, Qty: cnv.Qty, Detail: cnv.Detail, Image: cnv.Image}
	}

	return res
}

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func FailResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}
