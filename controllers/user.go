package controllers

import (
    _ "../models"
    _ "../services"
    "github.com/gin-gonic/gin"
)

func Login(c *gin.Context){
    //models.User
    c.JSON(200, gin.H{"message": "login"})
}

func Logout(c *gin.Context){
    c.JSON(200, gin.H{"message": "logout"})
}