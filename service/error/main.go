package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"github.com/rakin92/go-sam-starter/lambdautils"
)

// ErrorHandler handles errors in the dead letter queue.
func ErrorHandler(ctx context.Context, event events.SQSEvent, svc sqsiface.SQSAPI) error {
	for _, message := range event.Records {

		// Handle the error here.
		log.Println("error handled:", message.MessageId)

		// Delete the message from the dead letter queue.
		// Currently failing to delete message on mock events
		lambdautils.DeleteMessage(svc, message.ReceiptHandle)
	}
	return nil
}

func handler(ctx context.Context, event events.SQSEvent) error {
	sess := session.Must(session.NewSession())
	svc := sqs.New(sess)
	return ErrorHandler(ctx, event, svc)
}

func main() {
	lambda.Start(handler)
}
