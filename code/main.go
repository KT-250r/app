package main

import (
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    "github.com/mattn/go-sqlite3"
)

type Todo struct {
    gorm.Model
    Text string
    Status string
}

finc dbInit(){
    db, err := gorm.Open("sqlite3", "test.sqliete3")
    if err != nil {
        panic("DB can't open..!")
    }
    db.AutoMigrate(&Todo{})
    defer db.Close()
}

func main() {
    data := "Hello Go/Gin!"
    router := gin.Default()
    router.LoadHTMLGlob("templates/*.html")

    router.GET("/", func(ctx *gin.Context){
	    ctx.HTML(200, "index.html", gin.H{"data": data})
    })

    router.Run()
}

