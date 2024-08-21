package routers

import "github.com/billowdev/email-job-temporal/internal/adapters/http/handlers"

func (r RouterImpls) CreateEmailRoute(h handlers.IEmailHandler) {
	r.route.Post("/emails/send",
		h.HandleSendEmail)
}
