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
	userID, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	// Create channels for communication between goroutines
	userCh := make(chan User)
	errCh := make(chan error)

	go func ()  {
		user, _ := getUserByID(pool, int(userID))
		if user.ID == 0 {
			errCh <- fmt.Errorf("user with id %v not found", userID)
			return
		}
		userCh <- user
	}()

	// Wait for results from goroutines
	var user User
	select {
		case user = <-userCh:
		case err := <-errCh:
			c.JSON(http.StatusBadRequest, gin.H{
				"error_message": err,
			})
			return
		}

	// Create channels for communication between goroutines
	todosCh := make(chan []Todo)

	// Fetch todos concurrently
	go func() {
		todos, err := getTodosByUserID(pool, int(user.ID))
		if err != nil {
			errCh <- err
			return
		}
		todosCh <- todos
	}()

	// Wait for results from goroutines
	var todos []Todo
	select {
	case todos = <-todosCh:
	case err := <-errCh:
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": fmt.Sprintf("%v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"todos":   todos,
		"user_id": userID,
	})
}

func addTodoToUser(c *gin.Context) {
	user_id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user, _ := getUserByID(pool, int(user_id))
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