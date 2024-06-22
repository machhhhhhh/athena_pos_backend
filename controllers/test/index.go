package testcontroller

import (
	"athena-pos-backend/controllers"
	"athena-pos-backend/models"
	"athena-pos-backend/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/matthewhartstonge/argon2"
)

var USER_1 = models.User{
	UserID:    1,
	FirstName: "test",
	LastName:  "test",
	Username:  "test",
	Password:  "test",
}
var USER_2 = models.User{
	UserID:    2,
	FirstName: "test",
	LastName:  "test",
	Username:  "test",
	Password:  "test",
}

var ALL_USER []models.User = []models.User{USER_1, USER_2}

// Test Service Godoc
// @Summary Test Controller.
// @Description test controller.
// @ID TestGetFiber
// @Accept json
// @Produce json
// @Param search query string false "string valid"
// @Param limit query int false "int valid"
// @Param page  query int false "int valid"
// @Router /test [get]
// @Tags Test
// @Security ApiKeyAuth
// @Success 200 {array} models.User "Successfully Retrieved Data"
// @Failure 400 {object} controllers.ErrorResponse "Bad Request"
// @Failure 401 {object} controllers.ErrorResponse "Unauthorized"
// @Failure 403 {object} controllers.ErrorResponse "Role Forbidden"
// @Failure 404 {object} controllers.ErrorResponse "Not Found"
// @Failure 409 {object} controllers.ErrorResponse "Conflict"
// @Failure 429 {object} controllers.ErrorResponse "Rate Limit Exceeded"
// @Failure 500 {object} controllers.ErrorResponse "Internal Server Error"
// @Failure 502 {object} controllers.ErrorResponse "Bad Gateway"
func TestGetFiber(context *fiber.Ctx) error {

	userMe, err := services.GetContextUserFiber(context)
	if err != nil {
		return controllers.ErrorHandlerFiber(context, http.StatusInternalServerError, err.Error(), "TestGetFiber | get user from token")
	}

	fmt.Println("===========")
	fmt.Println("userMe", userMe)
	fmt.Println("===========")

	user := GetUser()
	return context.Status(fiber.StatusOK).JSON(controllers.SuccessResponse{
		Message:  "Successfully Retrieved User",
		Response: user,
	})
}

func TestCreateFiber(context *fiber.Ctx) error {

	userMe, err := services.GetContextUserFiber(context)
	if err != nil {
		return controllers.ErrorHandlerFiber(context, http.StatusInternalServerError, err.Error(), "TestCreateFiber | get user from token")
	}

	fmt.Println("===========")
	fmt.Println("userMe", userMe)
	fmt.Println("===========")

	body := models.User{
		FirstName: "test",
		LastName:  "test",
		Username:  "test",
		Password:  "test",
	}

	body = CreateUser(body)

	return context.Status(fiber.StatusCreated).JSON(controllers.SuccessResponse{
		Message:  "Successfully Create User",
		Response: body,
	})
}
func TestGetGin(context *gin.Context) {

	userMe, err := services.GetContextUserGin(context)
	if err != nil {
		controllers.ErrorHandlerGin(context, http.StatusInternalServerError, err.Error(), "TestGetGin | get user from token")
		return
	}

	fmt.Println("===========")
	fmt.Println("userMe", userMe)
	fmt.Println("===========")

	user := GetUser()

	context.JSON(http.StatusOK, controllers.SuccessResponse{
		Message:  "Successfully Retrieved User",
		Response: user,
	})
}
func TestCreateGin(context *gin.Context) {

	userMe, err := services.GetContextUserGin(context)
	if err != nil {
		controllers.ErrorHandlerGin(context, http.StatusInternalServerError, err.Error(), "TestGetGin | get user from token")
		return
	}

	fmt.Println("===========")
	fmt.Println("userMe", userMe)
	fmt.Println("===========")

	body := models.User{
		FirstName: "test",
		LastName:  "test",
		Username:  "test",
		Password:  "test",
	}

	body = CreateUser(body)

	context.JSON(http.StatusCreated, controllers.SuccessResponse{
		Message:  "Successfully Create User",
		Response: body,
	})
}

func GetPayloadLoginFiber(context *fiber.Ctx) error {

	body := LoginRequest{
		Username: "test",
		Password: "test",
	}

	token, err := services.GenerateTokenJWT(body)
	if err != nil {
		return controllers.ErrorHandlerFiber(context, http.StatusInternalServerError, err.Error(), "GetPayload | generate token")
	}

	return context.Status(fiber.StatusOK).JSON(controllers.SuccessResponse{
		Message:  "Get Payload Login",
		Response: token,
	})
}

func GetPayloadLoginGin(context *gin.Context) {

	body := LoginRequest{
		Username: "test",
		Password: "test",
	}

	token, err := services.GenerateTokenJWT(body)
	if err != nil {
		controllers.ErrorHandlerGin(context, http.StatusInternalServerError, err.Error(), "GetPayload | generate token")
		return
	}

	context.JSON(http.StatusOK, controllers.SuccessResponse{
		Message:  "Get Payload Login",
		Response: token,
	})
}

func LoginGin(context *gin.Context) {

	// get body
	body, err, status_code := GetBodyGin(context)
	if err != nil {
		controllers.ErrorHandlerGin(context, status_code, err.Error(), "LoginGin | get body")
		return
	}

	// find user
	user, err := FindUser(0, body.Username)
	if err != nil {
		controllers.ErrorHandlerGin(context, http.StatusNotFound, err.Error(), "LoginGin | validate body")
		return
	}

	// TODO: this is in Test Mode => devel for test currently
	argon := argon2.DefaultConfig()
	user_password, err := argon.HashEncoded([]byte(user.Password))
	if err != nil {
		controllers.ErrorHandlerGin(context, http.StatusInternalServerError, err.Error(), "LoginGin | validate body")
		return
	}

	// check the hash password
	ok, err := argon2.VerifyEncoded([]byte(body.Password), []byte(user_password))
	if err != nil {
		controllers.ErrorHandlerGin(context, http.StatusInternalServerError, err.Error(), "LoginGin | check Argon")
		return
	}

	if !ok {
		controllers.ErrorHandlerGin(context, http.StatusUnauthorized, "Incorrect Password", "LoginGin | compare new password with old password")
		return
	}

	// response for access_token
	access_token, err := services.AESEncrypted(&services.ObjectAES{UserID: user.UserID})
	if err != nil {
		controllers.ErrorHandlerGin(context, http.StatusInternalServerError, err.Error(), "LoginGin | encrypt access_token")
		return
	}

	context.JSON(http.StatusOK, controllers.SuccessResponse{
		Response: access_token,
	})
}

func LoginFiber(context *fiber.Ctx) error {
	// get body
	body, err, status_code := GetBodyFiber(context)
	if err != nil {
		return controllers.ErrorHandlerFiber(context, status_code, err.Error(), "LoginFiber | get body")
	}

	// find user
	user, err := FindUser(0, body.Username)
	if err != nil {
		return controllers.ErrorHandlerFiber(context, http.StatusNotFound, err.Error(), "LoginFiber | validate body")
	}

	// TODO: this is in Test Mode => devel for test currently
	argon := argon2.DefaultConfig()
	user_password, err := argon.HashEncoded([]byte(user.Password))
	if err != nil {
		return controllers.ErrorHandlerFiber(context, http.StatusInternalServerError, err.Error(), "LoginFiber | validate body")
	}

	// check the hash password
	ok, err := argon2.VerifyEncoded([]byte(body.Password), []byte(user_password))
	if err != nil {
		return controllers.ErrorHandlerFiber(context, http.StatusInternalServerError, err.Error(), "LoginFiber | check Argon")
	}

	if !ok {
		return controllers.ErrorHandlerFiber(context, http.StatusUnauthorized, "Incorrect Password", "LoginFiber | compare new password with old password")
	}

	// response for access_token
	access_token, err := services.AESEncrypted(&services.ObjectAES{UserID: user.UserID})
	if err != nil {
		return controllers.ErrorHandlerFiber(context, http.StatusInternalServerError, err.Error(), "LoginFiber | encrpt access_token")
	}

	return context.Status(fiber.StatusOK).JSON(controllers.SuccessResponse{
		Response: access_token,
	})
}
