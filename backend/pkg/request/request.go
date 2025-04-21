package request

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Request struct {
	validator *validator.Validate
}

type CustomRequestInterface interface {
	DoValidate(echo.Context, interface{}) interface{}
	BadRequest(c echo.Context, err interface{}) error
	Pagination() *Pagination
	Response(c echo.Context, p *Pagination, total int, result interface{}) error
}

func NewCustomRequest() *Request {
	return &Request{
		validator: validator.New(),
	}
}
