package kmsdecrypt

import "github.com/aws/aws-sdk-go/aws/session"

// KmsDecrypter stores the AWS Session used for KMS decryption.
type KmsDecrypter struct {
	session *session.Session
}

type resultKeyValue struct {
	key string
	val string
	err error
}

type resultString struct {
	str string
	err error
}
