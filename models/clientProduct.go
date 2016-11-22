package models

import(
    "github.com/jinzhu/gorm"
)

type ClientProduct struct{
    gorm.Model
    User User
    UserID int
    Product Product
    ProductID int
}

func init(){
    if !DB.HasTable(&ClientProduct{}) {
        DB.AutoMigrate(&ClientProduct{})
    }
}

func (cp ClientProduct) TableName() string {
    return "client_products"
}