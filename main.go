package main

import (
	"database/sql"
	"flag"

	_ "github.com/go-sql-driver/mysql"
)

var host = flag.String("host", "localhost", "Set MySQL hostname")
var user = flag.String("user", "root", "Set MySQL username")
var password = flag.String("password", "root", "Set MySQL user password")
var port = flag.Int("port", 3306, "Set MySQL listen port")

func init() {
	flag.Parse()
}

func main() {
	db, err := sql.Open("mysql", "root:@/my_database")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}
