package queue

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/manjdk/Carbon-Based-Life-Forms/api/domain"
	"github.com/manjdk/Carbon-Based-Life-Forms/api/domain/usecase"
)

type SQSClientIFace interface {
	Publisher
	Consumer
}

type Publisher interface {
	Publish(msg *domain.QueueMessage) error
}

type Consumer interface {
	RunFactoryConsumer(
		meltUC usecase.MeltMineralUC,
		condenseUC usecase.CondenseMineralUC,
		fractureUC usecase.FractureMineralUC,
	)
	RunManagerConsumer(factorySQS Publisher)
}

type sqsClient struct {
	Svc       *sqs.SQS
	QueueURL  *string
	QueueName string
}

type consumerMessage struct {
	domain.QueueMessage
	sqsMessage *sqs.Message
}

func NewSQSClient(queueName, queueHost, awsRegion string) (SQSClientIFace, error) {
	awsConfig := aws.NewConfig().
		WithRegion(awsRegion).
		WithEndpoint(queueHost)

	awsSession, err := session.NewSession(awsConfig)
	if err != nil {
		return nil, err
	}

	svc := sqs.New(awsSession)
	queueURL, err := getSqsURL(svc, queueName)
	if err != nil {
		return nil, err
	}

	return &sqsClient{
		Svc:       svc,
		QueueURL:  queueURL,
		QueueName: queueName,
	}, nil
}

func getSqsURL(svc *sqs.SQS, queueName string) (*string, error) {
	params := &sqs.GetQueueUrlInput{QueueName: aws.String(queueName)}
	resp, err := svc.GetQueueUrl(params)
	if err != nil {
		return nil, err
	}

	return resp.QueueUrl, nil
}
