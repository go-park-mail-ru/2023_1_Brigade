package serialization

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/mailru/easyjson"
	"github.com/stretchr/testify/require"
)

func TestMarshal(t *testing.T) {
	e := echo.New()
	e.JSONSerializer = EasyJsonSerializer{}

	msg := Obj{Str: "ok"}
	e.GET("", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, msg)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	require.Equal(t, rec.Code, http.StatusOK)

	received := Obj{}
	easyjson.Unmarshal(rec.Body.Bytes(), &received)
	require.Equal(t, received, msg)
}

func TestUnmarshal(t *testing.T) {
	e := echo.New()
	e.JSONSerializer = EasyJsonSerializer{}

	msg := Obj{Str: "ok"}

	e.POST("", func(ctx echo.Context) error {
		received := Obj{}
		if err := ctx.Bind(&received); err != nil {
			return err
		}
		require.Equal(t, received, msg)
		return ctx.JSON(http.StatusOK, received)
	})

	payload, err := easyjson.Marshal(msg)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)
	require.Equal(t, http.StatusOK, rec.Code)

	receivedResp := Obj{}
	err = easyjson.Unmarshal(rec.Body.Bytes(), &receivedResp)
	require.NoError(t, err)
	require.Equal(t, receivedResp, msg)
}
