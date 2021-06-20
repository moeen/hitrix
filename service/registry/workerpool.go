package registry

import (
	"github.com/coretrix/hitrix/service"
	"github.com/coretrix/hitrix/service/component/workerpool"
	"github.com/sarulabs/di"
)

func ServiceDefinitionWorkerPool(size int, poolFunc func(interface{})) *service.Definition {
	return &service.Definition{
		Name:   service.WorkerPool,
		Global: true,
		Build: func(ctn di.Container) (interface{}, error) {
			return workerpool.NewAntsWorkerPool(size, poolFunc)
		},
	}
}
