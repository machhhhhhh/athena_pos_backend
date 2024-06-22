package routes

import (
	testcontroller "athena-pos-backend/controllers/test"
	"athena-pos-backend/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

// TODO: from doc fiber cannot declare 2 router group from same parent path
// TODO: Fiber uses method chaining directly on the router instance
// TODO: Gin uses methods directly on the router group
func InitTestFiberRoutes(route fiber.Router) {

	router := route.Group("/test")

	// Define routes
	router.Get("/", middlewares.AuthenticateFiber(), testcontroller.TestGetFiber)
	router.Post("/", middlewares.AuthenticateFiber(), testcontroller.TestCreateFiber)

	// test login route
	router.Get("/get-payload/login", testcontroller.GetPayloadLoginFiber)
	router.Post("/login", testcontroller.LoginFiber)
}

func InitTestGinRoutes(route *gin.RouterGroup) {
	private_router := route.Group("/test").Use(middlewares.AuthenticateGin())
	public_router := route.Group("/test")

	// Define routes
	private_router.GET("/", testcontroller.TestGetGin)
	private_router.POST("/", testcontroller.TestCreateGin)

	// test login route
	public_router.GET("/get-payload/login", testcontroller.GetPayloadLoginGin)
	public_router.POST("/login", testcontroller.LoginGin)
}
