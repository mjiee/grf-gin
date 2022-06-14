package conf

type Jwt struct {
	Secret string `mapstructure:"secret" json:"secret" yaml:"secret" validate:"required"`
	// expiration time, > 1h, h.
	ExpiresAt int `mapstructure:"expires_at" json:"expires_at" yaml:"expires_at" validate:"required"`
}
