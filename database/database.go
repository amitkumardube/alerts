package database

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

func mysql_database_con (){
	// this function will return the refrence to database connection so that query execution can be done on top of that
	fmt.Println("Connecting to mysql database")

}
