# kmsdecrypt

[![Go Report Card](https://goreportcard.com/badge/github.com/dwtechnologies/kmsdecrypt)](https://goreportcard.com/report/github.com/dwtechnologies/kmsdecrypt)

This package can be used to quickly decrypt one or more strings using AWS KMS.
It also supports decrypting ENV variables automatically or based on ENV Key name.

## Download

`go get -u github.com/dwtechnologies/kmsdecrypt`

## Functions

### New

    New takes AWS Region r and creates a KmsDecrypter. Returns new KmsDecrypter and error.

### KmsDecrypter.DecryptEnv

    DecryptEnv will return a map[string]string decrypted Key-Value pairs from ENV variables that includes the supplied Marker m in it's Key.
    If m is an empty string it will default to "KMS_DECRYPT". The returned maps Key will have the marker removed from it's Key-name.
    Returns map[string]string and error.

### KmsDecrypter.DecryptEnvAuto

    DecryptEnvAuto will decrypt values from ENV variables automatically. It will check that the ENV value is divisible by 4, otherwise the ENV vill be ignored.
    Please note that any KMS decryption error will be treated as the ENV was not encrypted. So it can potentially be dangeroud, so use with causion.
    Returns a map[string]string of decrypted Key-Value pairs.

### KmsDecrypter.DecryptMap

    DecryptMap will decrypt values from map[string]string m.
    Returns a map[string]string of decrypted Key-Value pairs.

### KmsDecrypter.DecryptStringSlice

    DecryptStringSlice will decrypt values from []string s.
    Returns a []string of decrypted strings.

### KmsDecrypter.DecryptString

    DecryptString will string s. Returns a decrypted string.

## Example

```go
package main

import (
    "fmt"
    "strings"

    "github.com/dwtechnologies/kmsdecrypt"
)

func main() {
    // Create the decrypter for AWS Region eu-west-1
    decrypt, err := kmsdecrypt.New("eu-west-1")

    // Decrypt all ENV variables that has "KMS_DECRYPT" in their name and return them in Key-Value map.
    envs := decrypt.DecryptEnv("KMS_DECRYPT")
    fmt.Println(envs)

    // Decrypt a slice of strings and return an decrypted slice.
    slice := decrypt.DecryptStringSlice([]string{"ENCRYPTED1", "ENCRYPTED2", "ENCRYPTED3"})
    fmt.Println(slice)
}
```