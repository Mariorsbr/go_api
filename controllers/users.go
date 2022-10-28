package controllers 

import (
	"net/http"

  	"github.com/gin-gonic/gin"
  	"api/models"
)

func GetUser(c *gin.Context){
	var users []models.User
  	models.DB.Find(&users)

  	c.JSON(http.StatusOK, gin.H{"data": users})
}

func DeleteUser(c *gin.Context ){
	id := c.Param("id")
	var users []models.User
	models.DB.Find(&users)
	models.DB.Delete(&models.User{}, id)
	
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func CreateUser(c *gin.Context){
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	models.DB.Create(&user)
	c.IndentedJSON(http.StatusCreated, user)

}
