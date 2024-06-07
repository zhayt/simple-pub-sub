# Kafka Getting Started
Simple pub sub with kafka
# How to Run
### With Makefile
1. Deploy kafka and zookeeper
```shell
make build
```
2. Start producer
```shell
make run-producer
```
3. Start consumer
```shell
make run-consumer
```
### Manually
1. Deploy kafka and zookeeper
```shell
./run.sh
```
2. Start producer
```shell
go run producer/main.go --bootstrap-server localhost:9093 --topic {Topic}
```
3. Start consumer
```shell
go run consumer/main.go --bootstrap-server localhost:9093 --topic log
```
