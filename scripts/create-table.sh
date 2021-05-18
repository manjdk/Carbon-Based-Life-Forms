#!/bin/sh

ENDPOINT_URL="http://localhost:8000"
MINERAL_TABLE="Mineral"

aws dynamodb describe-table \
--region eu-west-1 \
--endpoint-url ${ENDPOINT_URL} \
--table-name ${MINERAL_TABLE} \

if [ $? -ne 0 ];
then

    aws dynamodb create-table \
    --region eu-west-1 \
    --endpoint-url ${ENDPOINT_URL} \
    --table-name ${MINERAL_TABLE} \
    --attribute-definitions \
        AttributeName=id,AttributeType=S \
        AttributeName=clientId,AttributeType=S \
    --key-schema \
        AttributeName=id,KeyType=HASH \
    --provisioned-throughput ReadCapacityUnits=10,WriteCapacityUnits=10 \
    --global-secondary-indexes '[
   {
        "IndexName": "clientId-index",
        "KeySchema": [
            {"AttributeName": "clientId", "KeyType": "HASH"}
         ],
        "Projection": {
            "ProjectionType": "ALL"
        },
        "ProvisionedThroughput": {
            "ReadCapacityUnits": 10,"WriteCapacityUnits": 10
        }
   }
]'

fi