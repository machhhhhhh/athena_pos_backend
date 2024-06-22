package testcontroller

import (
	"athena-pos-backend/models"
	"athena-pos-backend/utils"
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func GetBodyGin(context *gin.Context) (LoginRequest, error) {
	var body LoginRequest
	context_body := context.MustGet("body")
	data, err := json.Marshal(context_body)
	if err != nil {
		return LoginRequest{}, errors.New("Cannot Marshal body")
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		return LoginRequest{}, errors.New("Please input the correct_type body")
	}
	return body, nil
}
func GetBodyFiber(context *fiber.Ctx) (LoginRequest, error) {
	var body LoginRequest
	context_body := context.Locals("body")
	data, err := json.Marshal(context_body)
	if err != nil {
		return LoginRequest{}, errors.New("Cannot Marshal body")
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		return LoginRequest{}, errors.New("Please input the correct_type body")
	}
	return body, nil
}

func validateBody(body LoginRequest) (LoginRequest, error) {
	body.Username = utils.TrimString(body.Username)
	body.Password = utils.TrimString(body.Password)
	if body.Username == "" {
		return body, errors.New("username is empty")
	}
	if body.Password == "" {
		return body, errors.New("password is empty")
	}

	return body, nil
}

func FindUserByID(user_id interface{}) (models.User, error) {

	var user models.User

	if len(ALL_USER) != 0 {
		for i := range ALL_USER {
			if ALL_USER[i].UserID == user_id.(int) {
				return ALL_USER[i], nil
			}
		}
	}
	return user, errors.New(user_id.(string) + " not founded")
}

func FindUserByUsername(username string) (models.User, error) {

	var user models.User

	if len(ALL_USER) != 0 {
		for i := range ALL_USER {
			if ALL_USER[i].Username == username {
				return ALL_USER[i], nil
			}
		}
	}
	return user, errors.New(username + " not founded")
}
