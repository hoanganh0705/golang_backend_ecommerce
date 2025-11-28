package settings

type Config struct {
	Mysql  MySQLSetting `mapstructure:"mysql"` // mapstructure ở đây nghĩa là trỏ tới file local.yaml
	Logger LogSetting   `mapstructure:"logger"`
}

type MySQLSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	UserName        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Dbname          string `mapstructure:"dbname"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}

type LogSetting struct {
	Log_level     string `mapstructure:"log_level"`
	Log_file_name string `mapstructure:"log_file_name"`
	Max_size      int    `mapstructure:"max_size"`
	Max_backups   int    `mapstructure:"max_backups"`
	Max_age       int    `mapstructure:"max_age"`
	Compress      bool   `mapstructure:"compress"`
}
