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
	return gin.HandlerFunc(func(c *gin.Context) {
		token_header := c.GetHeader("Authorization")

		if token_header == "" {
			controllers.ErrorHandlerGin(c, http.StatusUnauthorized, "Unauthorized.", "AuthenticateGin | validate api header")
			defer c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !strings.Contains(token_header, "Bearer") {
			controllers.ErrorHandlerGin(c, http.StatusUnauthorized, "Please use Authorization Bearer.", "AuthenticateGin | validate api header")
			defer c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		access_token := strings.TrimSpace(strings.Replace(token_header, "Bearer", "", 1))

		aes_user, err := services.AESDecrypted(access_token)
		if err != nil {
			controllers.ErrorHandlerGin(c, http.StatusUnauthorized, err.Error(), "AuthenticateGin | decrypt token")
			defer c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		user, err := testcontroller.FindUserByID(aes_user.UserID)
		if err != nil {
			controllers.ErrorHandlerGin(c, http.StatusUnauthorized, err.Error(), "AuthenticateGin | find user")
			defer c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// TODO: hide password => if connect DB query to omit it
		user.Password = ""

		c.Set("user", user)
		c.Next()
	})
}

func AuthenticateFiber() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenHeader := c.Get("Authorization")

		if tokenHeader == "" {
			controllers.ErrorHandlerFiber(c, http.StatusUnauthorized, "Unauthorized.", "AuthenticateFiber | validate api header")
			return c.SendStatus(http.StatusUnauthorized)
		}

		if !strings.Contains(tokenHeader, "Bearer") {
			controllers.ErrorHandlerFiber(c, http.StatusUnauthorized, "Please use Authorization Bearer.", "AuthenticateFiber | validate api header")
			return c.SendStatus(http.StatusUnauthorized)
		}

		accessToken := strings.TrimSpace(strings.Replace(tokenHeader, "Bearer", "", 1))

		aesUser, err := services.AESDecrypted(accessToken)
		if err != nil {
			controllers.ErrorHandlerFiber(c, http.StatusUnauthorized, err.Error(), "AuthenticateFiber | decrypt token")
			return c.SendStatus(http.StatusUnauthorized)
		}

		user, err := testcontroller.FindUserByID(aesUser.UserID)
		if err != nil {
			controllers.ErrorHandlerFiber(c, http.StatusUnauthorized, err.Error(), "AuthenticateFiber | find user")
			return c.SendStatus(http.StatusUnauthorized)
		}
		// TODO: hide password => if connect DB query to omit it
		user.Password = ""

		c.Locals("user", user)
		return c.Next()
	}
}
