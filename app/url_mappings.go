package app

import "gin-rest-microservice/controllers"

func mapUrls() {
	router.POST("/users", controllers.UsersController.Create)
	router.GET("users/:id", controllers.UsersController.Get)
}
