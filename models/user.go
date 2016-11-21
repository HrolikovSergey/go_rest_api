package models

import(
    "github.com/jinzhu/gorm"
    "time"
)

type User struct {
    gorm.Model
    Provider string
    Uid string
    EncryptedPassword string
    ResetPasswordToken string
    ResetPasswordSentAt *time.Time
    RememberCreatedAt *time.Time
    SignInCount int
    CurrentSignInAt *time.Time
    CurrentSignInIp string
    LastSignInIp string
    ConfirmationToken string
    ConfirmedAt *time.Time
    ConfirmationSentAt *time.Time
    UnconfirmedEmail string
    Name string `grom:"size:50"`
    Nickname string `grom:"size:50"`
    Image string `grom:"size:255"`
    Email string `grom:"size:255,index"`
    Tokens string
}


func (u User) Get() User{
    return User{}
}

func (u User) Save(){
    //gorm.Open("mysql", )
}

//func (u User) BeforeSave() bool{
//    return false;
//}