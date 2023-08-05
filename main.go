package main

import (
	"context"
	"flag"
	"main/services"
	"main/utils"

	"github.com/aws/aws-sdk-go-v2/config"
)

func main() {

	action := flag.String("action", "", "")
	file := flag.String("file", "", "")
	flag.Parse()

	cfg, err := config.LoadDefaultConfig(context.TODO())
	utils.Check(err)

	if *action == "createKey" {
		services.CreateKey(cfg)
	} else if *action == "createAlias" {
		services.CreateAlias(cfg)
	} else if *action == "encryptFile" {
		services.EncryptFile(cfg, file)
	} else {
		panic("no valid option selected")
	}

}
