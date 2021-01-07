package config

import (
	"os"
	"fmt"
	"log"
	"encoding/json"
)

var (
    DB_type string
    DB_hostname string
    DB_user string
    DB_password string
    DB_name string
    Exp_tab_name string
    Exp_col_name string
    Exp_id_col_name string
    Exp_alert_threshold_days int
    Stock_tab_name string
    Stock_col_name string
    Stock_id_col_name string
    Stock_alert_threshold_quantity int
)

type Config struct {
	Database_config Database_config `json:"database_config,omitempty"`
	Expiry_alert_config Expiry_alert_config `json:"expiry_alert_config,omitempty"`
	Stock_alert_config Stock_alert_config `json:"stock_alert_config,omitempty"`
	Email_config Email_config `json:"email_config,omitempty"`
}

type Expiry_alert_config struct {
    Table_name string `json:"table_name,omitempty"`
    Expiry_column_name string `json:"expiry_column_name,omitempty"`
    Id_column_name string `json:"id_column_name,omitempty"`
    Expiry_alert_threshold_days int `json:"expiry_alert_threshold_days,omitempty"`
}

type Stock_alert_config struct {
    Table_name string `json:"table_name,omitempty"`
    Stock_column_name string `json:"stock_column_name,omitempty"`
    Id_column_name string `json:"id_column_name,omitempty"`
    Stock_alert_threshold_quantity int `json:"stock_alert_threshold_quantity,omitempty"`
}

type Database_config struct {
	Database_type string `json:"type,omitempty"`
	Host string `json:"host,omitempty"`
	User string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	Database string `json:"database,omitempty"`
}

type Email_config struct {
	Smtp string `json:"smtp_server,omitempty"`
}

func init(){
    fmt.Println("calling the init function from config package")
	// checking the number of arguments passed
	if len(os.Args) < 2 {
		fmt.Println("Missing the file name as argument. The config file must be passed at argument.")
		os.Exit(1)
	}
	// we need to read the config file as json as map the values with the struct values
	// the config file will be passed as argument in order to get the complete file name
	file_name := os.Args[1]
	file , err := os.Open(file_name)
	if err != nil { 
		log.Fatal(err)
		return
	}
	decode_file := json.NewDecoder(file)
	decode_file.UseNumber()
	var conf Config
	err = decode_file.Decode(&conf)
	if err != nil {
		log.Fatal(err)
		return
	}
	// setting values for global variables to use in other packages.

	/* Initializing the database configuration */
	DB_type = conf.Database_config.Database_type
	DB_hostname = conf.Database_config.Host
	DB_user = conf.Database_config.User
	DB_password = conf.Database_config.Password
	DB_name = conf.Database_config.Database

	/* Initializing the product expiry configuration */
    Exp_tab_name = conf.Expiry_alert_config.Table_name
    Exp_col_name = conf.Expiry_alert_config.Expiry_column_name
    Exp_id_col_name = conf.Expiry_alert_config.Id_column_name
    Exp_alert_threshold_days = conf.Expiry_alert_config.Expiry_alert_threshold_days

    /*Initializing the product stock configuration */
    Stock_tab_name = conf.Stock_alert_config.Table_name
    Stock_col_name = conf.Stock_alert_config.Stock_column_name
    Stock_id_col_name = conf.Stock_alert_config.Id_column_name
    Stock_alert_threshold_quantity = conf.Stock_alert_config.Stock_alert_threshold_quantity

	fmt.Println("Database type to connect to : " , DB_type)
	fmt.Println("Database Host Name : " , DB_hostname)
	fmt.Println("Database User Name : " , DB_user)
	fmt.Println("Database User Password : " , DB_password)
	fmt.Println("Database Name : " , DB_name)
}
