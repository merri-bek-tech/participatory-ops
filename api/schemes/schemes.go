package schemes

import (
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
	Name    string   `json:"name"`
	Id      string   `json:"id"`
	Domains []string `json:"domains"`
}

func GetIndex(c echo.Context) error {
	return c.JSON(http.StatusOK, schemes)
}
