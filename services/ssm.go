package services

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

func PutParameter(cfg aws.Config, keyId string, keyDescription string) {
	ssmName := "key_id"
	ssmOverride := true

	ssmClient := ssm.NewFromConfig(cfg)
	ssmInput := &ssm.PutParameterInput{
		Name:        &ssmName,
		Value:       &keyId,
		Description: &keyDescription,
		Type:        types.ParameterTypeString,
		Overwrite:   &ssmOverride,
	}
	ssmClient.PutParameter(context.TODO(), ssmInput)
}
