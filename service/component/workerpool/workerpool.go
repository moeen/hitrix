package workerpool

type WorkerPool interface {
	Invoke(job interface{}) error
}
