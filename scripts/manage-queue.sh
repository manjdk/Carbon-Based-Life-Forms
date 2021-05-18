#!/bin/sh

ENV=${ENV-local}
SQS_ENDPOINT="http://localhost:4566"
EVENTS_QUEUE="factory-events manager-events"

if [ "$1" = "-d" ] || [ "$1" = "-delete" ]; then
    echo "Deleting queues"

    for queue in ${EVENTS_QUEUE}
    do
        aws --endpoint-url=${SQS_ENDPOINT} \
        sqs delete-queue \
        --queue-url ${SQS_ENDPOINT}/queue/${queue}
    done
    exit 0
fi

echo "Creating queues"

for queue in ${EVENTS_QUEUE}
do
    aws --endpoint-url=${SQS_ENDPOINT} \
    sqs create-queue \
    --queue-name ${queue}
done
