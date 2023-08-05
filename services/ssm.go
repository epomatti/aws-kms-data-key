package services

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

var ssmName string = "key_id"

func PutParameter(cfg aws.Config, keyId string, keyDescription string) {
	override := true

	client := ssm.NewFromConfig(cfg)
	input := &ssm.PutParameterInput{
		Name:        &ssmName,
		Value:       &keyId,
		Description: &keyDescription,
		Type:        types.ParameterTypeString,
		Overwrite:   &override,
	}
	client.PutParameter(context.TODO(), input)
}

func GetKey(cfg aws.Config) string {
	client := ssm.NewFromConfig(cfg)

	input := &ssm.GetParameterInput{
		Name: &ssmName,
	}

	response, err := client.GetParameter(context.TODO(), input)
	if err != nil {
		panic(err)
	}

	return *response.Parameter.Value
}
