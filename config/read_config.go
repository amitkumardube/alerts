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
)

type Config struct {
	Database_config Database_config `json:"database_config,omitempty"`
	Alert_config []Alert_config `json:"alert_config,omitempty"`
	Email_config Email_config `json:"email_config,omitempty"`
}

type Database_config struct {
	Database_type string `json:"type,omitempty"`
	Host string `json:"host,omitempty"`
	User string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	Database string `json:"database,omitempty"`
}

type Alert_config struct {
	Table_name string `json:"table_name,omitempty"`
	Column_name string `json:"column_name,omitempty"`
	Expiry_alert_threshold_days int `json:"expiry_alert_threshold_days,omitempty"`
	Stock_alert_threshold_quantity int `json:"stock_alert_threshold_quantity,omitempty"`
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
	DB_type = conf.Database_config.Database_type
	DB_hostname = conf.Database_config.Host
	DB_user = conf.Database_config.User
	DB_password = conf.Database_config.Password
	DB_name = conf.Database_config.Database
	fmt.Println("Database type to connect to : " , DB_type)
	fmt.Println("Database Host Name : " , DB_hostname)
	fmt.Println("Database User Name : " , DB_user)
	fmt.Println("Database User Password : " , DB_password)
	fmt.Println("Database Name : " , DB_name)
	fmt.Println("Database type to connect to : " , conf.Alert_config[0].Expiry_alert_threshold_days)
	fmt.Println("Database type to connect to : " , conf.Alert_config[1].Stock_alert_threshold_quantity)
	fmt.Println("Database type to connect to : " , conf.Email_config.Smtp)
}
