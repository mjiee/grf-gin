package conf

// ali oss
type Oss struct {
	Region          string `mapstructure:"region" json:"region" yaml:"region" validate:"required"`
	Endpoint        string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint" validate:"required"`
	Bucket          string `mapstructure:"bucket" json:"bucket" yaml:"bucket" validate:"required"`
	AccessKeyId     string `mapstructure:"access_key_id" json:"access_key_id" yaml:"access_key_id" validate:"required"`
	AccessKeySecret string `mapstructure:"access_key_secret" json:"access_key_secret" yaml:"access_key_secret" validate:"required"`
	// sts角色ARN
	RoleArn string `mapstructure:"role_arn" json:"role_arn" yaml:"role_arn" validate:"required"`
	// 自定义角色会话名称，用来区分不同的令牌
	RoleSessionName string `mapstructure:"role_session_name" json:"role_session_name" yaml:"role_session_name" validate:"required"`
}
