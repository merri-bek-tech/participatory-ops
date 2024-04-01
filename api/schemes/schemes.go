package schemes

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// These are hard coded for now as this tools is a prototype being co-developed with the Merri-bek tech
// organisation. Eventually these will be stored somewhere and configurable. It is the intent that an
// installation of parops could support many organisations, each with their own schemes, as a place for
// them to start out, and then they could move their data onto a self-hosted instance when ready.
var schemes = []SchemeIdentity{
	{"Merri-bek Tech", "mbt", []string{"parops.merri-bek.tech"}},
	{"Merri-bek Tech Dev", "mbt-dev", []string{"dev.parops.merri-bek.tech", "localhost"}},
}

type SchemeIdentity struct {
	Name      string   `json:"name"`
	Id        string   `json:"id"`
	Hostnames []string `json:"hostnames"`
}

type HandlerFuncWithScheme func(c echo.Context, scheme *SchemeIdentity) error

func GetIndex(c echo.Context) error {
	return c.JSON(http.StatusOK, schemes)
}

func GetScheme(id string) (scheme *SchemeIdentity, exists bool) {
	for _, scheme := range schemes {
		if scheme.Id == id {
			return &scheme, true
		}
	}
	return nil, false
}

func WithScheme(next HandlerFuncWithScheme) echo.HandlerFunc {
	return func(c echo.Context) error {
		schemeId := c.Param("schemeId")
		if schemeId == "" {
			log.Println("schemeID not found")
			return c.String(http.StatusBadRequest, "schemeId is required")
		}

		scheme, exists := GetScheme(schemeId)
		if !exists {
			log.Println("scheme not found", schemeId)
			return c.String(http.StatusNotFound, "scheme not found")
		}

		return next(c, scheme)
	}
}
