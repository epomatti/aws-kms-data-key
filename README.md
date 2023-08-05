# aws-kms-data-key

This code will create a KMS key and use it for encryption.

Get the dependencies:

```
go get
go build
```

Create the KMS key:

```sh
./main -action="createKey"
```

```sh
./main -action="createAlias"
```

## Reference

- [AWS KMS Go SDK V2](https://docs.aws.amazon.com/code-library/latest/ug/go_2_kms_code_examples.html)
- [AWS SSM Go SDK V2](https://github.com/awsdocs/aws-doc-sdk-examples/blob/main/gov2/ssm/PutParameter/PutParameterv2.go)
