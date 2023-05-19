package serialization

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mailru/easyjson"
)

type EasyJsonSerializer struct{}

func (ejs EasyJsonSerializer) Serialize(ctx echo.Context, data interface{}, indent string) error {
	suitable := data.(easyjson.Marshaler)
	blob, err := easyjson.Marshal(suitable)
	if err != nil {
		return err
	}
	return ctx.JSONBlob(http.StatusOK, blob)
}

func (ejs EasyJsonSerializer) Deserialize(c echo.Context, i interface{}) error {
	return nil
}
