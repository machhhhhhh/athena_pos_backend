package services

import (
	"athena-pos-backend/models"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func GetContextUserGin(context *gin.Context) (models.User, error) {
	user := context.MustGet("user")

	var userMe models.User

	data, err := json.Marshal(user)
	if err != nil {
		return userMe, err
	}
	err = json.Unmarshal(data, &userMe)
	if err != nil {
		return userMe, err
	}

	return userMe, nil
}
func GetContextUserFiber(context *fiber.Ctx) (models.User, error) {
	user := context.Locals("user")

	var userMe models.User

	fmt.Println("user", user)

	data, err := json.Marshal(user)
	if err != nil {
		return userMe, err
	}
	err = json.Unmarshal(data, &userMe)
	if err != nil {
		return userMe, err
	}

	return userMe, nil
}
