package main

import(
    "fmt"
    "./services"
    "./models"
    "./settings"
    "./controllers"
    "github.com/gin-gonic/gin"
)

func main(){
    fmt.Println(authentication.GenerateToken(models.User{}))
    //gin.SetMode(gin.ReleaseMode)
    api := gin.Default()
    authorizationRequired := api.Group("/")
    authorizationRequired.Use(authentication.IsAuthorized())
    {
        authorizationRequired.GET("/authRequired", controllers.Login)
        authorizationRequired.GET("/authRequired1", controllers.Login)
    }
    api.GET("/page1", controllers.Login)
    api.GET("/page2", controllers.Login)

    api.Run(":"+settings.Get().ServerPort)
}