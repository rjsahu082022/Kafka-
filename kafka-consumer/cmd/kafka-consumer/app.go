package kafkaconsumer

import (
	"context"
	"fmt"
	"kafka/golang-kafka/kafka-consumer/internal/config"
	"kafka/golang-kafka/kafka-consumer/internal/consumer"
	"kafka/golang-kafka/kafka-consumer/internal/log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/sync/errgroup"
)

const (
	colName = "Messages"
)

var collection *mongo.Collection

type Application struct {
	db         *mongo.Collection
	log        log.Logger
	cfg        *config.Config
	services   *services
	router     *gin.Engine
	consumer   *consumer.Consumer
	httpServer *http.Server
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

	clientOption := options.Client().ApplyURI(a.cfg.Uri)
	client, err := mongo.Connect(context.Background(), clientOption)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database(a.cfg.Name).Collection(colName)

	a.db = collection

	services := buildServices(config, collection)
	a.services = services

	router := gin.Default()
	a.router = router
	a.SetupHandlers()

	consumer, err := a.BuildReader()

	if err != nil {
		a.log.Fatalf("failed to initialize consumer: %s ", err)
		return
	}

	a.consumer = consumer
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		err = a.consumer.Consume(ctx, []string{a.cfg.Consumer.Address})
		if err != nil {
			a.log.Fatal(err)
			return err
		}
		return nil
	})

}

func (a *Application) Start(ctx context.Context) {
	a.httpServer = &http.Server{
		Addr:    fmt.Sprintf("%v", a.cfg.Server.Host) + ":" + fmt.Sprintf("%v", a.cfg.Server.Port),
		Handler: a.router,
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
	a.log.Info("shutting down....")
	return
}
