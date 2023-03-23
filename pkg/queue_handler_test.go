package pkg

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestQueueHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(queueHandlerTestSuite))
}

type queueHandlerTestSuite struct {
	suite.Suite
	topicHandlerMock *MockTopicHandler
	queueMock        *MockQueue
	queueHandler     *queueHandler
	testMessage      int
}

func (s *queueHandlerTestSuite) SetupTest() {
	s.topicHandlerMock = &MockTopicHandler{}
	s.queueMock = &MockQueue{}
	s.testMessage = 1
	s.queueHandler = &queueHandler{
		topicHandler: s.topicHandlerMock,
		queueClient:  s.queueMock,
	}
}

func (s *queueHandlerTestSuite) Test_ListenOnTopics_Ack_On_Unrecoverable_Error() {
	testErr := fmt.Errorf("test err %w", ErrorUnrecoverable)
	s.topicHandlerMock.On("ProcssTopic1", s.testMessage).Return(testErr)
	s.queueMock.On("Ack", s.testMessage).Return()

	s.queueHandler.ListenOnTopics(s.testMessage)

	s.topicHandlerMock.AssertCalled(s.T(), "ProcssTopic1", s.testMessage)
	s.queueMock.AssertCalled(s.T(), "Ack", s.testMessage)
	s.queueMock.AssertNotCalled(s.T(), "Pushback", mock.Anything)
}

func (s *queueHandlerTestSuite) Test_ListenOnTopics_Pushback_On_Recoverable_Error() {
	testErr := fmt.Errorf("test err %w", ErrorRecoverable)
	s.topicHandlerMock.On("ProcssTopic1", s.testMessage).Return(testErr)
	s.queueMock.On("Pushback", s.testMessage).Return()

	s.queueHandler.ListenOnTopics(s.testMessage)

	s.topicHandlerMock.AssertCalled(s.T(), "ProcssTopic1", s.testMessage)
	s.queueMock.AssertCalled(s.T(), "Pushback", s.testMessage)
	s.queueMock.AssertNotCalled(s.T(), "Ack", mock.Anything)
}

func (s *queueHandlerTestSuite) Test_ListenOnTopics_Ack_On_Success() {
	s.topicHandlerMock.On("ProcssTopic1", s.testMessage).Return(nil)
	s.queueMock.On("Ack", s.testMessage).Return()

	s.queueHandler.ListenOnTopics(s.testMessage)

	s.topicHandlerMock.AssertCalled(s.T(), "ProcssTopic1", s.testMessage)
	s.queueMock.AssertCalled(s.T(), "Ack", s.testMessage)
	s.queueMock.AssertNotCalled(s.T(), "Pushback", mock.Anything)
}
