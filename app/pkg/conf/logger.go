package conf

// Logger config model
type Log struct {
	Compress bool `mapstructure:"compress" json:"compress" yaml:"compress"`
	// Debug Level: -1, Info: 0, Warn: 1, Error: 2, DPanic: 3, Panic: 4, Fatal: 5
	Level int8 `mapstructure:"level" json:"level" yaml:"level" validate:"oneof=-1 0 1 2 3 4 5"`
	// MB
	MaxSize int `mapstructure:"max_size" json:"max_size" yaml:"max_size" validate:"required"`
	// day
	MaxAge int `mapstructure:"max_age" json:"max_age" yaml:"max_age" validate:"required"`
	// number of logs
	MaxBackups int `mapstructure:"max_backups" json:"max_backups" yaml:"max_backups" validate:"required"`
	// storage directory
	Filename string `mapstructure:"filename" json:"filename" yaml:"filename" validate:"required"`
	// skip logging for specific path
	SkipPaths []string `mapstructure:"skip_paths" json:"skip_paths" yaml:"skip_paths"`
}
