package kmsdecrypt

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

// New takes AWS Region r and creates a KmsDecrypter. Returns new KmsDecrypter and error.
func New(r string) (*KmsDecrypter, error) {
	// If region is left empty, default to eu-west-1
	if r == "" {
		r = "eu-west-1"
	}

	// Use the default credential provider, it will check in the following order (1. ENV VARS, 2. CONFIG FILE, 3. EC2 ROLE).
	// In most cases we will use the EC2 role for providing access to the KMS key.
	cfg := &aws.Config{
		Region: &r,
	}

	// Create the session to AWS.
	sess, err := session.NewSession(cfg)
	if err != nil {
		return nil, err
	}

	return &KmsDecrypter{session: sess}, nil
}
