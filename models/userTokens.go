package models

import(
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "time"
    //"../settings"
)

type UserToken struct {
    gorm.Model
    UserId uint
    Token  string
    ValidTill *time.Time
}

func UserTokenSave(userId uint, token string){
    tokenRecord := UserToken{
        UserId: userId,
        Token: token,
        //ValidTill: time.Now().Add(time.Duration(time.Second * settings.Get().AuthTokenExpire)),
    }
    DB.Create(tokenRecord)
}