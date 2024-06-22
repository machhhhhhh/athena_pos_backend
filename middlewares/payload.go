package middlewares

import (
	"athena-pos-backend/controllers"
	"athena-pos-backend/services"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

type Body struct {
	Payload string `json:"payload"`
}

func ValidateHTTPMethodGin() gin.HandlerFunc {
	return gin.HandlerFunc(func(context *gin.Context) {

		method := context.Request.Method

		if method != http.MethodPost && method != http.MethodPatch && method != http.MethodPut {
			context.Next()
			return
		}

		var body Body

		// binding JSON
		if err := context.ShouldBindJSON(&body); err != nil {
			if _, ok := err.(*json.UnmarshalTypeError); ok {
				// Handle type mismatch error
				controllers.ErrorHandlerGin(context, http.StatusBadRequest, "Invalid JSON. Payload must be a string.", "ValidateHTTPMethodGin | check type json")
				defer context.AbortWithStatus(http.StatusBadRequest)
				return
			}

			controllers.ErrorHandlerGin(context, http.StatusBadRequest, err.Error(), "ValidateHTTPMethodGin | ShouldBindJSON")
			defer context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if body.Payload == "" {
			controllers.ErrorHandlerGin(context, http.StatusBadRequest, "Empty Payload is not allowed", "ValidateHTTPMethodGin | validate payload")
			defer context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		claims, err, status_code := services.ReadTokenJWT(body.Payload)
		if err != nil {
			controllers.ErrorHandlerGin(context, status_code, err.Error(), "ValidateHTTPMethodGin | read jwt")
			defer context.AbortWithStatus(status_code)
			return
		}

		// set to context
		payload := claims["payload"]
		context.Set("body", payload)
		context.Next()
	})
}

func ValidateHTTPMethodFiber() fiber.Handler {
	return func(context *fiber.Ctx) error {
		method := context.Method()

		if method != http.MethodPost {
			return context.Next()
		}

		var body Body

		// binding JSON
		if err := context.BodyParser(&body); err != nil {
			if _, ok := err.(*json.UnmarshalTypeError); ok {
				// Handle type mismatch error
				controllers.ErrorHandlerFiber(context, http.StatusBadRequest, "Invalid JSON. Payload must be a string.", "ValidateHTTPMethodFiber | check type json")
				return context.SendStatus(http.StatusBadRequest)
			}

			controllers.ErrorHandlerFiber(context, http.StatusBadRequest, err.Error(), "ValidateHTTPMethodFiber | BodyParser")
			return context.SendStatus(http.StatusBadRequest)
		}

		if body.Payload == "" {
			controllers.ErrorHandlerFiber(context, http.StatusBadRequest, "Empty Payload is not allowed", "ValidateHTTPMethodFiber | validate payload")
			return context.SendStatus(http.StatusBadRequest)
		}

		claims, err, statusCode := services.ReadTokenJWT(body.Payload)
		if err != nil {
			controllers.ErrorHandlerFiber(context, statusCode, err.Error(), "ValidateHTTPMethodFiber | read jwt")
			return context.SendStatus(statusCode)
		}

		// set to context
		payload := claims["payload"]
		context.Locals("body", payload)
		return context.Next()
	}
}
