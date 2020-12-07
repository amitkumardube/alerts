package config

import (
	"os"
	"fmt"
	"log"
	"encoding/json"
)

type Config struct {
	database_config Database_config `json:"database_config,omitempty"`
	alert_config []Alert_config `json:"alert_config,omitempty"`
	email_config Email_config `json:"email_config,omitempty"`
}

type Database_config struct {
	database_type string `json:"database_type,omitempty"`
	host string `json:"host,omitempty"`
	user string `json:"user,omitempty"`
	password string `json:"password,omitempty"`
	database string `json:"database,omitempty"`
}

type Alert_config struct {
	table_name string `json:"table_name,omitempty"`
	column_name string `json:"column_name,omitempty"`
	expiry_alert_threshold_days int `json:"expiry_alert_threshold_days,omitempty"`
	stock_alert_threshold_quantity int `json:"stock_alert_threshold_quantity,omitempty"`
}

type Email_config struct {
	smtp string `json:"smtp,omitempty"`
}

func init(){
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
	fmt.Println(conf.database_config.database_type)
}

