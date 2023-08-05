package services

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
)

func CreateKey(cfg aws.Config) {
	client := kms.NewFromConfig(cfg)

	description := "external-cmk"
	tagKey := "Name"
	tagValue := "external-cmk"
	tags := []types.Tag{{TagKey: &tagKey, TagValue: &tagValue}}

	input := &kms.CreateKeyInput{
		Description: &description,
		Tags:        tags,
	}

	result, err := client.CreateKey(context.TODO(), input)

	if err != nil {
		fmt.Println("Got error creating key:")
		fmt.Println(err)
		return
	}

	keyId := *result.KeyMetadata.KeyId
	keyDescription := *result.KeyMetadata.Description
	fmt.Println(keyId)

	PutParameter(cfg, keyId, keyDescription)
}

func CreateAlias(cfg aws.Config) {
	aliasName := "alias/externalKey"
	key := GetKey(cfg)
	client := kms.NewFromConfig(cfg)
	input := &kms.CreateAliasInput{
		AliasName:   &aliasName,
		TargetKeyId: &key,
	}
	_, err := client.CreateAlias(context.TODO(), input)
	if err != nil {
		panic(err)
	}
}
