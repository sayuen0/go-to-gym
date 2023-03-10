// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"github.com/sayuen0/go-to-gym/internal/models"
	"github.com/sayuen0/go-to-gym/internal/session"
	"sync"
)

// Ensure, that RepositoryMock does implement session.Repository.
// If this is not the case, regenerate this file with moq.
var _ session.Repository = &RepositoryMock{}

// RepositoryMock is a mock implementation of session.Repository.
//
//	func TestSomethingThatUsesRepository(t *testing.T) {
//
//		// make and configure a mocked session.Repository
//		mockedRepository := &RepositoryMock{
//			CreateFunc: func(ctx context.Context, sess *models.Session, expires int) (string, error) {
//				panic("mock out the Create method")
//			},
//			DeleteByIDFunc: func(ctx context.Context, userID string) error {
//				panic("mock out the DeleteByID method")
//			},
//			GetByIDFunc: func(ctx context.Context, id string) (*models.Session, error) {
//				panic("mock out the GetByUUID method")
//			},
//		}
//
//		// use mockedRepository in code that requires session.Repository
//		// and then make assertions.
//
//	}
type RepositoryMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(ctx context.Context, sess *models.Session, expires int) (string, error)

	// DeleteByIDFunc mocks the DeleteByID method.
	DeleteByIDFunc func(ctx context.Context, userID string) error

	// GetByIDFunc mocks the GetByID method.
	GetByIDFunc func(ctx context.Context, id string) (*models.Session, error)

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Sess is the sess argument value.
			Sess *models.Session
			// Expires is the expires argument value.
			Expires int
		}
		// DeleteByID holds details about calls to the DeleteByID method.
		DeleteByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UUID is the userID argument value.
			UserID string
		}
		// GetByID holds details about calls to the GetByID method.
		GetByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
	}
	lockCreate     sync.RWMutex
	lockDeleteByID sync.RWMutex
	lockGetByID    sync.RWMutex
}

// Create calls CreateFunc.
func (mock *RepositoryMock) Create(ctx context.Context, sess *models.Session, expires int) (string, error) {
	if mock.CreateFunc == nil {
		panic("RepositoryMock.CreateFunc: method is nil but Repository.Create was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Sess    *models.Session
		Expires int
	}{
		Ctx:     ctx,
		Sess:    sess,
		Expires: expires,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(ctx, sess, expires)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//
//	len(mockedRepository.CreateCalls())
func (mock *RepositoryMock) CreateCalls() []struct {
	Ctx     context.Context
	Sess    *models.Session
	Expires int
} {
	var calls []struct {
		Ctx     context.Context
		Sess    *models.Session
		Expires int
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// DeleteByID calls DeleteByIDFunc.
func (mock *RepositoryMock) DeleteByID(ctx context.Context, userID string) error {
	if mock.DeleteByIDFunc == nil {
		panic("RepositoryMock.DeleteByIDFunc: method is nil but Repository.DeleteByID was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		UserID string
	}{
		Ctx:    ctx,
		UserID: userID,
	}
	mock.lockDeleteByID.Lock()
	mock.calls.DeleteByID = append(mock.calls.DeleteByID, callInfo)
	mock.lockDeleteByID.Unlock()
	return mock.DeleteByIDFunc(ctx, userID)
}

// DeleteByIDCalls gets all the calls that were made to DeleteByID.
// Check the length with:
//
//	len(mockedRepository.DeleteByIDCalls())
func (mock *RepositoryMock) DeleteByIDCalls() []struct {
	Ctx    context.Context
	UserID string
} {
	var calls []struct {
		Ctx    context.Context
		UserID string
	}
	mock.lockDeleteByID.RLock()
	calls = mock.calls.DeleteByID
	mock.lockDeleteByID.RUnlock()
	return calls
}

// GetByID calls GetByIDFunc.
func (mock *RepositoryMock) GetByID(ctx context.Context, id string) (*models.Session, error) {
	if mock.GetByIDFunc == nil {
		panic("RepositoryMock.GetByIDFunc: method is nil but Repository.GetByUUID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGetByID.Lock()
	mock.calls.GetByID = append(mock.calls.GetByID, callInfo)
	mock.lockGetByID.Unlock()
	return mock.GetByIDFunc(ctx, id)
}

// GetByIDCalls gets all the calls that were made to GetByID.
// Check the length with:
//
//	len(mockedRepository.GetByIDCalls())
func (mock *RepositoryMock) GetByIDCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockGetByID.RLock()
	calls = mock.calls.GetByID
	mock.lockGetByID.RUnlock()
	return calls
}
