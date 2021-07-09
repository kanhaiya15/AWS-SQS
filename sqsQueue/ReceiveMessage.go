package sqsqueue

import (
	"fmt"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var (
	sqsMaxMessages     int64 = 10
	sqsPollWaitSeconds int64 = 0
)

// ReceiveMessage sends a message to an Amazon SQS queue
func ReceiveMessage(sess *session.Session) {
	svc := sqs.New(sess)
	for {
		out, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(os.Getenv("AWS_SQS_PACKET_CAPTURE_QUEUE")),
			MaxNumberOfMessages: aws.Int64(sqsMaxMessages),
			WaitTimeSeconds:     aws.Int64(sqsPollWaitSeconds),
			AttributeNames: []*string{
				aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
			},
			MessageAttributeNames: []*string{
				aws.String(sqs.QueueAttributeNameAll),
			},
		})

		if err != nil {
			fmt.Printf("failed to fetch sqs message %v", err)
			continue
		}
		var wg sync.WaitGroup
		for _, m := range out.Messages {
			wg.Add(1)
			go processAndDeleteMessage(&wg, sess, m)
		}
		wg.Wait()
	}
}

func processAndDeleteMessage(wg *sync.WaitGroup, sess *session.Session, m *sqs.Message) {
	defer wg.Done()

	fmt.Println("-------Message-------")
	fmt.Println("Body : ", *m.Body)
	fmt.Println("Attributes : ")
	for _, v := range m.Attributes {
		fmt.Printf("%v", *v)
	}
	fmt.Println("\nMessageAttributes : ")
	for k, v := range m.MessageAttributes {
		fmt.Printf("[Key : %v ] [Value : %v]\n", k, *v.StringValue)
	}
	fmt.Println("----------------------")

	DeleteMessage(sess, m)
}
