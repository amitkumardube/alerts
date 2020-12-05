package config

type config struct {
	database_config database_config `json:"database_config,omitempty"`
}

type database_config struct {
	:w

}

type alert_config struct {
	table
}

type email_config struct {


} 



