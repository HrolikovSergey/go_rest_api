package controllers

import (
    "../models"
    "../services"
    "../models/forms"
    "github.com/gin-gonic/gin"
    "fmt"
)

func Login(c *gin.Context){
    response := gin.H{}
    formData := &forms.Login{}
    err := models.Validate(c, formData)
    if(err != nil){
        c.JSON(400, response)
        return
    }
    user := models.UserFindByCredentials(formData.Email, authentication.GetPassword(formData.Password))
    fmt.Println(user)
    if user != (models.User{}){
        response["authToken"] = authentication.GenerateToken(user)
        //models.UserTokenSave(user.ID, response["authToken"].(string))
    } else {
        response["error"] = "User not found"
    }
    c.JSON(200, response)
}

func Signup(c *gin.Context){
    user := models.User{
        Nickname: "test",
        Email: "asd@asd.com",
        EncryptedPassword: authentication.GetPassword("asd"),
    }

    fmt.Print(user.ID)

    user.Create()
}

func Logout(c *gin.Context){
    c.JSON(200, gin.H{"message": "logout"})
}