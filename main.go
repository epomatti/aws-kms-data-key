package main

// Reference: https://docs.aws.amazon.com/code-library/latest/ug/go_2_kms_code_examples.html
import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}
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

	fmt.Println(*result.KeyMetadata.KeyId)
}
