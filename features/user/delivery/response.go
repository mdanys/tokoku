package delivery

import "tokoku/features/user"

type RegisterResponse struct {
	ID    uint   `json:"id_user" form:"id_user"`
	Name  string `json:"name" form:"name"`
	Role  string `json:"role" form:"role"`
	Email string `json:"email" form:"email"`
}

type UpdateResponse struct {
	ID    uint   `json:"id_user" form:"id_user"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

type LoginResponse struct {
	ID    uint   `json:"id_user" form:"id_user"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Role  string `json:"role" form:"role"`
	Token string `json:"token" form:"token"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "register":
		cnv := core.(user.Core)
		res = RegisterResponse{ID: cnv.ID, Name: cnv.Name, Role: cnv.Role, Email: cnv.Email}
	case "update":
		cnv := core.(user.Core)
		res = UpdateResponse{ID: cnv.ID, Name: cnv.Name, Email: cnv.Email}
	case "login":
		cnv := core.(user.Core)
		res = LoginResponse{ID: cnv.ID, Name: cnv.Name, Email: cnv.Email, Role: cnv.Role, Token: cnv.Token}
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
