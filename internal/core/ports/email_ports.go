package ports

import "github.com/billowdev/email-job-temporal/internal/core/domain"

type IEmailService interface {
	SendEmail(data domain.EmailDto) error
}
