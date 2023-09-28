# Kafka-consumer

This application demonstrates how kafka consumer works. We are using 
[segmentio pkg](github.com/segmentio/kafka-go) library to implement a kafka producer.

### Quickstart

You need to have mongodb installed locally as this app using it.
#### Steps:-
1. Just clone this repository.
2. Run the application using "go run cmd/main.go"
3. That's it the application is up and ready!

### Description 

We have used [GIN](https://github.com/gin-gonic/gin) web framework to implement REST API.

This application has consumer group called "group1" and this group of consumer consumes messages from a remote kafka producer and perists it in Mongodb database.

Currently it supports only one GET route "/records" which fetches persisted data from database.

It also has a custom logger to log different errors and events.


