package service

import "love_knot/internal/app/storage/cache"

var _ IEmailService = (*EmailService)(nil)

type IEmailService interface {
	Verify()
}

type EmailService struct {
	Storage *cache.EmailStorage
}

func (e *EmailService) Verify() {

}
