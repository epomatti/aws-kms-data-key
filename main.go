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

	switch *action {
	case "createKey":
		services.CreateKey(cfg)
	case "createAlias":
		services.CreateAlias(cfg)
	case "encryptFile":
		services.EncryptFile(cfg, file)
	case "decryptFile":
		services.DecryptFile(cfg, file)
	default:
		panic("no valid action selected")
	}

}
