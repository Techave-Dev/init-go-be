package tools

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guregu/null/v5"
)

type ResponseStatus = string

const (
	ResponseStatusSuccess ResponseStatus = "success"
	ResponseStatusFail    ResponseStatus = "fail"
	ResponseStatusError   ResponseStatus = "error"
)

type ResponseBody struct {
	Status  ResponseStatus `json:"status"`
	Message null.String    `json:"message"`
	Data    any            `json:"data"`
}

type R = map[string]any

func Null() null.String             { return null.StringFromPtr(nil) }
func String(str string) null.String { return null.StringFrom(str) }

func Success(c *fiber.Ctx, code int, message null.String, data R) error {
	return c.Status(code).JSON(ResponseBody{
		Status:  ResponseStatusSuccess,
		Message: message,
		Data:    data,
	})
}

func Fail(c *fiber.Ctx, code int, message null.String, data ...string) error {
	return c.Status(code).JSON(ResponseBody{
		Status:  ResponseStatusFail,
		Message: message,
		Data:    data,
	})
}

func Error(c *fiber.Ctx, code int, message null.String) error {
	return c.Status(code).JSON(ResponseBody{
		Status:  ResponseStatusError,
		Message: message,
	})
}
