package users

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// Returns user's langauge
func GetLang(db *pgx.Conn, tg_id int64) string {
	var lang string
	err := db.QueryRow(context.Background(), "select lang from users where tg_id = $1", fmt.Sprintf("%d", tg_id)).Scan(&lang)
	if err != nil {
		fmt.Println(err)
	}

	return lang
}

func SetLang(db *pgx.Conn, lang string, tg_id int64) {
	_, err := db.Exec(context.Background(), "update users set lang = $1 where tg_id = $2", lang,
		fmt.Sprintf("%d", tg_id))
	if err != nil {
		panic(err)
	}
}

// Checks if the user exists
func UserExists(db *pgx.Conn, tg_id int64) bool {
	var exists bool = true
	var data string
	err := db.QueryRow(context.Background(), "SELECT lang FROM users WHERE tg_id = $1", fmt.Sprintf("%d", tg_id)).Scan(&data)
	if err != nil {
		exists = false
	}
	return exists
}

// Inserts user id of a user to the table tg_id, default lang is ru
func RegMin(db *pgx.Conn, tg_id int64) {
	_, err := db.Exec(context.Background(), "insert into users (tg_id) values ($1)", fmt.Sprintf("%d", tg_id))
	if err != nil {
		panic(err)
	}
}
