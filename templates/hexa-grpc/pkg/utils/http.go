package utils

import (
	"github.com/billowdev/exclusive-go-hexa/pkg/configs"
	"github.com/gofiber/fiber/v2"
)

type APIResponse struct {
	StatusCode    string      `json:"status_code"`
	StatusMessage string      `json:"status_message"`
	Data          interface{} `json:"data"`
}

func NewSuccessResponse(c *fiber.Ctx, message string, data interface{}) error {
	if message == "" {
		message = "The process was successful"
	}
	response := APIResponse{
		StatusCode:    configs.API_SUCCESS_CODE,
		StatusMessage: message,
		Data:          data,
	}
	return c.Status(200).JSON(response)
}

func NewErrorResponse(c *fiber.Ctx, message string, data interface{}) error {
	if message == "" {
		message = "The process was failed"
	}
	response := APIResponse{
		StatusCode:    configs.API_ERROR_CODE,
		StatusMessage: message,
		Data:          data,
	}
	return c.Status(200).JSON(response)
}
