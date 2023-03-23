package pkg

import (
	"errors"
)

//go:generate mockery --inpackage --name=QueueHandler
type QueueHandler interface {
	ListenOnTopics(m int)
}

type queueHandler struct {
	topicHandler TopicHandler
	queueClient  Queue
}

func NewQueueHandler(topicHandler TopicHandler, queueClient Queue) QueueHandler {
	return &queueHandler{
		topicHandler: topicHandler,
		queueClient:  queueClient,
	}
}

func (q *queueHandler) ListenOnTopics(m int) {
	err := q.topicHandler.ProcssTopic1(m)
	if err != nil {
		if errors.Is(err, ErrorRecoverable) {
			q.queueClient.Pushback(m)
			return
		} else if errors.Is(err, ErrorUnrecoverable) {
			q.queueClient.Ack(m)
			return
		} else {
			// handle exceptions...
			return
		}
	}
	q.queueClient.Ack(m)
}
