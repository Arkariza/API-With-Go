package main

import (
	"github.com/user/api_item/auth"
	"github.com/user/api_item/controller/itemcontroller"
	"github.com/user/api_item/controller/usercontroller"
	"github.com/user/api_item/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectionDatabase()

	r.POST("/api/login", usercontroller.Login)   

	authorized := r.Group("/")
	authorized.Use(auth.JWTAuthMiddleware()) 

	{
		// Item
		authorized.GET("api/items", itemcontroller.Index)
		authorized.GET("api/items/:id", itemcontroller.Show)
		authorized.POST("api/items", itemcontroller.Create)
		authorized.PUT("api/items/:id", itemcontroller.Update)  
		authorized.DELETE("api/items/:id", itemcontroller.Delete) 

		// User
		authorized.GET("/api/users/:id", usercontroller.Show)
		authorized.DELETE("/api/users/:id", usercontroller.Delete)
		authorized.POST("/api/users", usercontroller.Create)
		authorized.GET("/api/users", usercontroller.Index)
		authorized.PUT("/api/users/:id", usercontroller.Update)
	}

	r.Run()
}
