package sqsqueue

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func Setup(cfg *aws.Config) (*session.Session, error) {
	sess, err := session.NewSession(cfg)
	return sess, err
}
