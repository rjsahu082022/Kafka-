package message

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"kafka/golang-kafka/kafka-producer/internal/log"
	"kafka/golang-kafka/kafka-producer/internal/producer"
	"kafka/golang-kafka/kafka-producer/internal/producer/model"
)

func RegisterHandlers(router *gin.Engine, producer, producer2 producer.Produce, logger log.Logger) {
	res := resource{producer, producer2, logger}
	router.POST("/message", res.Entry)
	router.POST("/message2", res.Messaging)

}

type resource struct {
	producer1 producer.Produce
	producer2 producer.Produce
	log       log.Logger
}

func (res resource) Entry(c *gin.Context) {
	var message model.Message
	if err := c.BindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	messageBytes, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error marshalling message:", err)
		return
	}

	err = res.producer1.Produce(messageBytes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Message received!"})
}

func (res resource) Messaging(c *gin.Context) {
	var message model.Message
	if err := c.BindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	messageBytes, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error marshalling message:", err)
		return
	}

	err = res.producer2.Produce(messageBytes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Message received!"})
}
