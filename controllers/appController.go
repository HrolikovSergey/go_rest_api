package controllers

import(
    "github.com/gin-gonic/gin"
)

//func validate(c *gin.Context, formName interface{}) interface{}{
//    c.Bind(formName)
//    return formName
//}

func DoSomeThing(c *gin.Context){
    c.JSON(200, gin.H{"message": "do some auth required action"})
}