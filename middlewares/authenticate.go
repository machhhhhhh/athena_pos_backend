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

		user, err := testcontroller.FindUserByID(aes_user.UserID)
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
		tokenHeader := context.Get("Authorization")

		if tokenHeader == "" {
			controllers.ErrorHandlerFiber(context, http.StatusUnauthorized, "Unauthorized.", "AuthenticateFiber | validate api header")
			return context.SendStatus(http.StatusUnauthorized)
		}

		if !strings.Contains(tokenHeader, "Bearer") {
			controllers.ErrorHandlerFiber(context, http.StatusUnauthorized, "Please use Authorization Bearer.", "AuthenticateFiber | validate api header")
			return context.SendStatus(http.StatusUnauthorized)
		}

		accessToken := strings.TrimSpace(strings.Replace(tokenHeader, "Bearer", "", 1))

		aesUser, err := services.AESDecrypted(accessToken)
		if err != nil {
			controllers.ErrorHandlerFiber(context, http.StatusUnauthorized, err.Error(), "AuthenticateFiber | decrypt token")
			return context.SendStatus(http.StatusUnauthorized)
		}

		user, err := testcontroller.FindUserByID(aesUser.UserID)
		if err != nil {
			controllers.ErrorHandlerFiber(context, http.StatusUnauthorized, err.Error(), "AuthenticateFiber | find user")
			return context.SendStatus(http.StatusUnauthorized)
		}
		// TODO: hide password => if connect DB query to omit it
		user.Password = ""

		context.Locals("user", user)
		return context.Next()
	}
}
