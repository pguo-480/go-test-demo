package pkg

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestTopicHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(topicHandlerTestSuite))
}

type topicHandlerTestSuite struct {
	suite.Suite
	magicMock    *MockMagic
	topicHandler TopicHandler
	testMessage  int
}

func (s *topicHandlerTestSuite) SetupTest() {
	s.magicMock = &MockMagic{}
	s.testMessage = 1
	s.topicHandler = &topicHandler{
		magic: s.magicMock,
	}
}

func (s *topicHandlerTestSuite) Test_Procss_Topic1_Returns_Error_When_BuildMagicBook_Returns_Error() {
	testErr := fmt.Errorf("test error")
	s.magicMock.On("BuildMagicBook", s.testMessage).Return(nil, testErr)
	err := s.topicHandler.ProcssTopic1(s.testMessage)

	s.Assert().ErrorIs(err, testErr)
	s.magicMock.AssertCalled(s.T(), "BuildMagicBook", s.testMessage)
}

func (s *topicHandlerTestSuite) Test_ProcssTopic1_Returns_Error_When_Magic_Returns_Error() {
	testErr := fmt.Errorf("test error")
	testBook := &magicBook{
		filed1: 1,
	}
	s.magicMock.On("BuildMagicBook", s.testMessage).Return(testBook, nil)
	s.magicMock.On("Magic", testBook).Return(testErr)
	err := s.topicHandler.ProcssTopic1(s.testMessage)

	s.Assert().ErrorIs(err, testErr)
	s.magicMock.AssertCalled(s.T(), "BuildMagicBook", s.testMessage)
	s.magicMock.AssertCalled(s.T(), "Magic", testBook)
}

func (s *topicHandlerTestSuite) Test_ProcssTopic1_HappyPath() {
	testBook := &magicBook{
		filed1: 1,
	}
	s.magicMock.On("BuildMagicBook", s.testMessage).Return(testBook, nil)
	s.magicMock.On("Magic", testBook).Return(nil)
	err := s.topicHandler.ProcssTopic1(s.testMessage)

	s.Assert().NoError(err)
	s.magicMock.AssertCalled(s.T(), "BuildMagicBook", s.testMessage)
	s.magicMock.AssertCalled(s.T(), "Magic", testBook)
}
