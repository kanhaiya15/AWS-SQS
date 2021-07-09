package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	sqsqueue "github.com/kanhaiya15/AWS-SQS/sqsQueue"
)

func main() {
	sess, err := sqsqueue.Setup(&aws.Config{})
	if err != nil {
		panic(err)
	}
	sendMessageInput := &sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"SeqId": {
				DataType:    aws.String("String"),
				StringValue: aws.String("734246346234cv6c6c65465c65455cv622"),
			},
			"Component": {
				DataType:    aws.String("String"),
				StringValue: aws.String("Packet Capture1"),
			},
		},
		MessageBody: aws.String("Packet Capture"),
		QueueUrl:    aws.String(os.Getenv("AWS_SQS_PACKET_CAPTURE_QUEUE")),
	}
	smq, err := sqsqueue.SendMessage(sess, sendMessageInput)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Sent message to queue %v\n", smq)
	sqsqueue.ReceiveMessage(sess)
}
