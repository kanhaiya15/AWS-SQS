package sqsqueue

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// DeleteMessage sends a message to an Amazon SQS queue
func DeleteMessage(sess *session.Session, m *sqs.Message) (*sqs.DeleteMessageOutput, error) {
	svc := sqs.New(sess)
	return svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      aws.String(os.Getenv("AWS_SQS_PACKET_CAPTURE_QUEUE")),
		ReceiptHandle: m.ReceiptHandle,
	})
}
