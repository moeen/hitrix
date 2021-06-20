package workerpool

import "github.com/panjf2000/ants/v2"

type AntsWorkerPool struct {
	ants *ants.PoolWithFunc
}

func NewAntsWorkerPool(s int, f func(interface{})) (WorkerPool, error) {
	p, err := ants.NewPoolWithFunc(s, f)
	if err != nil {
		return nil, err
	}

	return &AntsWorkerPool{p}, nil
}

func (a *AntsWorkerPool) Invoke(job interface{}) error {
	return a.Invoke(job)
}
