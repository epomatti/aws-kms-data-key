package main

// Reference: https://docs.aws.amazon.com/code-library/latest/ug/go_2_kms_code_examples.html
// Reference: https://github.com/awsdocs/aws-doc-sdk-examples/blob/main/gov2/ssm/PutParameter/PutParameterv2.go
import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	kmsTypes "github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	ssmTypes "github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	// KMS
	kmsClient := kms.NewFromConfig(cfg)

	description := "external-cmk"
	tagKey := "Name"
	tagValue := "external-cmk"
	tags := []kmsTypes.Tag{{TagKey: &tagKey, TagValue: &tagValue}}

	input := &kms.CreateKeyInput{
		Description: &description,
		Tags:        tags,
	}

	result, err := kmsClient.CreateKey(context.TODO(), input)

	if err != nil {
		fmt.Println("Got error creating key:")
		fmt.Println(err)
		return
	}

	keyId := *result.KeyMetadata.KeyId
	keyDescription := *result.KeyMetadata.Description
	fmt.Println(keyId)

	// SSM
	ssmName := "key_id"
	ssmOverride := true

	ssmClient := ssm.NewFromConfig(cfg)
	ssmInput := &ssm.PutParameterInput{
		Name:        &ssmName,
		Value:       &keyId,
		Description: &keyDescription,
		Type:        ssmTypes.ParameterTypeString,
		Overwrite:   &ssmOverride,
	}
	ssmClient.PutParameter(context.TODO(), ssmInput)
}
