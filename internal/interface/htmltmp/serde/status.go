package serde

import (
	"github.com/SUT-technology/judgino/internal/domain/model"
	"github.com/SUT-technology/judgino/pkg/reqid"
	"github.com/labstack/echo/v4"
)

func Response(c echo.Context, status int, message model.UserMessage, data any) error {

	reqID, _ := reqid.RequestID(c.Request().Context())
	return c.JSON(status, model.Response{
		Msg:          message,
		Data:         data,
		TrackingCode: reqID,
	})
}
