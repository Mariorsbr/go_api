package main

import(

	"api/models"
	"api/controllers"
	//"database/sql"
	//"gorm.io/gorm"
	//"gorm.io/driver/postgres"

	//"net/http"
    "github.com/gin-gonic/gin"

)


func main(){
	router := gin.Default()

	models.ConnectDatabase() // new
	router.GET("/users",controllers.GetUser)
	router.DELETE("/users/:id",controllers.DeleteUser)
	router.POST("users", controllers.CreateUser)
	

	router.Run()
}

