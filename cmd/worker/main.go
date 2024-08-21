package main

import (
	"log"
	"log/slog"

	"github.com/billowdev/email-job-temporal/internal/adapters/temporal/worker"
	"github.com/billowdev/email-job-temporal/pkg/configs"
	"go.temporal.io/sdk/client"
	temporalLog "go.temporal.io/sdk/log"
)

func main() {
	logger := temporalLog.NewStructuredLogger(slog.Default())
	hostPort := client.DefaultHostPort
	if configs.TEMPORAL_CLIENT_URL != "" {
        hostPort = configs.TEMPORAL_CLIENT_URL
    }

	temporalClient, err := client.Dial(client.Options{
		HostPort: hostPort,
		Logger:   logger,
	})

	worker.RegisterTemporalWorkflow(temporalClient)
	if err != nil {
		log.Fatal("Failed to start Temporal worker:", err)
	}

	defer temporalClient.Close()

}
