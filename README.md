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
