package queue

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/manjdk/Carbon-Based-Life-Forms/domain"
	"github.com/manjdk/Carbon-Based-Life-Forms/domain/usecase"
	"github.com/manjdk/Carbon-Based-Life-Forms/log"
)

func (s *sqsClient) RunFactoryConsumer(
	meltUC usecase.MeltMineralUC,
	condenseUC usecase.CondenseMineralUC,
	fractureUC usecase.FractureMineralUC,
) {
	for {
		messages, err := s.receiveMessages()
		if err != nil {
			continue
		}

		for i := range messages {
			switch messages[i].Action {
			case domain.Melt:
				if err := meltUC.Melt(messages[i].MineralID); err != nil {
					log.ErrorZ(messages[i].TraceID, err).
						Str("mineralID", messages[i].MineralID).
						Msg("Failed to melt mineral")
				}
			case domain.Condense:
				if err := condenseUC.Condense(messages[i].MineralID); err != nil {
					log.ErrorZ(messages[i].TraceID, err).
						Str("mineralID", messages[i].MineralID).
						Msg("Failed to condense mineral")
				}
			case domain.Fracture:
				if err := fractureUC.Fracture(messages[i].MineralID); err != nil {
					log.ErrorZ(messages[i].TraceID, err).
						Str("mineralID", messages[i].MineralID).
						Msg("Failed to fracture mineral")
				}
			default:
				log.ErrorZ(messages[i].TraceID, nil).
					Str("mineralID", messages[i].MineralID).
					Msg("Unsupported mineral action")
			}

			if err := s.deleteMessage(messages[i].sqsMessage); err != nil {
				log.ErrorZ(messages[i].TraceID, err).
					Str("mineralID", messages[i].MineralID).
					Msg("Failed to delete SQS message")
			}
		}
	}
}

func (s *sqsClient) RunManagerConsumer(factorySQS Publisher) {
	for {
		messages, err := s.receiveMessages()
		if err != nil {
			continue
		}

		for i := range messages {
			if err := factorySQS.Publish(&messages[i].QueueMessage); err != nil {
				log.ErrorZ(messages[i].TraceID, err).
					Str("mineralID", messages[i].MineralID).
					Msg("Failed to publish SQS message to factory")
				continue
			}

			if err := s.deleteMessage(messages[i].sqsMessage); err != nil {
				log.ErrorZ(messages[i].TraceID, err).
					Str("mineralID", messages[i].MineralID).
					Msg("Failed to delete SQS message")
			}
		}
	}
}

func (s *sqsClient) receiveMessages() ([]consumerMessage, error) {
	msgResult, err := s.Svc.ReceiveMessage(
		&sqs.ReceiveMessageInput{
			AttributeNames: []*string{
				aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
			},
			MessageAttributeNames: []*string{
				aws.String(sqs.QueueAttributeNameAll),
			},
			QueueUrl:            s.QueueURL,
			MaxNumberOfMessages: aws.Int64(1),
			VisibilityTimeout:   aws.Int64(5),
		})
	if err != nil {
		return nil, err
	}

	return translateSQSMessagesToDomain(msgResult.Messages), nil
}

func (s *sqsClient) deleteMessage(m *sqs.Message) error {
	params := &sqs.DeleteMessageInput{
		QueueUrl:      s.QueueURL,
		ReceiptHandle: aws.String(*m.ReceiptHandle),
	}
	_, err := s.Svc.DeleteMessage(params)
	return err
}

func translateSQSMessagesToDomain(sqsMessages []*sqs.Message) []consumerMessage {
	messages := make([]consumerMessage, 0)
	for i := range sqsMessages {
		if sqsMessages[i] == nil {
			continue
		}

		msg := new(consumerMessage)
		if err := json.Unmarshal([]byte(*sqsMessages[i].Body), msg); err != nil {
			continue
		}

		msg.sqsMessage = sqsMessages[i]
		messages = append(messages, *msg)
	}

	return messages
}
