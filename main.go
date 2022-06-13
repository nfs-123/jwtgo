package main

import (
	"context"
	"jwtgo/controllers"
	"jwtgo/database"
	"jwtgo/middleware"
	"jwtgo/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func main() {
	r := initRouter()
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "HELLO",
	// 	})
	// })

	// r.POST("/addData", func(c *gin.Context) {
	// 	client := database.DBinstance()
	// 	var user models.User
	// 	if err := c.Bind(&user); err != nil {
	// 		c.JSON(500, gin.H{
	// 			"error": "error in binding",
	// 		})
	// 		return
	// 	}

	// 	result, err := client.Database("mydb").Collection("user").InsertOne(context.Background(), user)

	// 	if err != nil {
	// 		c.JSON(500, gin.H{
	// 			"error": "error in inserting document",
	// 		})
	// 		return
	// 	}
	// 	c.JSON(201, result)

	// })
	r.Run(":8081")

}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/addData", func(ctx *gin.Context) {
			client := database.DBinstance()
			var user models.User
			if err := ctx.Bind(&user); err != nil {
				ctx.JSON(500, gin.H{
					"error": "error in binding",
				})
				return
			}

			result, err := client.Database("mydb").Collection("user").InsertOne(context.Background(), user)

			if err != nil {
				ctx.JSON(500, gin.H{
					"error": "error in inserting document",
				})
				return
			}
			ctx.JSON(201, result)

		})
		secured := api.Group("/secured").Use(middleware.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
