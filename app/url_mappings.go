package app

import (
	"github.com/tajud99n/bookstore_users-api/controllers/users"
)

func mapURLS() {
	router.POST("/users", users.CreateUser)
	router.GET("/users/:userId", users.GetUser)
	router.PUT("/users/:userId", users.UpdateUser)
	router.PATCH("/users/:userId", users.UpdateUser)
	router.DELETE("/users/:userId", users.DeleteUser)
	router.GET("/internal/users/search", users.Search)
	router.POST("/users/login", users.Login)
}
