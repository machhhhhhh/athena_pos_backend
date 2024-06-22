package routes

import (
	testcontroller "athena-pos-backend/controllers/test"
	"athena-pos-backend/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func InitTestFiberRoutes(route fiber.Router) {

	public_router := route.Group("/test")
	private_router := route.Group("/test").Use(middlewares.AuthenticateFiber())

	// Define routes
	private_router.Get("/", testcontroller.TestGetFiber)
	private_router.Post("/", testcontroller.TestCreateFiber)

	// test login route
	public_router.Get("/get-payload/login", testcontroller.GetPayloadLoginFiber)
	public_router.Post("/login", testcontroller.LoginFiber)
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
