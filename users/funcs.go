package users

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func GetLang(db *pgx.Conn, tg_id int64) string {
	var lang string
	err := db.QueryRow(context.Background(), "select lang from users where tg_id = $1", fmt.Sprintf("%d", tg_id)).Scan(&lang)
	if err != nil {
		panic(err)
	}
	return lang
}

func UserExists(db *pgx.Conn, tg_id int64) bool {
	var exists bool = true
	var data string
	err := db.QueryRow(context.Background(), "SELECT EXISTS (SELECT 1 FROM your_table WHERE tg_id = $1)", fmt.Sprintf("%d", tg_id)).Scan(&data)
	if err != nil {
		exists = false
	}
	// fmt.Println(data)
	return exists
}

// Inserts user id of a user to the table tg_id
func RegMin(db *pgx.Conn, tg_id int64) {
	_, err := db.Exec(context.Background(), "insert into users (tg_id) values ($1)", fmt.Sprintf("%d", tg_id))
	if err != nil {
		panic(err)
	}
}
