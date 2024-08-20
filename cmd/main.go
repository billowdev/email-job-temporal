package main

import (
	"fmt"

	"github.com/billowdev/email-job-temporal/internal/adapters/temporal/worker"
)

func main() {

	fmt.Println("Hello world !")

	WorkflowClient := worker.WorkflowClient()
	worker.RegisterTemporalWorkflow(WorkflowClient)
	defer WorkflowClient.Close()
	// emailSrv := services.NewEmailService(WorkflowClient)

	// err := emailSrv.SendEmail(domain.EmailDto{
	// 	Sender:       configs.SMTP_SENDER,
	// 	Receiver:     "testmail@billowdev.com",
	// 	Subject:      "TEST",
	// 	HTMLTemplate: emailtemplates.TEST_HTML_TEMPLATE,
	// 	CC:           []string{"empmail01@billowdev.com", "empmail02@billowdev.com"},
	// })
	// fmt.Println("---")
	// fmt.Println(err)
	// fmt.Println("---")
}
