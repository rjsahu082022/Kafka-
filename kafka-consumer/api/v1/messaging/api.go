package messaging

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"kafka/golang-kafka/kafka-consumer/internal/log"
	"kafka/golang-kafka/kafka-consumer/internal/service/messaging/model"
)

func RegisterHandlers(router *gin.Engine, svc model.Service, logger log.Logger) {
	res := resource{svc, logger}
	router.GET("/records", res.Records)
}

type resource struct {
	service model.Service
	log     log.Logger
}

func (res resource) Records(c *gin.Context) {
	records, err := res.service.Get(c.Request.Context(), res.log)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(http.StatusOK, records)
}
