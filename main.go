package main

import (
	"athena-pos-backend/middlewares"
	"athena-pos-backend/routes"
	socket_service "athena-pos-backend/sockets"
	service_usb "athena-pos-backend/usb"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"athena-pos-backend/docs"

	gin_cors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	fiber_cors "github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gofiber/contrib/swagger"
)

func main() {

	// Set timezone
	os.Setenv("TZ", "Asia/Bangkok")

	// TODO: Goroutines here (on middleware => representative)
	// TODO: have to command => go middlewares.FocusUserRepresentExpire()

	// Setup routes
	SetupGinRoutes()
	// SetupFiberRoutes()

}

func SetupGinRoutes() {
	// Load .env
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file : " + err.Error())
	}

	cors_origin := os.Getenv("ATHENA_CORS_ORIGIN")
	gin_mode := os.Getenv("ATHENA_GIN_MODE")
	server_port := os.Getenv("ATHENA_SERVER_PORT")
	socket_port := os.Getenv("ATHENA_SOCKET_PORT")
	jwt_secret := os.Getenv("ATHENA_JWT_SECRET")
	aes_iv := os.Getenv("ATHENA_AES_IV")
	aes_key := os.Getenv("ATHENA_AES_KEY")

	if cors_origin == "" || server_port == "" || socket_port == "" || jwt_secret == "" || gin_mode == "" || aes_key == "" || aes_iv == "" {
		panic("environment not founded")
	}

	gin.SetMode(gin_mode)
	app := gin.Default()

	// rate_limiter middleware
	app.Use(middlewares.RateLimiterGin(100, time.Minute))

	// cors
	app.Use(gin_cors.New(gin_cors.Config{
		AllowOrigins:     []string{cors_origin},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowAllOrigins:  false,
	}))

	// services payload
	app.Use(middlewares.ValidateHTTPMethodGin())

	// api
	SetUpAPIGin(app)

	// USB Machine
	service_usb.SetupUSB()

	// Start Gin Server
	go func() {
		fmt.Println("Server Listening on port:", server_port)
		app.Run(":" + server_port)
	}()

	// Socket move socket to last section
	socket_service.SetupSocket(socket_port)

}

func SetupFiberRoutes() {

	// Load .env
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file : " + err.Error())
	}

	cors_origin := os.Getenv("ATHENA_CORS_ORIGIN")
	server_port := os.Getenv("ATHENA_SERVER_PORT")
	socket_port := os.Getenv("ATHENA_SOCKET_PORT")
	jwt_secret := os.Getenv("ATHENA_JWT_SECRET")
	aes_iv := os.Getenv("ATHENA_AES_IV")
	aes_key := os.Getenv("ATHENA_AES_KEY")

	if cors_origin == "" || server_port == "" || socket_port == "" || jwt_secret == "" || aes_key == "" || aes_iv == "" {
		panic("environment not founded")
	}

	// Create a new Fiber instance
	app := fiber.New(fiber.Config{
		// TODO: (Prefork) need to understand this doc this is affect to socket (Production must set to true)
		Prefork: false, // if set true make sure dockerfile setting to [CMD ./app]
		// ProxyHeader: true, // make loading balance [c.IP()]
		CaseSensitive: true,
		StrictRouting: false, // true => allow only cors (need to read doc more) // now development use false
		ServerHeader:  "Fiber",
		AppName:       "Athena v1.0.0",
		Immutable:     true, // ***
	})

	// cors
	app.Use(fiber_cors.New(fiber_cors.Config{
		AllowOrigins:     cors_origin,
		AllowMethods:     "GET,POST,PUT,DELETE,PATCH",
		AllowHeaders:     "*",
		ExposeHeaders:    "*",
		AllowCredentials: true,
	}))

	// rate_limiter middleware
	app.Use(middlewares.RateLimiterFiber(100, time.Minute)) // 1 minute per 100 requests for 1 IP

	// file upload routes
	app.Static("/api/assets", "./uploads")

	// services payload
	app.Use(middlewares.ValidateHTTPMethodFiber())

	// api
	SetUpAPIFiber(app)

	// USB Machine
	service_usb.SetupUSB()

	// Start Fiber Server
	go func() {
		fmt.Println("Server Listening on port:", server_port)
		log.Fatal(app.Listen(":" + server_port))
	}()

	// Socket move socket to last section
	socket_service.SetupSocket(socket_port)
}

func SetUpAPIFiber(app *fiber.App) {
	// api path
	api := app.Group("/api") // {{domain_url}}/api

	// all router
	routes.InitTestFiberRoutes(api)

	// swagger base path
	app.Use(swagger.New(swagger.Config{
		BasePath: "/api",
		FilePath: "./docs/swagger.json",
		Path:     "docs",
	}))

	app.Get("/swagger/*", fiberSwagger.WrapHandler)
}
func SetUpAPIGin(app *gin.Engine) {
	// api path
	api := app.Group("/api") // {{domain_url}}/api

	// all router
	routes.InitTestGinRoutes(api)

	// swagger base path
	docs.SwaggerInfo.BasePath = "/api"

	// swagger routes
	app.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		ginSwagger.DefaultModelsExpandDepth(-1),
		ginSwagger.PersistAuthorization(true)),
	)

	http.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.Dir("./docs"))))
}
