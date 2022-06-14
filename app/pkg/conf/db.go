package conf

// Relational database
type Db struct {
	Driver string `mapstructure:"driver" json:"driver" yaml:"driver"`
	// host:port address
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr" validate:"addr"`
	Database string `mapstructure:"database" json:"database" yaml:"database" validate:"required"`
	UserName string `mapstructure:"username" json:"username" yaml:"username" validate:"required"`
	Password string `mapstructure:"password" json:"password" yaml:"password" validate:"required"`
	Charset  string `mapstructure:"charset" json:"charset" yaml:"charset" validate:"required"`
	// The maximum amount of time a connection may be reused, s.
	MaxLifeTime int `mapstructure:"max_life_time" json:"max_life_time" yaml:"max_life_time" validate:"required"`
	// The maximum number of mapstructure:"" idle connection
	MaxIdleConns int `mapstructure:"max_idle_conns" json:"max_idle_conns" yaml:"max_idle_conns" validate:"required"`
	// The maximum number of mapstructure:"" connection
	MaxOpenConns int    `mapstructure:"max_open_conns" json:"max_open_conns" yaml:"max_open_conns" validate:"required"`
	LogFile      string `mapstructure:"log_file" json:"log_file" yaml:"log_file" validate:"required"`
	// log level: 1 Slient, 2 Error, 3 Warn, 4 Info
	LogLevel int8 `mapstructure:"log_level" json:"log_level" yaml:"log_level" validate:"oneof=1 2 3 4"`
	// slow log tiem, s.
	SlowLog int `mapstructure:"slow_log" json:"slog_log" yaml:"slow_log" validate:"required"`
}
