package models

import(
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "time"
)

type User struct {
    gorm.Model
    Provider string `gorm:"default:'email'"`
    Uid string
    Password string `gorm:"-"`
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
    Name string `gorm:"size:50"`
    Nickname string `gorm:"size:50"`
    Image string `gorm:"size:255"`
    Email string `gorm:"size:255,index"`
    Tokens string
}

func UserFindByCredentials(email string, password string) (user User){
    DB.Where("encrypted_password = ? AND email = ?", password, email).First(&user)
    return
}

func (user User) Create(){
    DB.Create(&user)
}