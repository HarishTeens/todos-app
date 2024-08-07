package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)
var pool *pgxpool.Pool

func initDB() {
	dbUrl := os.Getenv("DB_URL")
	config, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		fmt.Println(err)
	}
	pool, err = pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		fmt.Println(err)
	}
	pool.AcquireAllIdle(context.Background())
}

func insertUser(pool *pgxpool.Pool, insertUser User) (User, error) {
	res, err := pool.Query(context.Background(), "INSERT INTO users (name) VALUES ($1) RETURNING id", insertUser.Name)
	if err != nil {
		return User{}, errors.Join(errors.New("error inserting user: "), err)
	}
	defer res.Close()
	for res.Next() {
		err := res.Scan(&insertUser.ID)
		if err != nil {
			return User{}, errors.Join(errors.New("error inserting user: "), err)
		}
	}
	return insertUser, nil
}

func getUserByID(pool *pgxpool.Pool, id int) (User, error) {
	var user User
	res, err := pool.Query(context.Background(), "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return User{}, errors.Join(errors.New("error getting user: "), err)
	}
	defer res.Close()
	for res.Next() {
		err := res.Scan(
			&user.ID, 
			&user.Name,
			&user.CreatedAt,
			&user.DeletedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return User{}, errors.Join(errors.New("error getting user: "), err)
		}
	}
	return user, nil
}

func getTodosByUserID(pool *pgxpool.Pool, userId int) ([]Todo, error) {
	var todos []Todo
	rows, err := pool.Query(context.Background(), "SELECT * FROM todos WHERE user_id = $1", userId)
	if err != nil {
		return []Todo{}, errors.Join(errors.New("error getting todos: "), err)
	}
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Todo, &todo.User_Id, &todo.CreatedAt, &todo.DeletedAt, &todo.UpdatedAt)
		if err != nil {
			return []Todo{}, errors.Join(errors.New("error getting todos: "), err)
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func addTodosForUser(pool *pgxpool.Pool, user_id int, todo Todo) (Todo, error) {
	res, err := pool.Query(context.Background(), "INSERT INTO todos (todo, user_id) VALUES ($1, $2) RETURNING id", todo.Todo, user_id)
	if err != nil {
		return Todo{}, errors.Join(errors.New("error inserting todo: "), err)
	}
	defer res.Close()
	for res.Next() {
		err := res.Scan(&todo.ID)
		if err != nil {
			return Todo{}, errors.Join(errors.New("error inserting todo: "), err)
		}
	}
	return todo, nil
}
