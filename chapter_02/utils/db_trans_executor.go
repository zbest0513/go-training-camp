package utils

type TransTaskExecutor interface {
	exec() (int64, error)
}

type TransInsertTaskExecutor struct {
	method func([]interface{}) (int64, error)
	target []interface{}
}

func (receiver *TransInsertTaskExecutor) exec() (int64, error) {
	return receiver.method(receiver.target)
}

func (receiver *TransInsertTaskExecutor) NewInsertTaskExecutor(method func([]interface{}) (int64, error), target []interface{}) *TransInsertTaskExecutor {
	return &TransInsertTaskExecutor{
		target: target,
		method: method,
	}
}

type TransUpdateTaskExecutor struct {
	method func(interface{}, *WhereGenerator, []string) (int64, error)
	target interface{}
	where  *WhereGenerator
	sets   []string
}

func (receiver *TransUpdateTaskExecutor) exec() (int64, error) {
	return receiver.method(receiver.target, receiver.where, receiver.sets)
}
func (receiver *TransUpdateTaskExecutor) NewUpdateTaskExecutor(method func(interface{}, *WhereGenerator, []string) (int64, error), target interface{}, where *WhereGenerator, sets []string) *TransUpdateTaskExecutor {
	return &TransUpdateTaskExecutor{
		target: target,
		where:  where,
		sets:   sets,
		method: method,
	}
}

type TransDeleteTaskExecutor struct {
	method func(interface{}, *WhereGenerator) (int64, error)
	target interface{}
	where  *WhereGenerator
}

func (receiver *TransDeleteTaskExecutor) exec() (int64, error) {
	return receiver.method(receiver.target, receiver.where)
}
func (receiver *TransDeleteTaskExecutor) NewDeleteTaskExecutor(method func(interface{}, *WhereGenerator) (int64, error), target interface{}, where *WhereGenerator) *TransDeleteTaskExecutor {
	return &TransDeleteTaskExecutor{
		target: target,
		where:  where,
		method: method,
	}
}
