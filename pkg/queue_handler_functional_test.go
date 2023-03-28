package pkg

import (
	"fmt"
	"testing"

	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestQueueHandlerIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(queueHandlerIntegrationTestSuite))
}

type queueHandlerIntegrationTestSuite struct {
	suite.Suite
	dbMock       *MockDb
	queueMock    *MockQueue
	magic        Magic
	topicHandler TopicHandler
	queueHandler QueueHandler
}

func (s *queueHandlerIntegrationTestSuite) SetupTest() {
	s.dbMock = &MockDb{}
	s.queueMock = &MockQueue{}
	s.magic = NewMagic(s.dbMock)
	s.topicHandler = NewTopicHandler(s.magic)
	s.queueHandler = NewQueueHandler(s.topicHandler, s.queueMock)

}

func (s *queueHandlerIntegrationTestSuite) Test_ListenOnTopics_Ack_When_Message_Is_1() {
	testMessage := 1
	s.queueMock.On("Ack", testMessage).Return()

	s.queueHandler.ListenOnTopics(testMessage)

	s.queueMock.AssertCalled(s.T(), "Ack", testMessage)
	s.queueMock.AssertNotCalled(s.T(), "Pushback", mock.Anything)
	s.dbMock.AssertNotCalled(s.T(), "Put", mock.Anything)
}

func (s *queueHandlerIntegrationTestSuite) Test_ListenOnTopics_Ack_Message_When_DB_Returns_409_error() {
	testMessage := 2
	s.queueMock.On("Ack", testMessage).Return()
	s.dbMock.On("Put", testMessage).Return(ErrorDb409)
	s.queueHandler.ListenOnTopics(testMessage)

	s.queueMock.AssertCalled(s.T(), "Ack", testMessage)
	s.queueMock.AssertNotCalled(s.T(), "Pushback", mock.Anything)
	s.dbMock.AssertCalled(s.T(), "Put", testMessage)
}

func (s *queueHandlerIntegrationTestSuite) Test_ListenOnTopics_PushBack_Message_When_DB_Returns_other_error() {
	testMessage := 2
	s.queueMock.On("Pushback", testMessage).Return()
	s.dbMock.On("Put", testMessage).Return(fmt.Errorf("other error"))
	s.queueHandler.ListenOnTopics(testMessage)

	s.queueMock.AssertNotCalled(s.T(), "Ack", mock.Anything)
	s.queueMock.AssertCalled(s.T(), "Pushback", testMessage)
	s.dbMock.AssertCalled(s.T(), "Put", testMessage)
}

func (s *queueHandlerIntegrationTestSuite) Test_ListenOnTopics_Ack_Message_When_In_HappyPath() {
	testMessage := 2
	s.queueMock.On("Ack", testMessage).Return()
	s.dbMock.On("Put", testMessage).Return(nil)
	s.queueHandler.ListenOnTopics(testMessage)

	s.queueMock.AssertCalled(s.T(), "Ack", mock.Anything)
	s.queueMock.AssertNotCalled(s.T(), "Pushback", testMessage)
	s.dbMock.AssertCalled(s.T(), "Put", testMessage)
}
