// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	models "kvado_test_task/pkg/models"

	mock "github.com/stretchr/testify/mock"
)

// LibraryManager is an autogenerated mock type for the LibraryManager type
type LibraryManager struct {
	mock.Mock
}

// GetAuthorsByBook provides a mock function with given fields: book
func (_m *LibraryManager) GetAuthorsByBook(book string) ([]models.Author, error) {
	ret := _m.Called(book)

	var r0 []models.Author
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]models.Author, error)); ok {
		return rf(book)
	}
	if rf, ok := ret.Get(0).(func(string) []models.Author); ok {
		r0 = rf(book)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Author)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(book)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBooksByAuthor provides a mock function with given fields: author
func (_m *LibraryManager) GetBooksByAuthor(author string) ([]models.Book, error) {
	ret := _m.Called(author)

	var r0 []models.Book
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]models.Book, error)); ok {
		return rf(author)
	}
	if rf, ok := ret.Get(0).(func(string) []models.Book); ok {
		r0 = rf(author)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Book)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(author)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewLibraryManager creates a new instance of LibraryManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLibraryManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *LibraryManager {
	mock := &LibraryManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
