package services

import (
	"context"
	"log"

	"github.com/billowdev/email-job-temporal/internal/adapters/temporal/workflows"
	"github.com/billowdev/email-job-temporal/internal/core/domain"
	"github.com/billowdev/email-job-temporal/internal/core/ports"
	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
)

type EmailServiceImpls struct {
	wc client.Client
}

func NewEmailService(wc client.Client) ports.IEmailService {
	return &EmailServiceImpls{
		wc: wc,
	}
}

// SendEmail implements ports.IEmailService.
func (e *EmailServiceImpls) SendEmail(data domain.EmailDto) error {
	from := data.Sender
	to := data.Receiver
	subject := data.Subject
	htmlTemplate := data.HTMLTemplate
	cc := data.CC

	_ = from
	_ = to
	_ = subject
	_ = htmlTemplate
	_ = cc
	workflowID := "email_" + uuid.New().String()
	wo := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: "email_worker",
		// CronSchedule: "* * * * *",
	}
	we, err := e.wc.ExecuteWorkflow(context.Background(), wo, workflows.SendEmailWithTemplateTask, data)
	if err != nil {
		return err
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
	return nil
}
