package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/manjdk/Carbon-Based-Life-Forms/config"
	"github.com/manjdk/Carbon-Based-Life-Forms/domain"
)

const (
	keyID         = "id"
	clientIDIndex = "clientId"
)

type MineralCreateIFace interface {
	Create(mineral *domain.Mineral) error
}

type MineralsGetIFace interface {
	GetAll() ([]domain.Mineral, error)
}

type MineralsGetByClientIDIFace interface {
	GetByClientID(clientID string) ([]domain.Mineral, error)
}

type MineralGetByIDIFace interface {
	Get(mineralID string) (*domain.Mineral, error)
}

type MineralUpdateStateIFace interface {
	Update(mineral *domain.Mineral) error
}

type MineralDeleteIFace interface {
	Delete(mineralID string) error
}

type DynamoDB struct {
	db           *dynamodb.DynamoDB
	MineralTable string
}

func NewDynamoDB(config *config.Config) (*DynamoDB, error) {
	awsConfig := aws.NewConfig().
		WithRegion(config.AwsRegion).
		WithEndpoint(config.DatabaseHost)

	awsSession, err := session.NewSession(awsConfig)
	if err != nil {
		return nil, err
	}

	return &DynamoDB{
		db:           dynamodb.New(awsSession),
		MineralTable: config.MineralTable,
	}, nil
}
