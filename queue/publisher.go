package queue

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/manjdk/Carbon-Based-Life-Forms/api/domain"
)

func (s *sqsClient) Publish(msg *domain.QueueMessage) error {
	b, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	input := &sqs.SendMessageInput{
		MessageBody:  aws.String(string(b)),
		QueueUrl:     s.QueueURL,
		DelaySeconds: aws.Int64(5),
	}

	_, err = s.Svc.SendMessage(input)
	return err
}
