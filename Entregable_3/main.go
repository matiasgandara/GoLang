package main

//_ "github.com/go-sql-driver/mysql"

func main() {
	var db xsqlx.DB
	db = sqlx.Open("mysql", ":memory")

}
