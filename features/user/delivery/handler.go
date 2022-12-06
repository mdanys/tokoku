package delivery

import (
	"net/http"
	"strings"
	"tokoku/features/user"
	"tokoku/utils/middlewares"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var validate = validator.New()

type userHandler struct {
	srv user.Service
}

func New(e *echo.Echo, srv user.Service) {
	handler := userHandler{srv: srv}
}

func (uh *userHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RegisterFormat
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}

		er := validate.Struct(input)
		if er != nil {
			if strings.Contains(er.Error(), "min") {
				return c.JSON(http.StatusBadRequest, FailResponse("min. 4 character"))
			} else if strings.Contains(er.Error(), "email") {
				return c.JSON(http.StatusBadRequest, FailResponse("must input valid email"))
			} else if strings.Contains(er.Error(), "password") {
				return c.JSON(http.StatusBadRequest, FailResponse("password must include alphabet and numeric"))
			}
			return c.JSON(http.StatusBadRequest, FailResponse(er.Error()))
		}

		cnv := ToCore(input)
		res, err := uh.srv.Create(cnv)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate") {
				return c.JSON(http.StatusBadRequest, FailResponse("duplicate email on database"))
			} else if strings.Contains(err.Error(), "password") {
				return c.JSON(http.StatusBadRequest, FailResponse("cannot encrypt password"))
			}
			return c.JSON(http.StatusInternalServerError, FailResponse("there is problem on server."))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("success create user", ToResponse(res, "register")))
	}
}

func (uh *userHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, role := middlewares.ExtractToken(c)
		if userId == 0 {
			return c.JSON(http.StatusUnauthorized, FailResponse("cannot validate token"))
		} else if role == "" {
			return c.JSON(http.StatusUnauthorized, FailResponse("cannot recognize role"))
		} else {
			var input UpdateFormat
			err := c.Bind(&input)
			if err != nil {
				return c.JSON(http.StatusBadRequest, FailResponse("cannot bind update data"))
			}

			er := validate.Struct(input)
			if er != nil {
				if strings.Contains(er.Error(), "min") {
					return c.JSON(http.StatusBadRequest, FailResponse("min. 4 character"))
				} else if strings.Contains(er.Error(), "email") {
					return c.JSON(http.StatusBadRequest, FailResponse("must input valid email"))
				} else if strings.Contains(er.Error(), "password") {
					return c.JSON(http.StatusBadRequest, FailResponse("password must include alphabet and numeric"))
				}
				return c.JSON(http.StatusBadRequest, FailResponse(er.Error()))
			}

			cnv := ToCore(input)
			res, err := uh.srv.Update(cnv, userId)
			if err != nil {
				if strings.Contains(err.Error(), "found") {
					return c.JSON(http.StatusNotFound, FailResponse("data not found."))
				}
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}

			return c.JSON(http.StatusAccepted, SuccessResponse("success update user", ToResponse(res, "update")))
		}
	}
}

func (uh *userHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, role := middlewares.ExtractToken(c)
		if userId == 0 {
			return c.JSON(http.StatusUnauthorized, FailResponse("cannot validate token"))
		} else if role == "" {
			return c.JSON(http.StatusUnauthorized, FailResponse("cannot recognize role"))
		} else {
			err := uh.srv.Delete(userId)
			if err != nil {
				if strings.Contains(err.Error(), "found") {
					return c.JSON(http.StatusNotFound, FailResponse("data not found."))
				}
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}

			return c.JSON(http.StatusOK, SuccessResponse("success delete data", ToResponse(err, "")))
		}
	}
}

func (uh *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginFormat
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}

		cnv := ToCore(input)
		res, err := uh.srv.Login(cnv)
		if err != nil {
			if strings.Contains(err.Error(), "an invalid client request") {
				return c.JSON(http.StatusBadRequest, FailResponse("email doesn't exist."))
			} else if strings.Contains(err.Error(), "password not match") {
				return c.JSON(http.StatusBadRequest, FailResponse("password not match."))
			}
			return c.JSON(http.StatusInternalServerError, FailResponse("there is problem on server"))
		}

		return c.JSON(http.StatusOK, SuccessResponse("success login", ToResponse(res, "login")))
	}
}
