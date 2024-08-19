package main

import (
	"fmt"

	"github.com/billowdev/email-job-temporal/internal/adapters/temporal/worker"
	"github.com/billowdev/email-job-temporal/internal/core/domain"
	"github.com/billowdev/email-job-temporal/internal/core/services"
	"github.com/billowdev/email-job-temporal/pkg/configs"
	emailtemplates "github.com/billowdev/email-job-temporal/pkg/configs/email_templates"
)

func main() {

	fmt.Println("Hello world !")
	WorkflowClient := worker.WorkflowClient()
	worker.RegisterTemporalWorkflow(WorkflowClient)
	defer WorkflowClient.Close()
	fmt.Println("---")
	fmt.Println(configs.SMTP_HOST)
	fmt.Println(configs.SMTP_SENDER)
	fmt.Println(configs.SMTP_PASSWORD)
	fmt.Println("---")
	emailSrv := services.NewEmailService(WorkflowClient)
	_ = emailSrv

	err := emailSrv.SendEmail(domain.EmailDto{
		Sender:       configs.SMTP_SENDER,
		Receiver:     "sender_test@billowdev.com",
		Subject:      "TEST",
		HTMLTemplate: emailtemplates.TEST_HTML_TEMPLATE,
		CC:           []string{"cctest1@billowdev.com", "cctest2@billowdev.com"},
	})
	fmt.Println("---")
	fmt.Println(err)
	fmt.Println("---")
}
