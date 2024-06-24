package testcontroller

import (
	"athena-pos-backend/models"
	"athena-pos-backend/utils"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

// func GetBodyGin(context *gin.Context,body LoginRequest) (LoginRequest, error, int) {
// 	context_body := context.MustGet("body")
// 	data, err := json.Marshal(context_body)
// 	if err != nil {
// 		return LoginRequest{}, errors.New("Cannot Marshal Body"), http.StatusInternalServerError
// 	}
// 	err = json.Unmarshal(data, &body)
// 	if err != nil {
// 		return LoginRequest{}, errors.New("Please input the correct_type body"), http.StatusBadRequest
// 	}

// 	// validate body
// 	body, err = validateBody(body)
// 	if err != nil {
// 		return LoginRequest{}, err, http.StatusBadRequest
// 	}

// 	// body.
// 	return body, nil, http.StatusOK
// }

func GetBodyGin[T any](context *gin.Context) (T, error, int) {
	var body T
	context_body := context.MustGet("body")
	data, err := json.Marshal(context_body)
	if err != nil {
		return body, errors.New("Cannot Marshal Body"), http.StatusInternalServerError
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		return body, errors.New("Please input the correct_type body"), http.StatusBadRequest
	}

	return body, nil, http.StatusOK
}
func GetBodyFiber[T any](context *fiber.Ctx) (T, error, int) {
	var body T
	context_body := context.Locals("body")
	data, err := json.Marshal(context_body)
	if err != nil {
		return body, errors.New("Cannot Marshal Body"), http.StatusInternalServerError
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		return body, errors.New("Please input the correct_type body"), http.StatusBadRequest
	}

	return body, nil, http.StatusOK
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

func FindUser(user_id int, username string) (models.User, error) {

	var user models.User

	if len(ALL_USER) != 0 {
		for i := range ALL_USER {
			if ALL_USER[i].Username == username || ALL_USER[i].UserID == user_id {
				return ALL_USER[i], nil
			}
		}
	}
	return user, errors.New(username + " not founded")
}

func GetUser() []models.User {
	return ALL_USER
}

func CreateUser(body models.User) models.User {
	body.UserID = len(ALL_USER) + 1
	ALL_USER = append(ALL_USER, body)
	return body
}
func UpdateUser(user_id int, body models.User) {
	if len(ALL_USER) != 0 && user_id != 0 {
		for i := range ALL_USER {
			if ALL_USER[i].UserID == user_id {
				ALL_USER[i] = body
			}
		}
	}
}
func DeleteUser(user_id int) {
	if len(ALL_USER) != 0 && user_id != 0 {
		var all_new_user []models.User
		for i := range ALL_USER {
			if ALL_USER[i].UserID != user_id {
				all_new_user = append(all_new_user, ALL_USER[i])
			}
		}
		ALL_USER = all_new_user
	}
}
