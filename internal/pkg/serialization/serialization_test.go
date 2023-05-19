package serialization

import (
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
	e.GET("/marshal", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, msg)
	})
	
	req := httptest.NewRequest(http.MethodGet, "/marshal", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	require.Equal(t, rec.Code, http.StatusOK)

	received := Obj{}
	easyjson.Unmarshal(rec.Body.Bytes(), &received)
	require.Equal(t, received, msg)
}
