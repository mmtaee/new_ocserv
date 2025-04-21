package request

import (
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type ErrorResponse struct {
	Error string `json:"error" validate:"required"`
}

func (r *Request) BadRequest(c echo.Context, err interface{}, msg ...string) error {
	var response struct {
		Error   string
		Message []string
	}

	switch err.(type) {
	case error:
		var pqErr *pgconn.PgError
		if errors.As(err.(error), &pqErr) {
			response.Error = err.(*pgconn.PgError).Detail
		} else {
			response.Error = err.(error).Error()
		}
	case string:
		response.Error = err.(string)
	case map[string]interface{}:
		errs := err.(map[string]interface{})["error"]
		if errSlice, ok := errs.([]string); ok && len(errSlice) > 0 {
			response.Error = strings.Join(errSlice, ", ")
		}
	default:
		response.Error = "unknown error"
	}
	response.Message = append(response.Message, msg...)
	return c.JSON(http.StatusBadRequest, response)
}
