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
    Tokens []UserToken `gorm:"ForeignKey:user_id;"`
}

func init(){
    if !DB.HasTable(&User{}) {
        DB.AutoMigrate(&User{})
    }
}

func UserFindByCredentials(email string, password string) (user User){
    DB.Where("users.encrypted_password = ? AND users.email = ?", password, email).First(&user)
    return
}

func (user User) Create(){
    DB.Create(&user)
}

func UserByToken(token string, userId uint) (user User){
    user = User{}
    user.ID = userId
    //DB.LogMode(true)
    DB.Model(&user).Related(&UserToken{}).First(&user)
    //DB.LogMode(false)
    //DB.Model(&user).Related(UserToken{}).Where("token = ? AND valid_till < ?", token, time.Now()).First(&user)
    //DB.Where("users.id = ? AND tokens.token = ? AND valid_till < ?", userId, token, time.Now()).Related(UserToken{}).First(&user)
    return
}