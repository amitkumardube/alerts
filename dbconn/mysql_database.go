package dbconn

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"context"
	"time"
	conf "config"
)

var (
    user string = conf.DB_user
    password string = conf.DB_password
    host string = conf.DB_hostname
    name string = conf.DB_name
)

// forming the connection string

    func dsn(user string , pass string , host string , name string) string {
        return fmt.Sprintf("%s:%s@tcp(%s)/%s", user , pass, host, name)
    }

    func return_mysql_database_con (user string, pass string , host string , name string) *sql.DB {
	    // this function will return the reference to database connection so that query execution can be done on top of that
	    fmt.Println("Connecting to mysql database")
	    connection_str := dsn(user , pass , host , name)
	    fmt.Println("Returned Connection string : " + connection_str)
	    db, err := sql.Open("mysql", connection_str)
	    if err != nil {
	        log.Println(err.Error())
	    }
	    return db
    }

    func ping_database(db *sql.DB , dbname string) bool {
        ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancelfunc()
        err := db.PingContext(ctx)
        if err != nil {
            log.Printf("Errors %s pinging DB", err)
            return false
        }
        log.Printf("Connected to DB %s successfully\n", dbname)
        return true
    }

    func Get_expiry_data(){
        fmt.Println("Connecting to database to get product expiry data")
        db := return_mysql_database_con(user, password, host, name)
        if (ping_database(db , name)){
            fmt.Println("Connected")
        }else{
            fmt.Println("Connection failed")
            return
        }

    }

    func Get_stock_data(){
        fmt.Println("Connecting to database to get product stock data")
    }