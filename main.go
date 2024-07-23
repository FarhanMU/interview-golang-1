package main

import (
	"fmt"
	"net/http"
	"zen-test/app"
	"zen-test/app/helpers"
	"zen-test/app/middleware"

	docs "zen-test/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @contact.name hi.boedi8@gmail.com
// @contact.url https://www.github.com/hiboedi
// @contact.email hi.boedi8@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Bearer token authentication
func main() {
	router, appConfig := app.Init()

	docs.SwaggerInfo.Title = "Zenstore API"
	docs.SwaggerInfo.Description = "Sample of Zenstore API Documentation."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8000"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	authRouter := middleware.AuthMiddleware(router)

	fmt.Println("Welcome to " + appConfig.AppName)
	fmt.Println("Starting server on " + appConfig.AppURL)

	server := http.Server{
		Addr:    "localhost:" + appConfig.AppPort,
		Handler: authRouter,
	}
	err := server.ListenAndServe()
	helpers.PanicIfError(err)
}
