package pkg

//go:generate mockery --inpackage --name=Queue
type Queue interface {
	Ack(message int)
	Pushback(message int)
}

type queue struct {
}

func NewQueue() Queue {
	return &queue{}
}

func (q *queue) Ack(message int)      {}
func (q *queue) Pushback(message int) {}
