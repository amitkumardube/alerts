package dbconn

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"context"
	"time"
	conf "config"
	"strconv"
)

var (
    user string = conf.DB_user
    password string = conf.DB_password
    host string = conf.DB_hostname
    name string = conf.DB_name
)

// Defining to struct to store database returned values

type db_data struct{
    expiry_date string `json:"expiry_date,omitempty"`
    product_id int `json:"product_id,omitempty"`
    stock int `json:"stock,omitempty"`
}

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
        defer db.Close();
        tab_name := conf.Exp_tab_name
        exp_col_name := conf.Exp_col_name
        exp_id_col_name := conf.Exp_id_col_name
        exp_alert_threshold_days := conf.Exp_alert_threshold_days
        query := "SELECT "+exp_col_name+" , "+exp_id_col_name+" from "+tab_name+" where DATEDIFF("+exp_col_name+" , curdate()) <= "+strconv.Itoa(exp_alert_threshold_days)
        if (ping_database(db , name)){
            fmt.Println("Connected")
            results, err := db.Query(query)
            if err != nil {
                log.Println(err.Error())
            }
            for results.Next(){
                var data db_data
                // for each row, scan the result into our db_data composite object
                err = results.Scan(&data.expiry_date,&data.product_id)
                if err != nil {
                    log.Println(err.Error())
                }
                fmt.Printf("Expiry Date is : %s and Product ID is : %d \n",data.expiry_date,data.product_id)
            }
        }else{
            fmt.Println("Connection failed")
            return
        }
    }

    func Get_stock_data(){
        fmt.Println("Connecting to database to get product stock data")
        db := return_mysql_database_con(user, password, host, name)
        defer db.Close();
        tab_name := conf.Stock_tab_name
        stock_col_name := conf.Stock_col_name
        stock_id_col_name := conf.Stock_id_col_name
        stock_alert_threshold_quantity := conf.Stock_alert_threshold_quantity
        query := "SELECT "+stock_col_name+" , "+stock_id_col_name+" from "+tab_name+" where "+stock_col_name+" <= "+strconv.Itoa(stock_alert_threshold_quantity)
        if (ping_database(db , name)){
            fmt.Println("Connected")
            results, err := db.Query(query)
            if err != nil {
                log.Println(err.Error())
            }
            for results.Next(){
                var data db_data
                // for each row, scan the result into our db_data composite object
                err = results.Scan(&data.stock,&data.product_id)
                if err != nil {
                    log.Println(err.Error())
                }
                fmt.Printf("Stock is : %d and Product ID is : %d \n",data.stock,data.product_id)
            }
        }else{
            fmt.Println("Connection failed")
            return
        }
    }