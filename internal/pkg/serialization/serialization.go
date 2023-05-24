package serialization

import (
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/mailru/easyjson"
	"io"
)

type EasyJsonSerializer struct{}

func (ejs EasyJsonSerializer) Serialize(ctx echo.Context, data interface{}, indent string) error {
	switch data.(type) {
	case echo.Map:
		dataMap := data.(echo.Map)
		//result := make(map[string]interface{})
		//for k, v := range dataMap {
		//	result[k.(string)] = v
		//}

		bytes, err := json.Marshal(dataMap)
		if err != nil {
			return err
		}

		return ctx.JSONBlob(ctx.Response().Status, bytes)
		//return json.Marshal(result)
	}

	marshable := data.(easyjson.Marshaler)
	blob, err := easyjson.Marshal(marshable)
	if err != nil {
		return err
	}

	return ctx.JSONBlob(ctx.Response().Status, blob)
}

func (s EasyJsonSerializer) Deserialize(ctx echo.Context, i interface{}) error {
	data, err := io.ReadAll(ctx.Request().Body)
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

	switch i.(type) {
	case echo.Map:
		//dataMap := data.(echo.Map)
		//result := make(map[string]interface{})
		//for k, v := range dataMap {
		//	result[k.(string)] = v
		//}

		//bytes, err := json.Marshal(dataMap)
		//if err != nil {
		//	return err
		//}
		//
		//return ctx.JSONBlob(ctx.Response().Status, bytes)
		////return json.Marshal(result)

		err := json.Unmarshal(data, i.(echo.Map))
		if err != nil {
			return err
		}
	}

	err := easyjson.Unmarshal(data, i.(easyjson.Unmarshaler))
	if err != nil {
		return err
	}

	return nil
}
