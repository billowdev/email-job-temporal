// configs.go
package configs

import (
	"strconv"

	"github.com/spf13/viper"
)

var (
	TEMPORAL_CLIENT_URL string
	APP_API_VERSION     string
	APP_NAME            string
	APP_ENV             string
	APP_DEBUG_MODE      bool
	SERVER_HTTP_PORT    string

	DB_DRY_RUN       bool
	DB_NAME          string
	DB_USERNAME      string
	DB_PASSWORD      string
	DB_HOST          string
	DB_PORT          string
	DB_SSL_MODE      string
	DB_RUN_MIGRATION bool
	DB_RUN_SEEDER    bool
	DB_SCHEMA        string

	SMTP_HOST                 string
	SMTP_PORT                 string
	SMTP_USERNAME             string
	SMTP_PASSWORD             string
	SMTP_SENDER               string
	SMTP_INSECURE_SKIP_VERIFY bool
	SMTP_START_TLS            bool
	SMTP_IS_AUTH_REQUIRED     bool
)

func init() {
	NewViperConfig()

	APP_API_VERSION = "v2"
	SERVER_HTTP_PORT = viper.GetString("APP_PORT")
	TEMPORAL_CLIENT_URL = viper.GetString("TEMPORAL_CLIENT_URL")

	var err error
	APP_DEBUG_MODE, err = strconv.ParseBool(viper.GetString("APP_DEBUG_MODE"))
	if err != nil {
		APP_DEBUG_MODE = false
	}

	APP_NAME = viper.GetString("APP_NAME")
	APP_ENV = viper.GetString("APP_ENV")
	DB_DRY_RUN = false

	DB_NAME = viper.GetString("DB_NAME")
	DB_USERNAME = viper.GetString("DB_USERNAME")
	DB_PASSWORD = viper.GetString("DB_PASSWORD")
	DB_HOST = viper.GetString("DB_HOST")
	DB_PORT = viper.GetString("DB_PORT")
	DB_SSL_MODE = viper.GetString("DB_SSL_MODE")
	if DB_SSL_MODE == "" {
		DB_SSL_MODE = "default"
	}
	DB_SCHEMA = viper.GetString("DB_SCHEMA")

	DB_RUN_MIGRATION, err = strconv.ParseBool(viper.GetString("DB_RUN_MIGRATION"))
	if err != nil {
		DB_RUN_MIGRATION = false
	}

	DB_RUN_SEEDER, err = strconv.ParseBool(viper.GetString("DB_RUN_SEEDER"))
	if err != nil {
		DB_RUN_SEEDER = false
	}

	SMTP_HOST = viper.GetString("SMTP_HOST")
	SMTP_PORT = viper.GetString("SMTP_PORT")
	SMTP_USERNAME = viper.GetString("SMTP_USERNAME")
	SMTP_PASSWORD = viper.GetString("SMTP_PASSWORD")
	SMTP_SENDER = viper.GetString("SMTP_SENDER")

	SMTP_START_TLS, err = strconv.ParseBool(viper.GetString("SMTP_START_TLS"))
	if err != nil {
		SMTP_START_TLS = false
	}

	SMTP_IS_AUTH_REQUIRED, err = strconv.ParseBool(viper.GetString("SMTP_IS_AUTH_REQUIRED"))
	if err != nil {
		SMTP_IS_AUTH_REQUIRED = false
	}

	SMTP_INSECURE_SKIP_VERIFY, err = strconv.ParseBool(viper.GetString("SMTP_INSECURE_SKIP_VERIFY"))
	if err != nil {
		SMTP_INSECURE_SKIP_VERIFY = false
	}

}
