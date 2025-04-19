package serde

import "github.com/labstack/echo/v4"

func BindRequestBody[T any](c echo.Context) (T, error) {
	t := new(T)
	if err := c.Bind(t); err != nil {
		return *t, err
	}
	if err := c.Validate(t); err != nil {
		return *t, err
	}
	return *t, nil
}

// func BindRequestParams[T any](c echo.Context) (T, error) {
// 	t := new(T)
// 	if err := c.QueryParamsBinder(t).BindQueryParams(); err != nil {
// 		return *t, err
// 	}
// 	if err := c.Validate(t); err != nil {
// 		return *t, err
// 	}
// 	return *t, nil
// }
