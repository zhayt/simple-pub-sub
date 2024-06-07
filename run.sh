#!/bin/bash

usage() {
    echo "Usage: $0 [-t topic_name]"
    exit 1
}

while getopts ":t:" opt; do
    case "${opt}" in
        t)
            Topic=${OPTARG}
            ;;
        *)
            usage
            ;;
    esac
done

if [ -z "$Topic" ]; then
    echo "Enter topic name:"
    read Topic
fi

if [ -z "$Topic" ]; then
    echo "Error: Topic name is required."
    usage
fi

# Docker operations
docker compose down -v
docker compose build
docker compose up -d

until docker exec kafka kafka-topics.sh --create --topic ${Topic} --bootstrap-server localhost:9093 --partitions 1 --replication-factor 1
do
    echo "Waiting for kafka creating topic"
    sleep 4
done

echo "I'm done!"
