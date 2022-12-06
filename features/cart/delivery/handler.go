package delivery

import (
	"net/http"
	"os"
	"strconv"
	"tokoku/features/cart"
	"tokoku/utils/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type cartHandler struct {
	srv cart.Service
}

func New(e *echo.Echo, srv cart.Service) {
	handler := cartHandler{srv: srv}
	e.POST("/cart", handler.Create(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	e.PUT("/cart/:id", handler.Update(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	e.DELETE("/cart/:id", handler.Delete(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	e.GET("/cart", handler.Show(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
}

func (ch *cartHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input AddFormat
		userId, _ := middlewares.ExtractToken(c)

		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("an invalid client request"))
		}
		input.UserID = userId

		cnv := ToCore(input)
		res, err := ch.srv.Create(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("there is problem on server"))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("success add to cart", ToResponse(res, "cart")))
	}
}

func (ch *cartHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UpdateFormat
		id, _ := strconv.Atoi(c.Param("id"))
		userId, _ := middlewares.ExtractToken(c)

		err := c.Bind(input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("an invalid client request"))
		}
		input.UserID = userId

		cnv := ToCore(input)
		res, err := ch.srv.Update(cnv, uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("there is problem on server"))
		}

		return c.JSON(http.StatusAccepted, SuccessResponse("success update cart", ToResponse(res, "cart")))
	}
}

func (ch *cartHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		err := ch.srv.Delete(uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("there is problem on server"))
		}

		return c.JSON(http.StatusOK, SuccessResponse("success delete cart", ToResponse(err, "")))
	}
}

func (ch *cartHandler) Show() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, _ := middlewares.ExtractToken(c)
		res, err := ch.srv.Show(userId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("An invalid client request"))
		}

		return c.JSON(http.StatusOK, SuccessResponse("success get cart", ToResponse(res, "all")))
	}
}
