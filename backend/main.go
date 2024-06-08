package main

import (
	"fmt"

	"example.com/m/controllers"
	"example.com/m/initializers"
	"github.com/gin-contrib/cors"

	// "github.com/machinebox/graphql"
	"github.com/gin-gonic/gin"
	// "github.com/hasura/go-graphql-client"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3001"}
	r.Use(cors.New(config))

	r.Static("/assets", "./assets")
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.POST("/upload", controllers.Upload)
	fmt.Println(fmt.Sprint(232323))

	r.Run()
}
