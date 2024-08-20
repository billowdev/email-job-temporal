package main

import (
	"fmt"
	"log"
	"log/slog"
	"sync"

	"github.com/billowdev/email-job-temporal/cmd/application"
	"github.com/billowdev/email-job-temporal/internal/adapters/temporal/worker"
	"github.com/billowdev/email-job-temporal/pkg/configs"
	"go.temporal.io/sdk/client"
	temporalLog "go.temporal.io/sdk/log"
)

func main() {
	var wg sync.WaitGroup
	logger := temporalLog.NewStructuredLogger(slog.Default())
	hostPort := func() string {
		if configs.TEMPORAL_CLIENT_URL != "" {
			return configs.TEMPORAL_CLIENT_URL
		}
		return client.DefaultHostPort
	}()
	temporalClient, err := client.Dial(client.Options{
		HostPort: hostPort,
		Logger:   logger,
	})

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err != nil {
			log.Fatalln("Unable to create temporal workflow client", err)
		}

		worker.RegisterTemporalWorkflow(temporalClient)
		if err != nil {
			log.Fatal("Failed to start Temporal worker:", err)
		}

		defer temporalClient.Close()
	}()

	params := configs.NewFiberHttpServiceParams()
	fiberConfig := configs.NewFiberHTTPService(params)
	httpFiber := application.AppContainer(fiberConfig, temporalClient)
	portString := fmt.Sprintf(":%v", params.Port)
	err = httpFiber.Listen(portString)
	if err != nil {
		log.Fatal("Failed to start golang Fiber server:", err)
	}
	wg.Wait()
}
