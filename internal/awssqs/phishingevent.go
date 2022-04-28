package awssqs

import (
	"encoding/json"
	"fmt"

	events "gosanta/pkg"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

//go:generate mockgen -destination=../mocks/phishingevent.go -package=mocks -source=./phishingevent.go

type SQSGetDeleter interface {
	ReceiveMessage(*sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error)
	DeleteMessage(*sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error)
}

type SQSClient struct {
	s        SQSGetDeleter
	queueURL string
}

func New(s SQSGetDeleter, queueURL string) SQSClient {
	return SQSClient{s: s, queueURL: queueURL}
}

func (sc *SQSClient) GetNextMessages() ([]events.PhishingEvent, error) {
	eventS := []events.PhishingEvent{}
	rmi := &sqs.ReceiveMessageInput{MaxNumberOfMessages: aws.Int64(10), QueueUrl: &sc.queueURL}

	rmo, err := sc.s.ReceiveMessage(rmi)
	if err != nil {
		return eventS, err
	}

	for _, msg := range rmo.Messages {
		msgBody := msg.Body
		if msgBody == nil {
			return eventS, fmt.Errorf("empty event message")
		}
		event := events.PhishingEvent{}
		err := json.Unmarshal([]byte(*msgBody), &event)
		if err != nil {
			return eventS, fmt.Errorf("could not unmarshal message to PhishingEvent: %v", err)
		}
		event.EventId = *msg.ReceiptHandle
		eventS = append(eventS, event)
	}
	return eventS, nil
}

func (sc *SQSClient) DeleteMessage(eventID string) error {
	dmi := sqs.DeleteMessageInput{QueueUrl: &sc.queueURL, ReceiptHandle: &eventID}
	_, err := sc.s.DeleteMessage(&dmi)
	if err != nil {
		return fmt.Errorf("could not delete message with id: %v: %v", eventID, err)
	}

	return nil
}
