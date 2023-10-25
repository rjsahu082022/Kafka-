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
    "message":"Hi Karan, Raj this side"
}
and sends this data into a kafka topic named "my-topic".

Any application can consume from this topic if it has required information about the topic and kafka server.

This application also supports message retry mechanism, so if a message fails to get produced then it retries 3 times before throwing an error.


Multiple Topic:-
  
  If we want to write on two topic the we have to give both the topic name in local.yaml. In this poc we are handling ony two topic. We can write on topic by using different end points 

  1. localhost:8087/message
  2. localhost:8087/message2

  Multiple Consumer:-

  If we want to read from multiple consumer then we have to start two consumer on different address. 

