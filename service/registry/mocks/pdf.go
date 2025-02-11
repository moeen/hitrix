package mocks

import (
	"github.com/sarulabs/di"

	"github.com/coretrix/hitrix/service"
)

func FakeServiceTemplate(fake interface{}) *service.DefinitionGlobal {
	return &service.DefinitionGlobal{
		Name: service.PDFService,
		Build: func(ctn di.Container) (interface{}, error) {
			return fake, nil
		},
	}
}
