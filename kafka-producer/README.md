# Kafka-producer

This application demonstrates a kafka producer. We are using 
[segmentio pkg](github.com/segmentio/kafka-go) library to implement a kafka producer.

### Quickstart

You need to have kafka installed locally.
#### Steps:-
1. Just clone this repository and start kafka server and zookeeper
2. Run the application using "go run cmd/main.go"
3. That's it the application is up and ready!

### Description 

We have used [GIN](https://github.com/gin-gonic/gin) web framework to implement REST API.

Currently it supports only one POST route "/message" which takes JSON body for example 
{
    "from":"Raj sahu",
     "to":"Karan",
    "message":"2015"
}
and sends this data into a kafka topic named "my-topic".

Any application can consume from this topic if it has required information about the topic and kafka server.

This application also supports message retry mechanism, so if a message fails to get produced then it retries 3 times before throwing an error.


