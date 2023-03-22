package messages

import (
	"github.com/labstack/echo/v4"
)

type Usecase interface {
	SendMessage(ctx echo.Context, jsonSentMessage []byte) error
	ReceiveMessage(ctx echo.Context) ([]byte, error)
}
