package activities

import (
	"context"
	"fmt"

	"github.com/billowdev/email-job-temporal/internal/core/domain"
	helpers "github.com/billowdev/email-job-temporal/pkg/helpers/email"
	"go.temporal.io/sdk/activity"
	"gorm.io/gorm"
)

// EmailActivities is the interface that defines the activities related to email sending.
type EmailActivities interface {
	SendEmailActivity(ctx context.Context, data domain.EmailDto) error
}
type EmailActivitiesImpl struct {
	DB *gorm.DB
}

// (a *EmailActivitiesImpl)
func SendEmailActivity(ctx context.Context, data domain.EmailDto) error {
	from := data.Sender
	to := data.Receiver
	subject := data.Subject
	htmlTemplate := data.HTMLTemplate
	cc := data.CC

	// emailLog := models.AuditLogEmailJob{
	// 	Title:   subject,
	// 	From:    from,
	// 	To:      to,
	// 	Success: false,
	// 	Error:   nil,
	// }

	// err := a.DB.Create(&emailLog).Error
	// if err != nil {
	// 	return err
	// }

	// Send the email
	err := helpers.SendEmail(from, to, subject, htmlTemplate, cc)
	if err != nil {
		// emailLog.Success = false
		errMsg := err.Error()
		// emailLog.Error = &errMsg
		// a.DB.Save(&emailLog)
		activity.GetLogger(ctx).Info(fmt.Sprintf("Send email: from: %s to: %s Result: ERROR-%v", from, to, errMsg))
		return err
	}

	// If successful
	// emailLog.Success = true
	// a.DB.Save(&emailLog)
	activity.GetLogger(ctx).Info(fmt.Sprintf("Send email: from: %s to: %s Result: SUCCESS", from, to))
	return nil
}
