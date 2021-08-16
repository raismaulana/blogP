package envconfig

import (
	"github.com/spf13/viper"
)

type EnvConfig struct {
	DBHost       string `mapstructure:"db_host"`       //
	DBPort       string `mapstructure:"db_port"`       //
	DBName       string `mapstructure:"db_name"`       //
	DBUser       string `mapstructure:"db_user"`       //
	DBPassword   string `mapstructure:"db_password"`   //
	AppName      string `mapstructure:"app_name"`      //
	AppBaseURL   string `mapstructure:"base_url"`      //
	AppPort      string `mapstructure:"port"`          //
	SecretKey    string `mapstructure:"secretkey"`     //
	SMTPHost     string `mapstructure:"smtp_host"`     //
	SMTPPort     int    `mapstructure:"smtp_port"`     //
	SMTPSender   string `mapstructure:"smtp_sender"`   //
	SMTPEmail    string `mapstructure:"smtp_email"`    //
	SMTPPassword string `mapstructure:"smtp_password"` //
}

func NewEnvConfig() (*EnvConfig, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var env EnvConfig
	err = viper.Unmarshal(&env)
	if err != nil {
		return nil, err
	}

	return &env, nil
}
