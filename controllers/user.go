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
    form := &forms.Login{}
    err := models.Validate(c, form)
    if err == nil {
        user := models.UserFindByCredentials(form.Email, authentication.GetPassword(form.Password))
        if user.ID != 0{
            response["authToken"] = authentication.GenerateToken(user)
            models.UserTokenSave(user.ID, response["authToken"].(string))
        } else {
            response["error"] = "User not found"
        }
        c.JSON(200, response)
    } else {
        response["error"] = "Validation fails"
        c.JSON(400, response)
    }
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
    models.DeleteToken(c.Request.Header.Get("AuthToken"))
    c.JSON(200, gin.H{"message": "ok"})
}