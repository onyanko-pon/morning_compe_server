package DataBase

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func New() *gorm.DB {
    var db *gorm.DB
    db = newDB()

    sqlDB, _ := db.DB()
    sqlDB.SetMaxOpenConns(14)
    return db
}


func newDB() *gorm.DB {

    c := newConfig()

    dsn := "user=" + c.Username + " host=" + c.Host + " password=" + c.Password + " dbname=" + c.DBName + " port=" + c.Port + " sslmode=disable TimeZone=Asia/Tokyo"

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        panic(err.Error())
    }
    return db
}