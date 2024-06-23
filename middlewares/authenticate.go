package middlewares

import (
	"athena-pos-backend/controllers"
	testcontroller "athena-pos-backend/controllers/test"
	"athena-pos-backend/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func AuthenticateGin() gin.HandlerFunc {
	return gin.HandlerFunc(func(context *gin.Context) {
		token_header := context.GetHeader("Authorization")

		if token_header == "" {
			controllers.ErrorHandlerGin(context, http.StatusUnauthorized, "Unauthorized.", "AuthenticateGin | validate api header")
			defer context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !strings.Contains(token_header, "Bearer") {
			controllers.ErrorHandlerGin(context, http.StatusUnauthorized, "Please use Authorization Bearer.", "AuthenticateGin | validate api header")
			defer context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		access_token := strings.TrimSpace(strings.Replace(token_header, "Bearer", "", 1))

		aes_user, err := services.AESDecrypted(access_token)
		if err != nil {
			controllers.ErrorHandlerGin(context, http.StatusUnauthorized, err.Error(), "AuthenticateGin | decrypt token")
			defer context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		user, err := testcontroller.FindUser(aes_user.UserID, "")
		if err != nil {
			controllers.ErrorHandlerGin(context, http.StatusUnauthorized, err.Error(), "AuthenticateGin | find user")
			defer context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// TODO: hide password => if connect DB query to omit it
		user.Password = ""

		context.Set("user", user)
		context.Next()
	})
}

func AuthenticateFiber() fiber.Handler {
	return func(context *fiber.Ctx) error {
		token_header := context.Get("Authorization")

		if token_header == "" {
			return controllers.ErrorHandlerFiber(context, http.StatusUnauthorized, "Unauthorized.", "AuthenticateFiber | validate api header")
		}

		if !strings.Contains(token_header, "Bearer") {
			return controllers.ErrorHandlerFiber(context, http.StatusUnauthorized, "Please use Authorization Bearer.", "AuthenticateFiber | validate api header")
		}

		access_token := strings.TrimSpace(strings.Replace(token_header, "Bearer", "", 1))

		aes_user, err := services.AESDecrypted(access_token)
		if err != nil {
			return controllers.ErrorHandlerFiber(context, http.StatusUnauthorized, err.Error(), "AuthenticateFiber | decrypt token")
		}

		user, err := testcontroller.FindUser(aes_user.UserID, "")
		if err != nil {
			return controllers.ErrorHandlerFiber(context, http.StatusUnauthorized, err.Error(), "AuthenticateFiber | find user")
		}
		// TODO: hide password => if connect DB query to omit it
		user.Password = ""

		context.Locals("user", user)
		return context.Next()
	}
}
