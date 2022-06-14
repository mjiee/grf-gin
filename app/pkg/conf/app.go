package conf

// app config model
type App struct {
	Name string `mapstructure:"name" json:"name" yaml:"name" validate:"required"`
	Mode string `mapstructure:"mode" json:"mode" yaml:"mode" validate:"oneof=production dev debug"`
	Addr string `mapstructure:"addr" json:"addr" yaml:"addr" validate:"addr"` // host:port address
}
