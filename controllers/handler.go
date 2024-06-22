package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

type RequestPayload struct {
	Payload string `json:"payload,omitempty"`
} // @name RequestPayload

type SuccessResponse struct {
	Message  string      `json:"message,omitempty"`
	Response interface{} `json:"response,omitempty"`
} // @name SuccessResponse

type ErrorResponse struct {
	Error        string `json:"error,omitempty"`
	ErrorSection string `json:"error_section,omitempty"`
} // @name ErrorResponse

func ErrorHandlerFiber(c *fiber.Ctx, error_status int, error_message string, error_section string) {
	c.Status(error_status).JSON(ErrorResponse{
		Error:        error_message,
		ErrorSection: error_section,
	})
}
func ErrorHandlerGin(context *gin.Context, error_status int, error_message string, error_section string) {
	context.JSON(error_status, ErrorResponse{
		Error:        error_message,
		ErrorSection: error_section,
	})
}

// func PreloadUnscopedHandler(db *gorm.DB) *gorm.DB {
// 	return db.Unscoped()
// }
