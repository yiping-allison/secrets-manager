# Secrets Manager #

Secrets Manager is a command line tool designed to provide basic encryption/decryption functionality to a secret file located on your machine.
It is mainly used to store passwords or API keys.

The **cipher** directory is a custom package which contains the encryption and decryption functionality that the CLI uses. The **secret** directory contains the main source code which starts Cobra execution.

If you would like to add more functionality to the CLI, you can do so in `vault.go`, located in the **secret** directory.

The CLI currently supports the following functionalities:
- [x] Set - sets a key value pair and stores it in your secret file
- [x] Get - gets the value associated with the given key
- [x] List - lists the names of all the stored keys in your secret file
- [x] Remove - removes a key value pair based on given key
- [x] Delete - removes an existing Secrets file from your user directory

----

### Installation ###

Secrets Manager uses Cobra to generate the CLI tool. You also need to install cobra before building the application.

```go
go get github.com/yiping-allison/secrets-manager/secret
go get github.com/spf13/cobra/cobra
```

These commands should install the necessary dependencies that the Secrets Manager requires.

### Building ###

After installation, you should be able to see secrets-manager src files in `src/github.com/yiping-allison/secrets-manager`.

To build the application, use the following command while in `src/github.com/yiping-allison/secrets-manager/secret`:

Terminal | Command
---------|--------
**Windows Powershell** | `go build -o secret.exe cmd/cli.go`

After building, you're done! You should be able to see `secret.exe` in your current directory.

----

### First-Time Usage ###

If it's your first time running the application, make sure to set an encoding key while setting your first key-value entry. The default
option is to include no key.

_Example_

`./secret set my_API_key 'some-random-key' -k 'myEncodingKey'`

Make sure to remember your key if you set one! You need the key for subsequent uses to your secret file.

If you accidently set the wrong encoding key, you can use the `delete` command and restart your Secret file. 

----

### CLI Commands ###

Function | Command
---------|---------
Set      | `./secret set my_API_key 'some-random-key' -k 'myEncodingKey'`
Get      | `./secret get my_API_key -k 'myEncodingKey'`
List     | `./secret list -k 'myEncodingKey'`
Remove   | `./secret rm my_API_key -k 'myEncodingKey'`
Delete   | `./secret delete -k 'myEncodingKey'`

----

**Please NOTE: This was built as an exercise for me to learn Go. This is not meant to be an industry grade password/API key encryptor.
As such, please do not attempt to distribute or sell any of the code for profit.**

**This project idea was inspired by an exercise from [Gophercises](https://gophercises.com/exercises/secret).**

If you would like to learn more about Go and its different functionalities, I highly recommend going over [Jon Calhoun's lessons](https://gophercises.com/).
