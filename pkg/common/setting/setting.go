package setting

import (
	"github.com/leeexeo/kon/log"
	"github.com/leeexeo/kon/orm"
	"github.com/spf13/viper"
)

type Config struct {
	Database orm.Config   `yaml:"Database"`
	Server   ServerConfig `yaml:"Server"`
	Log      log.Config   `yaml:"Log"`
}

type ServerConfig struct {
	ListenAddr string `yaml:"Host"`
	ListenPort int    `yaml:"Port"`
}

var globalConfig Config

func GlobalConfig() *Config {
	return &globalConfig
}

func defaultConfig() {
	globalConfig = Config{
		Database: orm.Config{
			User:        "root",
			Password:    "x",
			Host:        "127.0.0.1",
			Port:        "3306",
			Schema:      "kuchiki",
			MaxIdleConn: 20,
			MaxOpenConn: 5,
			Charset:     "utf8",
			Engine:      "InnoDB",
			Collate:     "utf8_bin",
		},
		Server: ServerConfig{
			ListenAddr: "0.0.0.0",
			ListenPort: 8888,
		},
		Log: log.Config{
			Filename:                  "/var/log/kuchiki/kuchiki.log",
			MaxSize:                   500,
			MaxAge:                    7,
			MaxBackups:                30,
			LocalTime:                 false,
			Compress:                  false,
			CallerSkip:                3,
			Level:                     "debug",
			Console:                   "stdout",
			GormLevel:                 "info",
			SqlSlowThreshold:          300,
			IgnoreRecordNotFoundError: false,
			IgnoreDuplicateError:      false,
		},
	}
}

func InitGlobal(configPath string) error {
	defaultConfig()
	v := viper.New()
	v.SetConfigFile(configPath)
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = v.Unmarshal(&globalConfig)
	if err != nil {
		panic(err)
	}
	return nil
}
