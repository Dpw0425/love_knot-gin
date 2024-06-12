package service

import (
	"love_knot/resource"
	"love_knot/utils/email"
)

var _ ITemplateService = (*TemplateService)(nil)

type ITemplateService interface {
	LoadTemplate(data map[string]string) ([]byte, error)
}

type TemplateService struct {
}

func (t *TemplateService) LoadTemplate(data map[string]string) ([]byte, error) {
	template, err := resource.Template().ReadFile("templates/email/verify_code.tmpl")
	if err != nil {
		return nil, err
	}

	return email.RenderTemplate(template, data)
}
