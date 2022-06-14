package conf

type Redis struct {
	// host:port address
	Addr string `mapstructure:"addr" json:"addr" yaml:"addr" validate:"required,addr"`
	// redis 6.0+
	UserName string `mapstructure:"username" json:"username" yaml:"username" validate:"required"`
	Password string `mapstructure:"password" json:"password" yaml:"password" validate:"required"`
	DB       int    `mapstructure:"db" json:"db" yaml:"db" validate:"oneof=0 1 2 3 4 5"`
	// Maximum number of socket connections.
	PoolSize int `mapstructure:"pool_size" json:"pool_size" yaml:"pool_size" validate:"required"`
	// Minimum number of idle connections which is useful when establishing
	MinIdleConns int `mapstructure:"min_idle_conns" json:"min_idle_conns" yaml:"min_idle_conns" validate:"required"`
}
