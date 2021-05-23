package app

import (
	"github.com/TriTranDev/bookstore_users_api/Controllers/ping"
	"github.com/TriTranDev/bookstore_users_api/Controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
