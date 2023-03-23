package main

import "gomock/pkg"

func main() {
	queueClient := pkg.NewQueue()
	db := pkg.NewDb()

	magic := pkg.NewMagic(db)
	topicHandler := pkg.NewTopicHandler(magic)
	queueHandler := pkg.NewQueueHandler(topicHandler, queueClient)

	queueHandler.ListenOnTopics(1)
}
