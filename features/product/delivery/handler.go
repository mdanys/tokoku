package delivery

import (
	"net/http"
	"strconv"
	"strings"
	"tokoku/features/product"

	"github.com/labstack/echo/v4"
)

type productHandler struct {
	srv product.Service
}

func New(e *echo.Echo, srv product.Service) {
	handler := productHandler{srv: srv}
	e.GET("/product", handler.ShowAll())
	e.GET("/product/:id_product", handler.ShowByID())
}

func (ph *productHandler) ShowAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ph.srv.ShowAll()
		if err != nil {
			if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusNotFound, FailResponse("data not found."))
			}
			return c.JSON(http.StatusInternalServerError, FailResponse("there is a problem on server."))
		}

		return c.JSON(http.StatusOK, SuccessResponse("success get all product", ToResponse(res, "all")))
	}
}

func (ph *productHandler) ShowByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		productId, err := strconv.Atoi(c.Param("id_product"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("id product must integer"))
		}

		res, err := ph.srv.ShowByID(uint(productId))
		if err != nil {
			if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusNotFound, FailResponse("data not found."))
			}
			return c.JSON(http.StatusInternalServerError, FailResponse("there is a problem on server."))
		}

		return c.JSON(http.StatusOK, SuccessResponse("success get product detail", ToResponse(res, "detail")))
	}
}
