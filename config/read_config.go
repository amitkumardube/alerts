package config

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
	smtp string `json:"smpt,omitempty"`
}



