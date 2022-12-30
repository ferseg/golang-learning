package config

import(
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
)

var (
    db * gorm.DB
)

func Connect() {
    d, err := gorm.Open(mysql.Open("fseg:Something@tcp(127.0.0.1:3306)/gosample?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
    if err != nil {
        panic(err)
    }
    db = d
}

func GetDb() *gorm.DB {
    return db
}
