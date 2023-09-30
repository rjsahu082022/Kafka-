package kafkaproducer

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"kafka/golang-kafka/kafka-producer/internal/config"
	"kafka/golang-kafka/kafka-producer/internal/log"
	"kafka/golang-kafka/kafka-producer/internal/producer"
)

const (
	bootstrapServer = "localhost:9092"
	timeout         = 5 * time.Second
)

type Application struct {
	log        log.Logger
	router     *gin.Engine
	httpServer *http.Server
	producer1  *producer.Producer
	producer2  *producer.Producer
	cfg        *config.Config
}

func (a *Application) Init(ctx context.Context, configFile string) {
	log := log.New().With(ctx)
	a.log = log

	config, err := config.Load(log, configFile)
	if err != nil {
		log.Fatalf("failed to read config: %s ", err)
		return
	}
	a.cfg = config

	a.producer1, a.producer2 = a.BuildWriter()

	router := gin.Default()
	a.router = router
	a.SetupHandlers()
}

func (a *Application) Start(ctx context.Context) {
	a.httpServer = &http.Server{
		Addr:              fmt.Sprintf("%v", a.cfg.Server.Host) + ":" + fmt.Sprintf("%v", a.cfg.Server.Port),
		Handler:           a.router,
		ReadHeaderTimeout: timeout,
	}
	go func() {
		defer a.log.Infof("server stopped listening")
		if err := a.httpServer.ListenAndServe(); err != nil {
			a.log.Errorf("failed to listen and serve: %v ", err)
			return
		}
		return
	}()
	a.log.Infof("http server started on %s ...", a.httpServer.Addr)
}

func (a *Application) Stop(ctx context.Context) {
	err := a.httpServer.Shutdown(ctx)
	if err != nil {
		a.log.Error(err)
	}
	a.CloseWriter()
	a.log.Info("shutting down....")
	return
}
