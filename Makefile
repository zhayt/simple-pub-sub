build:
	./run.sh -t log

run-producer:
	go run producer/main.go --bootstrap-server localhost:9093 --topic log

run-consumer:
	go run consumer/main.go --bootstrap-server localhost:9093 --topic log