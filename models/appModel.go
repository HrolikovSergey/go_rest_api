package models

import(
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/gin-gonic/gin"
    "../settings"
)

var (
    DB *gorm.DB
    err error
)

func init() {
    DB, err = gorm.Open("mysql",
        settings.Get().DbUser+":"+
        settings.Get().DbPassword+"@/"+
        settings.Get().DbName+
        "?charset="+settings.Get().DbCharset+
        "&parseTime=True")
    if err != nil {
        panic(err)
    }

}

func Validate(c *gin.Context, form interface{}) error{
    if err := c.Bind(form); err == nil {
        return nil
    }
    return err
}