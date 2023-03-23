// Code generated by mockery v2.16.0. DO NOT EDIT.

package pkg

import mock "github.com/stretchr/testify/mock"

// MockMagic is an autogenerated mock type for the Magic type
type MockMagic struct {
	mock.Mock
}

// BuildMagicBook provides a mock function with given fields: message
func (_m *MockMagic) BuildMagicBook(message int) (*magicBook, error) {
	ret := _m.Called(message)

	var r0 *magicBook
	if rf, ok := ret.Get(0).(func(int) *magicBook); ok {
		r0 = rf(message)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*magicBook)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(message)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Magic provides a mock function with given fields: book
func (_m *MockMagic) Magic(book *magicBook) error {
	ret := _m.Called(book)

	var r0 error
	if rf, ok := ret.Get(0).(func(*magicBook) error); ok {
		r0 = rf(book)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockMagic interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockMagic creates a new instance of MockMagic. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockMagic(t mockConstructorTestingTNewMockMagic) *MockMagic {
	mock := &MockMagic{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
