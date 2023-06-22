package http

import (
	"bytes"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	userMock "project/internal/microservices/user/usecase/mocks"
	imagesMock "project/internal/monolithic_services/images/usecase/mocks"
	myErrors "project/internal/pkg/errors"
	"project/internal/pkg/http_utils"
	"testing"
)

func TestHandlers_UploadUserAvatarsHandler_Error(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	e := echo.New()
	r := httptest.NewRequest("POST", "/signup/", bytes.NewReader(nil))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	ctx := e.NewContext(r, w)

	userUsecase := userMock.NewMockUsecase(ctl)
	imagesUsecase := imagesMock.NewMockUsecase(ctl)
	handler := NewImagesHandler(e, userUsecase, imagesUsecase)

	err := handler.UploadUserAvatarsHandler(ctx)
	require.Equal(t, http_utils.StatusCode(err), http_utils.StatusCode(myErrors.ErrInternal))
}

func TestHandlers_UploadChatAvatarsHandler_Error(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	e := echo.New()
	r := httptest.NewRequest("POST", "/signup/", bytes.NewReader(nil))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	ctx := e.NewContext(r, w)

	userUsecase := userMock.NewMockUsecase(ctl)
	imagesUsecase := imagesMock.NewMockUsecase(ctl)
	handler := NewImagesHandler(e, userUsecase, imagesUsecase)

	err := handler.UploadChatAvatarsHandler(ctx)
	require.Equal(t, http_utils.StatusCode(err), http_utils.StatusCode(myErrors.ErrInternal))
}

func TestHandlers_UploadChatImagesHandler_Error(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	e := echo.New()
	r := httptest.NewRequest("POST", "/signup/", bytes.NewReader(nil))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	ctx := e.NewContext(r, w)

	userUsecase := userMock.NewMockUsecase(ctl)
	imagesUsecase := imagesMock.NewMockUsecase(ctl)
	handler := NewImagesHandler(e, userUsecase, imagesUsecase)

	err := handler.UploadChatImagesHandler(ctx)
	require.Equal(t, http_utils.StatusCode(err), http_utils.StatusCode(myErrors.ErrInternal))
}
