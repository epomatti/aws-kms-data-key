package services

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"main/utils"
	"os"

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
	utils.Check(err)

	keyId := *result.KeyMetadata.KeyId
	keyDescription := *result.KeyMetadata.Description
	fmt.Println(keyId)

	PutParameter(cfg, keyId, keyDescription)
}

var keyAlias string = "alias/externalKey"

func CreateAlias(cfg aws.Config) {
	aliasName := "alias/externalKey"
	key := GetKey(cfg)
	client := kms.NewFromConfig(cfg)
	input := &kms.CreateAliasInput{
		AliasName:   &aliasName,
		TargetKeyId: &key,
	}
	_, err := client.CreateAlias(context.TODO(), input)
	utils.Check(err)
}

func generateDataKey(cfg aws.Config) *kms.GenerateDataKeyOutput {
	client := kms.NewFromConfig(cfg)
	input := &kms.GenerateDataKeyInput{
		KeyId:   &keyAlias,
		KeySpec: types.DataKeySpecAes256,
	}
	response, err := client.GenerateDataKey(context.TODO(), input)
	utils.Check(err)

	return response
}

func EncryptFile(cfg aws.Config, file *string) {
	output := generateDataKey(cfg)
	plainKey := output.Plaintext
	dat, err := os.ReadFile(*file)
	utils.Check(err)

	c, err := aes.NewCipher(plainKey)
	utils.Check(err)

	gcm, err := cipher.NewGCM(c)
	utils.Check(err)

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	cipheredText := gcm.Seal(nonce, nonce, dat, nil)

	path := fmt.Sprintf("tmp/%s.encrypted", *file)

	os.Mkdir("tmp", os.ModePerm)
	os.Remove(path)

	f, err := os.Create(path)
	utils.Check(err)

	defer f.Close()

	_, err2 := f.Write(cipheredText)
	utils.Check(err2)

	fmt.Println("Created encrypted file")
}
