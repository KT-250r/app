package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Text   string
	Status string
}

//DB init
func dbInit() {
	db, err := gorm.Open("sqlite3", "test.sqliete3")
	if err != nil {
		panic("DB cant open..!(dbInit)")
	}
	db.AutoMigrate(&Todo{})
	defer db.Close()
}

//DB add
func dbInsert(text string, status string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("DB cant open...!(dbInsert)")
	}
	db.Create(&Todo{Text: text, Status: status})
	defer db.Close()
}

//DB update
func dbUpdate(id int, text string, status string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("DB cant open...!")
	}
	var todo Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()
}

//DB delete
func dbDelete(id, int) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("DB cant open...!(dbDelete)")
	}
	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

//DB get all
func dbGetAll() []Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("DB cant open...!(dbGetAll())")
	}
	var todos []Todo
	db.order("created_at_desc").Find(&todos)
	db.Close()
	return todos
}

//DB get
func dbGetOne(id int) Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
}

func main() {
	data := "Hello Go/Gin!"
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{"data": data})
	})

	router.Run()
}
