package awssqs_test

import (
	"testing"
	"time"

	"gosanta/internal/awssqs"
	"gosanta/internal/mocks"
	events "gosanta/pkg"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetNextMessages(t *testing.T) {
	ctrl := gomock.NewController(t)
	sqsMock := mocks.NewMockSQSGetDeleter(ctrl)
	queueUrl := "some-queue-url"
	eventReader := awssqs.New(sqsMock, queueUrl)

	nbrM := int64(10)
	expRMI := sqs.ReceiveMessageInput{MaxNumberOfMessages: &nbrM, QueueUrl: &queueUrl}

	expBody := `{"user_id": 1, "created_at": "2022-03-28T14:41:34.0+00:00", "action": "opened", "email_ref": "f20416ef-15d5-4159-9bef-de150edfa970"}`
	expMsgID := "123-432"
	expMessages := []*sqs.Message{
		{
			Attributes:    nil,
			Body:          &expBody,
			ReceiptHandle: &expMsgID,
		},
	}
	expRMO := sqs.ReceiveMessageOutput{Messages: expMessages}

	expEvents := []events.PhishingEvent{
		{
			UserId:    1,
			CreatedAt: time.Date(2022, 3, 28, 14, 41, 34, 0, time.UTC),
			Action:    "opened",
			EmailRef:  "f20416ef-15d5-4159-9bef-de150edfa970",
			EventId:   expMsgID,
		},
	}

	sqsMock.EXPECT().ReceiveMessage(&expRMI).Return(&expRMO, nil)
	eventS, err := eventReader.GetNextMessages()

	assert.Nil(t, err)
	assert.Equal(t, expEvents[0].Action, eventS[0].Action)
	assert.Equal(t, expEvents[0].UserId, eventS[0].UserId)
	assert.Equal(t, expEvents[0].EmailRef, eventS[0].EmailRef)
	assert.Equal(t, expEvents[0].EventId, eventS[0].EventId)
}
