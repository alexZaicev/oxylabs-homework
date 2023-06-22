package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"oxylabs-homework/internal/drivers/cli"
)

const (
	shutdownTimeout = 30 * time.Second
)

func main() {
	os.Exit(run())
}

func run() int {
	app, err := initialize(context.Background())
	if err != nil {
		fmt.Printf("ERROR: failed to initialize application: %s\n", err)
		return cli.Failure
	}

	app.logger.Info("starting subscriber client")

	app.subscriberHandler.Serve()

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
