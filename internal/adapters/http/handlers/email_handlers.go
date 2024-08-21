package handlers

import (
	"github.com/billowdev/email-job-temporal/internal/core/domain"
	"github.com/billowdev/email-job-temporal/internal/core/ports"
	"github.com/billowdev/email-job-temporal/pkg/configs"
	"github.com/gofiber/fiber/v2"
)

type (
	IEmailHandler interface {
		HandleSendEmail(c *fiber.Ctx) error
	}
	EmailHandlerImpls struct {
		emailSrv ports.IEmailService
	}
)

func NewEmailHandler(
	emailSrv ports.IEmailService,
) IEmailHandler {
	return &EmailHandlerImpls{emailSrv: emailSrv}
}

// SendEmail implements IEmailHandler.
func (e *EmailHandlerImpls) HandleSendEmail(c *fiber.Ctx) error {
	var emailRequest domain.EmailDto
	// Parse the request body into the struct
	if err := c.BodyParser(&emailRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unable to parse request body",
		})
	}
	if emailRequest.Sender == "" {
		emailRequest.Sender = configs.SMTP_SENDER
	}
	err := e.emailSrv.SendEmail(domain.EmailDto{
		Sender:       emailRequest.Sender,
		Receiver:     emailRequest.Receiver,
		Subject:      emailRequest.Subject,
		HTMLTemplate: emailRequest.HTMLTemplate,
		CC:           emailRequest.CC,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{"error": err.Error()})
	}
	return c.JSON(map[string]interface{}{"status": "success"})
}
