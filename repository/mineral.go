package repository

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/manjdk/Carbon-Based-Life-Forms/domain"
)

func (d *DynamoDB) Create(mineral *domain.Mineral) error {
	attributes, err := dynamodbattribute.MarshalMap(mineral)
	if err != nil {
		return err
	}

	putReq := &dynamodb.PutItemInput{
		Item:                attributes,
		TableName:           aws.String(d.MineralTable),
		ConditionExpression: aws.String("attribute_not_exists(" + keyID + ")"),
	}

	_, err = d.db.PutItem(putReq)
	return err
}

func (d *DynamoDB) GetAll() ([]domain.Mineral, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(d.MineralTable),
	}

	resp, err := d.db.Scan(input)
	if err != nil {
		return nil, err
	}

	minerals := make([]domain.Mineral, 0)
	return minerals, dynamodbattribute.UnmarshalListOfMaps(resp.Items, &minerals)
}

func (d *DynamoDB) Get(mineralID string) (*domain.Mineral, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			keyID: {
				S: aws.String(mineralID),
			},
		},
		TableName: aws.String(d.MineralTable),
	}
	resp, err := d.db.GetItem(input)
	if err != nil {
		return nil, err
	}

	sc := new(domain.Mineral)
	return sc, dynamodbattribute.UnmarshalMap(resp.Item, sc)
}

func (d *DynamoDB) GetByClientID(clientID string) ([]domain.Mineral, error) {
	input := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":ci": {
				S: aws.String(clientID),
			},
		},
		KeyConditionExpression: aws.String(fmt.Sprintf("%s =:ci", clientIDIndex)),
		TableName:              aws.String(d.MineralTable),
		IndexName:              aws.String(fmt.Sprintf("%s-index", clientIDIndex)),
	}

	resp, err := d.db.Query(input)
	if err != nil {
		return nil, err
	}

	minerals := make([]domain.Mineral, 0)
	return minerals, dynamodbattribute.UnmarshalListOfMaps(resp.Items, &minerals)
}

func (d *DynamoDB) Update(mineral *domain.Mineral) error {
	attributes, err := dynamodbattribute.MarshalMap(mineral)
	if err != nil {
		return err
	}

	putReq := &dynamodb.PutItemInput{
		Item:                attributes,
		TableName:           aws.String(d.MineralTable),
		ConditionExpression: aws.String("attribute_exists(" + keyID + ")"),
	}

	_, err = d.db.PutItem(putReq)
	return err
}

func (d *DynamoDB) Delete(mineralID string) error {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			keyID: {
				S: aws.String(mineralID),
			},
		},
		TableName:           aws.String(d.MineralTable),
		ConditionExpression: aws.String("attribute_exists(" + keyID + ")"),
	}
	_, err := d.db.DeleteItem(input)
	return err
}
