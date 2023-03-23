package pkg

import "fmt"

//go:generate mockery --inpackage --name=TopicHandler
type TopicHandler interface {
	ProcssTopic1(message int) error
}

type topicHandler struct {
	magic Magic
}

func NewTopicHandler(magic Magic) TopicHandler {
	return &topicHandler{magic: magic}
}

func (t *topicHandler) ProcssTopic1(message int) error {
	book, err := t.magic.BuildMagicBook(message)
	if err != nil {
		return fmt.Errorf("BuildMagicBook BuildMagicBook error %w", err)
	}

	err = t.magic.Magic(book)
	if err != nil {
		return fmt.Errorf("BuildMagicBook Magic error %w", err)
	}

	return nil
}
