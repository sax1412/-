package db

import (
	"database/sql"
	"time"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root@/pachong?charset=utf8")
}

func Insert(name, title, url string) {
	exe, _ := db.Prepare("insert star set name = ?,url = ?, ct = ?,title = ?")
	exe.Exec(name, url, time.Now().UTC(), title)
}
