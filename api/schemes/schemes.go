package schemes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetIndex(c echo.Context) error {
	return c.JSON(http.StatusOK, true)
}
