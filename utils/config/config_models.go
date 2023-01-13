package config

type ConfigNodes struct {
	ConfigPath Paths `mapstructure:"config_path"`
	ConfigDB   DB    `mapstructure:"config_db"`
}

type Paths struct {
	Product_id string `mapstructure:"product_id"`
	Quantity   int    `mapstructure:"quantity"`
	Workera    string `mapstructure:"workera"`
	Workerb    string `mapstructure:"workerb"`
	Workerc    string `mapstructure:"workerc"`
	Workerd    string `mapstructure:"workerd"`
}

type DB struct {
	PostgreUser     string `mapstructure:"postgre_user"`
	PostgrePassword string `mapstructure:"postgre_password"`
	PostgreHost     string `mapstructure:"postgre_host"`
	PostgrePort     string `mapstructure:"postgre_port"`
	PostgreSchema   string `mapstructure:"postgre_schema"`
	PostgreSSLMode  string `mapstructure:"postgre_sslmode"`
}
