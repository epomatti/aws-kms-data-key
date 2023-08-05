# AWS KMS Data Key

This code will create a KMS key and generate a Data Key for encryption and decryption.

Build the executable:

```
go get
go build
```

Create the KMS key:

```sh
./main -action="createKey"
```

Create the key alias:

```sh
./main -action="createAlias"
```

Encrypt a file with **5KB** using a **Data Key** (4KB is the limit for standard keys in KMS):

```sh
./main -action="encryptFile" -file="hello5kb.txt"
```

Decrypt the file:

```sh
./main -action="decryptFile" -file="tmp/hello5kb.txt.encrypted"
```


#### Reference

- [AWS KMS Go SDK V2](https://docs.aws.amazon.com/code-library/latest/ug/go_2_kms_code_examples.html)
- [AWS SSM Go SDK V2](https://github.com/awsdocs/aws-doc-sdk-examples/blob/main/gov2/ssm/PutParameter/PutParameterv2.go)
