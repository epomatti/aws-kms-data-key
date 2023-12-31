# AWS KMS Data Key

This code will create a KMS key and generate a Data Key for encryption and decryption.

The envelope encryption process looks like this:

```mermaid
sequenceDiagram
    Client->>+AWS KMS: Create KMS Key
    Client->>+AWS KMS: Generate data key
    AWS KMS-->>-Client: Data key
    Client->>+Client: Encrypt content with data key
    Client->>+Client: Delete the unencrypted data key
    Client->>+Client: Envelope (append) encrypted data key and encrypted content
```

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
