package main

import (
	"database/sql"
	"flag"
	"fmt"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var host = flag.String("host", "localhost", "Set MySQL hostname")
var user = flag.String("user", "root", "Set MySQL username")
var password = flag.String("password", "root", "Set MySQL user password")
var port = flag.Int("port", 3306, "Set MySQL listen port")
var name = flag.String("name", "", "Set using database name in MySQL")

func init() {
	flag.Parse()
}

func main() {

	db, err := sql.Open("mysql", config().FormatDSN())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}

func config() *mysql.Config {
	c := mysql.NewConfig()
	c.DBName = *name
	c.User = *user
	c.Passwd = *password
	c.Addr = fmt.Sprintf("%s:%d", *host, *port)

	return c
}
