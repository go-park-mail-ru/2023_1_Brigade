package serialization

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mailru/easyjson"
)

type EasyJsonSerializer struct{}

func (ejs EasyJsonSerializer) Serialize(ctx echo.Context, data interface{}, indent string) error {
	marshable := data.(easyjson.Marshaler)
	blob, err := easyjson.Marshal(marshable)
	if err != nil {
		return err
	}

	return ctx.JSONBlob(http.StatusOK, blob)
}

func (s EasyJsonSerializer) Deserialize(ctx echo.Context, i interface{}) error {
	data, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return err
	}

	return s.deserializeBytes(data, i)
}

func (s EasyJsonSerializer) deserializeBytes(data []byte, i interface{}) error {
	if data == nil {
		return errors.New("nil data")
	}
	if i == nil {
		return errors.New("nil i")
	}

	err := easyjson.Unmarshal(data, i.(easyjson.Unmarshaler))
	if err != nil {
		return err
	}

	return nil
}
