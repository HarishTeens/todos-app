package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID  int64 `gorm:"primaryKey;type:autoIncrement"`
	Todo string
	User_Id int64
	CreatedAt pgtype.Date
	DeletedAt pgtype.Date
	UpdatedAt pgtype.Date
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
	godotenv.Load(".env")
	initDB()
	r.GET("/ping", pong)

	r.POST("/users", addUser)
	r.POST("/users/:id", addTodoToUser)

	r.GET("/todos/:id", getTodo)

	r.Run() 
}