#!/bin/sh

ENDPOINT_URL="http://localhost:8000"
MINERAL_TABLE="Mineral"

aws dynamodb delete-table \
--region eu-west-1 \
--endpoint-url ${ENDPOINT_URL} \
--table-name ${MINERAL_TABLE}

aws dynamodb delete-table \
--region eu-west-1 \
--endpoint-url ${ENDPOINT_URL} \
--table-name ${MINERAL_TABLE}