package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/leeexeo/kon/log"
	"github.com/leeexeo/kuchiki/pkg/common/setting"
	"github.com/leeexeo/kuchiki/pkg/models"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "./etc/config.yaml", "config path")
	flag.Parse()

	ctx := context.Background()

	// init global config
	fmt.Println("config path:", configPath)
	err := setting.InitGlobal(configPath)
	if err != nil {
		fmt.Println("init config failed")
		panic(err)
	}

	// setup log
	err = log.SetupGlobal(&setting.GlobalConfig().Log)
	if err != nil {
		fmt.Printf("setup global log failed,err:%s\n", err.Error())
		panic(err)
	}
	log.Debug(ctx, "setup log success", "a", "b")

	//setup models
	err = models.Setup()
	if err != nil {
		panic(err)
	}

}
