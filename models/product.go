package models

import(
    "github.com/jinzhu/gorm"
)

type Product struct{
    gorm.Model
}

func init(){
    if !DB.HasTable(&Product{}) {
        DB.AutoMigrate(&Product{})
    }
}