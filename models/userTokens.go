package models

import(
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "time"
    "../settings"
    "fmt"
)

type UserToken struct {
    gorm.Model
    Token  string `gorm:"type:varchar(1024)"`
    UserId uint
    ValidTill time.Time
}

func (UserToken) TableName() string {
    return "user_tokens"
}

func init(){
    if !DB.HasTable(&UserToken{}) {
        DB.AutoMigrate(&UserToken{})
    }
}

func UserTokenSave(userId uint, token string){
    tokenRecord := &UserToken{
        UserId: userId,
        Token: token,
        ValidTill: time.Now().Add(time.Second * time.Duration(settings.Get().AuthTokenExpire)),
    }
    DB.Create(tokenRecord)
}

func CheckToken(token string) bool{
    userToken := UserToken{}
    DB.LogMode(true)
    DB.Where("token=?", token).First(&userToken)
    DB.LogMode(false)
    fmt.Println(userToken);
    if userToken.ID == 0 {
        return false
    }
    //if userToken.ValidTill < time.Now() {
    //    DeleteToken(token)
    //    return false
    //}
    return true
}

func DeleteToken(token string){
    DB.Where("token=?", token).Delete(&UserToken{})
}