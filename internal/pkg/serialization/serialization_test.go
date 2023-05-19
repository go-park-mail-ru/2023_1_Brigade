package serialization

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestMarshal(t *testing.T) {
	e := echo.New()
	e.JSONSerializer = EasyJsonSerializer{}
	e.GET("/marshal", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, Obj{Str: "ok"})
	})
	req := httptest.NewRequest(http.MethodGet, "/marshal", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	require.Equal(t, rec.Code, http.StatusOK)
	require.Equal(t, rec.Body.String(), "{\"str\":\"ok\"}")
}
