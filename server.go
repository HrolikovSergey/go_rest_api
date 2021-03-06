package main

import(
    "./settings"
    "./services"
    "./controllers"
    "github.com/gin-gonic/gin"
)

func main(){
//    fmt.Println(authentication.GenerateToken(models.User{}))
    api := gin.Default()
    authorizationRequired := api.Group("/")
    authorizationRequired.Use(authentication.IsAuthorized())
    {
        authorizationRequired.GET("/authRequired1", controllers.DoSomeThing)
        authorizationRequired.GET("/logout", controllers.Logout)
    }
    api.POST("/login", controllers.Login)
    api.POST("/signup", controllers.Signup)

    api.Run(":"+settings.Get().ServerPort)
}