package main

import (
	"context"
	"flag"
	"main/services"

	"github.com/aws/aws-sdk-go-v2/config"
)

func main() {

	action := flag.String("action", "", "")
	flag.Parse()

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	if *action == "createKey" {
		services.CreateKey(cfg)
	} else if *action == "createAlias" {
		services.CreateAlias(cfg)
	}

}
