package main

import (
	"database/sql"
	"flag"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	command string
)

func init() {
	flag.StringVar(&command, "command", "insert", "command of sql: [insert|update]")
	flag.Parse()
}

func main() {
	wg := &sync.WaitGroup{}

	var sql = "update test set val = last_insert_id(val+1) where id = 1;"
	if command == "insert" {
		sql = "insert into test (id, val) values (2, 1) on duplicate key update val = last_insert_id(val+1);"
	}

	db := initdb()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				result, err := db.Exec(sql)
				if err != nil {
					panic(err.Error())
				}
				val, err := result.LastInsertId()
				if err != nil {
					panic(err.Error())
				}
				fmt.Println(val)
			}
		}()
	}

	wg.Wait()
}

func initdb() *sqlx.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3309)/test")
	if err != nil {
		panic(err.Error())
	}
	return sqlx.NewDb(db, "mysql")
}
