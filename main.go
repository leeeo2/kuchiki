package main

import (
	"flag"
	"fmt"

	_ "github.com/leeeoxeo/kon/log"
	"leeeoxeo.github.com/kuchiki/pkg/common/log"
	"leeeoxeo.github.com/kuchiki/pkg/common/setting"
	"leeeoxeo.github.com/kuchiki/pkg/models"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "./etc/config.yaml", "config path")
	flag.Parse()

	//ctx := context.Background()

	err := setting.InitGlobalConfig(configPath)
	if err != nil {
		fmt.Println("init config failed")
		panic(err)
	}

	fmt.Printf("config:%+v \n", setting.GlobalConfig())
	err = log.SetupGlobal(setting.GlobalConfig().Log)
	if err != nil {
		fmt.Println("init log failed", err)
		panic(err)
	}

	//setup models
	err = models.Setup()
	if err != nil {
		panic(err)
	}

}
