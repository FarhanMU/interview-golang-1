package router

import (
	"zen-test/app/middleware"
	"zen-test/app/web/controllers"

	"github.com/gorilla/mux"
)

func InitializeRouter(
	userController controllers.UserController,
	productController controllers.ProductController,
	orderController controllers.OrderController,
) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/user/login", userController.Login).Methods("POST")
	router.HandleFunc("/user/signup", userController.SignUp).Methods("POST")
	router.HandleFunc("/user/signup/{userId}", userController.Update).Methods("PUT")
	router.HandleFunc("/user/logout", userController.Logout).Methods("GET")

	router.HandleFunc("/products", productController.Create).Methods("POST")
	router.HandleFunc("/products", productController.FindAll).Methods("GET")
	router.HandleFunc("/products/{productId}", productController.Update).Methods("PUT")
	router.HandleFunc("/products/{productId}", productController.FindById).Methods("GET")
	router.HandleFunc("/products/{productId}", productController.Delete).Methods("DELETE")

	router.HandleFunc("/orders", orderController.FindAllOrder).Methods("GET")
	router.HandleFunc("/orders", orderController.CreateOrder).Methods("POST")

	router.Use(middleware.RecoverMiddleware)

	return router
}
