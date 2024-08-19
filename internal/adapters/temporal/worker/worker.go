package worker

import (
	"log"
	"log/slog"

	"github.com/billowdev/email-job-temporal/internal/adapters/temporal/activities"
	"github.com/billowdev/email-job-temporal/internal/adapters/temporal/workflows"
	"go.temporal.io/sdk/client"
	temporalLog "go.temporal.io/sdk/log"
	"go.temporal.io/sdk/worker"
)

func RegisterWorkflow(w worker.Registry) {
	w.RegisterWorkflow(workflows.SendEmailWithTemplateTask)
	w.RegisterActivity(activities.SendEmailActivity)
}
func WorkflowClient() client.Client {
	logger := temporalLog.NewStructuredLogger(slog.Default())
	c, err := client.Dial(client.Options{
		// HostPort: client.DefaultHostPort,
		HostPort: "localhost:7233",
		Logger:   logger,
		// HostPort: "172.16.40.38:7233",
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}

	return c
}
func RegisterTemporalWorkflow(c client.Client) {
	w := worker.New(c, "email_worker", worker.Options{})
	RegisterWorkflow(w)
	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln("Unable to start email_worker", err)
	}
	// return nil
}
