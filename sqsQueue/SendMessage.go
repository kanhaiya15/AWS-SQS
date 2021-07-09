package sqsqueue

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// SendMessage sends a message to an Amazon SQS queue
func SendMessage(sess *session.Session, sendMessageInput *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
	svc := sqs.New(sess)
	return svc.SendMessage(sendMessageInput)
}
