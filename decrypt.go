package kmsdecrypt

import (
	"encoding/base64"

	"github.com/aws/aws-sdk-go/service/kms"
)

// decryptKeyValue is ment to be run in a go-routine as sends the key, value and possibly any errors to the res channel.
func (d *KmsDecrypter) decryptKeyValue(key *string, val *string, res chan<- resultKeyValue) {
	decrypted, err := d.kmsDecrypt(*val)
	result := resultKeyValue{
		key: *key,
		val: decrypted,
	}
	// If we had any errors add the error to the result.
	if err != nil {
		result.err = err
	}

	// Send the result to the result channel.
	res <- result
}

// decryptString is ment to be run in a go-routine as sends the str and possibly any errors to the res channel.
func (d *KmsDecrypter) decryptString(str *string, res chan<- resultString) {
	decrypted, err := d.kmsDecrypt(*str)
	result := resultString{
		str: decrypted,
	}
	// If we had any errors add the error to the result.
	if err != nil {
		result.err = err
	}

	// Send the result to the result channel.
	res <- result
}

// kmsdecrypt uses aws kms to decrypt the value
func (d *KmsDecrypter) kmsDecrypt(val string) (string, error) {
	svc := kms.New(d.session)

	// Decode from base64 to []byte.
	decoded, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		return "", err
	}

	// Decrypt using KMS.
	params := &kms.DecryptInput{CiphertextBlob: decoded}
	resp, err := svc.Decrypt(params)
	if err != nil {
		return "", err
	}

	return string(resp.Plaintext), nil
}
