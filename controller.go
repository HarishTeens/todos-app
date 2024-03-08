package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Pong",
	})
}

func addUser(c *gin.Context) {
	insertUser,err := insertUser(pool, User{Name: "blah"})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": insertUser,
	})
}

func getTodo(c *gin.Context) {
	user_id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	user, _ := getUserById(pool, int(user_id))
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": fmt.Sprintf("user with id %v not found", user_id),
		})
		return
	}
	
	rowss, err := getTodosByUserId(pool, int(user_id))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": fmt.Sprintf("%v", err),
		})
		return			
	}

	c.JSON(http.StatusOK, gin.H{
		"todos": rowss,
		"user_id": user_id,
	})

}

func addTodoToUser(c *gin.Context) {
	user_id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user, _ := getUserById(pool, int(user_id))
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": fmt.Sprintf("user with id %v not found", user_id),
		})
		return
	}

	insertTodo, err := addTodosForUser(pool, int(user_id), Todo{Todo: "blah todo"})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"todo": insertTodo,
	})
}