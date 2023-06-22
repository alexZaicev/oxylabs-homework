package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
	"oxylabs-homework/internal/drivers/cli"
)

const (
	shutdownTimeout = 30 * time.Second
)

func main() {
	os.Exit(run())
}

func run() int {
	msgCh := make(chan string)
	defer close(msgCh)

	// initialize application
	app, err := initialize(msgCh, msgCh)
	if err != nil {
		fmt.Printf("ERROR: failed to initialize application: %s\n", err)
		return cli.Failure
	}

	app.logger.Info("starting publisher server")
	go app.publisherServer.Start(app.logger.With(zap.Any("server", "publisher")))

	app.logger.Info("starting subscriber server")
	go app.subscriberServer.Start(app.logger.With(zap.Any("server", "subscriber")))

	// Wait for SIGTERM signal to gracefully shutdown the application
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	app.Shutdown(ctx)

	app.logger.Info("shutdown successful")

	return cli.Success
}
