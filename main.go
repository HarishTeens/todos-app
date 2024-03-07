package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID  int64 `gorm:"primaryKey;type:autoIncrement"`
	Todo string
	User_Id int64
}
func (u *Todo) TableName() string {
    // custom table name, this is default
    return "public.todos"
}
  
type User struct {
	gorm.Model
	ID  int64 `gorm:"primaryKey;type:autoIncrement"`
	Name string
}
func (u *User) TableName() string {
    // custom table name, this is default
    return "public.users"
}

func main() {
	r := gin.Default()
	dsn := "host=localhost user=postgres password=postgres dbname=mydb port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/users", func(c *gin.Context) {
		insertUser := User{Name: "Surya"}

		db.Create(&insertUser)

		c.JSON(http.StatusOK, gin.H{
			"user": insertUser,
		})
	})


	r.GET("/todos/:id", func(c *gin.Context) {
		user_id := c.Param("id")
		todos := []Todo{}
		err := db.Find(&todos, "user_id = ?", user_id).Error

		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"todos": todos,
			"user_id": user_id,
		})
	})

	r.POST("/users/:id", func(c *gin.Context) {
		user_id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

		userRecord := User{}
		err := db.First(&userRecord,  "id = ?", user_id).Error

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error_message": "User not found",
			})
			return
		}

		insertTodo := Todo{Todo: "blah todo", User_Id: user_id}

		err = db.Create(&insertTodo).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error_message": err,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"todo": insertTodo,
		})
	})

	r.Run() 
}