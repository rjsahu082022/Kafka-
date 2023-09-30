package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	kafkaProducer "kafka/golang-kafka/kafka-producer/cmd/kafka-producer"
)

const (
	defaultConfPath = "./local.yaml"
)

func main() {
	application := &kafkaProducer.Application{}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	application.Init(ctx, defaultConfPath)
	application.Start(ctx)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sigterm
	application.Stop(ctx)

	defer func(cancel context.CancelFunc) {
		cancel()
		os.Exit(0)
	}(cancel)
}
