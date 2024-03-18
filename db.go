package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/exaring/otelpgx"
	"github.com/jackc/pgx/v5/pgxpool"
)
var pool *pgxpool.Pool

func initDB() {
	dbUrl := os.Getenv("DB_URL")
	config, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		fmt.Println(err)
	}
	config.ConnConfig.Tracer = otelpgx.NewTracer()
	pool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func insertUser(pool *pgxpool.Pool, insertUser User, ctx context.Context) (User, error) {
	res, err := pool.Query(ctx, "INSERT INTO users (name) VALUES ($1) RETURNING id", insertUser.Name)
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

func getUserByID(pool *pgxpool.Pool, id int, ctx context.Context) (User, error) {
	var user User
	res, err := pool.Query(ctx, "SELECT * FROM users WHERE id = $1", id)
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

func getTodosByUserID(pool *pgxpool.Pool, userId int, ctx context.Context) ([]Todo, error) {
	var todos []Todo
	rows, err := pool.Query(ctx, "SELECT * FROM todos WHERE user_id = $1", userId)
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

func addTodosForUser(pool *pgxpool.Pool, user_id int, todo Todo, ctx context.Context) (Todo, error) {
	res, err := pool.Query(ctx, "INSERT INTO todos (todo, user_id) VALUES ($1, $2) RETURNING id", todo.Todo, user_id)
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
