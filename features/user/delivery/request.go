package delivery

import "tokoku/features/user"

type RegisterFormat struct {
	Name     string `json:"name" form:"name" validate:"required,min=4"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,alpha,numeric"`
}

type UpdateFormat struct {
	Name     string `json:"name" form:"name" validate:"min=4"`
	Email    string `json:"email" form:"email" validate:"email"`
	Password string `json:"password" form:"password" validate:"alpha,numeric"`
}

type LoginFormat struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func ToCore(i interface{}) user.Core {
	switch i.(type) {
	case RegisterFormat:
		cnv := i.(RegisterFormat)
		return user.Core{Name: cnv.Name, Email: cnv.Email, Password: cnv.Password}
	case UpdateFormat:
		cnv := i.(UpdateFormat)
		return user.Core{Name: cnv.Name, Email: cnv.Email, Password: cnv.Password}
	case LoginFormat:
		cnv := i.(LoginFormat)
		return user.Core{Email: cnv.Email, Password: cnv.Password}
	}

	return user.Core{}
}
