package workflows

import (
	"time"

	"github.com/billowdev/email-job-temporal/internal/adapters/temporal/activities"
	"github.com/billowdev/email-job-temporal/internal/core/domain"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// SendEmailWithTemplateTask is the implementation of the email sending workflow.
func SendEmailWithTemplateTask(ctx workflow.Context, data domain.EmailDto) error {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second * 1,
			BackoffCoefficient: 2.0,
			MaximumInterval:    time.Second * 10,
			MaximumAttempts:    3,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	err := workflow.ExecuteActivity(ctx, activities.SendEmailActivity, data).Get(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
