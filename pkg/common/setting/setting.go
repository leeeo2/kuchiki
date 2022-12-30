package setting

import (
	"github.com/spf13/viper"
	"leeeoxeo.github.com/kuchiki/pkg/common/log"
)

type Config struct {
	Database Database     `yaml:"Database"`
	Server   ServerConfig `yaml:"Server"`
	Log      *log.Config  `yaml:"Log"`
}

type ServerConfig struct {
	ListenAddr string `yaml:"ListenAddr"`
	ListenPort int    `yaml:"ListenPort"`
}

type Database struct {
	//Type        string `yaml:"Type"`
	User        string `yaml:"User"`
	Password    string `yaml:"Password"`
	Host        string `yaml:"Host"`
	Address     string `yaml:"Address"`
	Port        string `yaml:"Port"`
	Schema      string `yaml:"Schema"`
	MaxIdleConn int    `yaml:"MaxIdleConn"`
	MaxOpenConn int    `yaml:"MaxOpenConn"`
	Charset     string `yaml:"Charset"`
	Engine      string `yaml:"Engine"`
	Collate     string `yaml:"Collate"`
	LogLevel    string `yaml:"LogLevel"`
}

var globalConfig Config

func GlobalConfig() *Config {
	return &globalConfig
}

func defaultConfig() {
	globalConfig = Config{
		Database: Database{
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
			LogLevel:    "DEBUG",
		},
		Server: ServerConfig{
			ListenAddr: "0.0.0.0",
			ListenPort: 8888,
		},
	}
}

func InitGlobalConfig(configPath string) error {
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
