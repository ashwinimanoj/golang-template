package config

import (
	"errors"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Logger LoggerConfig
	Server ServerConfig
}

func NewConfig() (*Config, []error) {
	var errs []error

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	//Replace prefix with the name of the application
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		errs = append(errs, err)
		return nil, errs
	}
	if validateErrs := config.Validate(); validateErrs != nil {
		errs = append(errs, validateErrs...)
	}

	return &config, errs
}

func (c Config) Validate() []error {
	v := reflect.ValueOf(c)
	t := v.Type()
	return validate(v, t)
}

func validate(v reflect.Value, t reflect.Type) []error {
	var errs []error

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Type().Kind() == reflect.Struct {
			if validateErrs := validate(field, field.Type()); validateErrs != nil {
				errs = append(errs, validateErrs...)
			}
		}

		tag := t.Field(i).Tag.Get("validate")
		if tag == "required" {
			if field.Interface() == reflect.Zero(field.Type()).Interface() {
				errs = append(errs, errors.New(t.Field(i).Name+" is required"))
			}
		}
	}

	return errs
}
