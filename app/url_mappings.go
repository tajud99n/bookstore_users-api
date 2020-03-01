package app

import (
	"github.com/tajud99n/bookstore_users-api/controllers/users"
)

func mapURLS() {
	router.POST("/users", users.CreateUser)
	router.GET("/users/:userId", users.GetUser)
}
