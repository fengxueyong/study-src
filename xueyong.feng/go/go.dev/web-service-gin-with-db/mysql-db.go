// 这个包用于对相册进行CRUD操作的内部接口类，提供给外面的模块或package对相册进行CRUD操作
package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:                 "album",
		Passwd:               "album",
		Net:                  "tcp",
		Addr:                 "121.37.106.152:3306",
		DBName:               "album",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

func Save(ab Album) Album {
	db.Exec("INSERT INTO album (id, title, artist, price) VALUES (?, ?, ?, ?)", ab.ID, ab.Title, ab.Artist, ab.Price)
	return ab
}

func FindById(id string) (Album, error) {
	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)

	var alb Album
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return alb, nil
}

func FindAll() ([]Album, error) {
	rows, _ := db.Query("SELECT * FROM album")

	var id string
	var title string
	var artist string
	var price float64

	// slice的定义，使用[]type的语法，type可以是基本类型，或定义的结构体等
	// 比如下面的[]Album就是类型是结构的slice，再比如var names []string等等
	// 看上去像一个没有长度的数组。
	// var names []string
	var albs []Album
	for rows.Next() {
		rows.Scan(&id, &title, &artist, &price)
		albs = append(albs, Album{ID: id, Title: title, Artist: artist, Price: price})
	}

	defer rows.Close()
	return albs, nil
}
