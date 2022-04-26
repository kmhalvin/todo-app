// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	uModel "github.com/kmhalpin/todoapp/internal/model/user"
	"github.com/kmhalpin/todoapp/internal/repository/user"
	"sync"
)

// Ensure, that RepositoryMock does implement user.Repository.
// If this is not the case, regenerate this file with moq.
var _ user.Repository = &RepositoryMock{}

// RepositoryMock is a mock implementation of user.Repository.
//
// 	func TestSomethingThatUsesRepository(t *testing.T) {
//
// 		// make and configure a mocked user.Repository
// 		mockedRepository := &RepositoryMock{
// 			DeleteUserFunc: func(ctx context.Context, id string) (string, error) {
// 				panic("mock out the DeleteUser method")
// 			},
// 			GetIDByUsernameFunc: func(ctx context.Context, username string) (string, error) {
// 				panic("mock out the GetIDByUsername method")
// 			},
// 			GetPasswordByUsernameFunc: func(ctx context.Context, username string) (string, error) {
// 				panic("mock out the GetPasswordByUsername method")
// 			},
// 			GetUserFunc: func(ctx context.Context) ([]uModel.User, error) {
// 				panic("mock out the GetUser method")
// 			},
// 			GetUserByIDFunc: func(ctx context.Context, id string) (uModel.User, error) {
// 				panic("mock out the GetUserByID method")
// 			},
// 			GetUsernameIsAvailableFunc: func(ctx context.Context, username string) error {
// 				panic("mock out the GetUsernameIsAvailable method")
// 			},
// 			InsertUserFunc: func(ctx context.Context, user uModel.User) (string, error) {
// 				panic("mock out the InsertUser method")
// 			},
// 		}
//
// 		// use mockedRepository in code that requires user.Repository
// 		// and then make assertions.
//
// 	}
type RepositoryMock struct {
	// DeleteUserFunc mocks the DeleteUser method.
	DeleteUserFunc func(ctx context.Context, id string) (string, error)

	// GetIDByUsernameFunc mocks the GetIDByUsername method.
	GetIDByUsernameFunc func(ctx context.Context, username string) (string, error)

	// GetPasswordByUsernameFunc mocks the GetPasswordByUsername method.
	GetPasswordByUsernameFunc func(ctx context.Context, username string) (string, error)

	// GetUserFunc mocks the GetUser method.
	GetUserFunc func(ctx context.Context) ([]uModel.User, error)

	// GetUserByIDFunc mocks the GetUserByID method.
	GetUserByIDFunc func(ctx context.Context, id string) (uModel.User, error)

	// GetUsernameIsAvailableFunc mocks the GetUsernameIsAvailable method.
	GetUsernameIsAvailableFunc func(ctx context.Context, username string) error

	// InsertUserFunc mocks the InsertUser method.
	InsertUserFunc func(ctx context.Context, user uModel.User) (string, error)

	// calls tracks calls to the methods.
	calls struct {
		// DeleteUser holds details about calls to the DeleteUser method.
		DeleteUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
		// GetIDByUsername holds details about calls to the GetIDByUsername method.
		GetIDByUsername []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Username is the username argument value.
			Username string
		}
		// GetPasswordByUsername holds details about calls to the GetPasswordByUsername method.
		GetPasswordByUsername []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Username is the username argument value.
			Username string
		}
		// GetUser holds details about calls to the GetUser method.
		GetUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// GetUserByID holds details about calls to the GetUserByID method.
		GetUserByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
		// GetUsernameIsAvailable holds details about calls to the GetUsernameIsAvailable method.
		GetUsernameIsAvailable []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Username is the username argument value.
			Username string
		}
		// InsertUser holds details about calls to the InsertUser method.
		InsertUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// User is the user argument value.
			User uModel.User
		}
	}
	lockDeleteUser             sync.RWMutex
	lockGetIDByUsername        sync.RWMutex
	lockGetPasswordByUsername  sync.RWMutex
	lockGetUser                sync.RWMutex
	lockGetUserByID            sync.RWMutex
	lockGetUsernameIsAvailable sync.RWMutex
	lockInsertUser             sync.RWMutex
}

// DeleteUser calls DeleteUserFunc.
func (mock *RepositoryMock) DeleteUser(ctx context.Context, id string) (string, error) {
	if mock.DeleteUserFunc == nil {
		panic("RepositoryMock.DeleteUserFunc: method is nil but Repository.DeleteUser was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockDeleteUser.Lock()
	mock.calls.DeleteUser = append(mock.calls.DeleteUser, callInfo)
	mock.lockDeleteUser.Unlock()
	return mock.DeleteUserFunc(ctx, id)
}

// DeleteUserCalls gets all the calls that were made to DeleteUser.
// Check the length with:
//     len(mockedRepository.DeleteUserCalls())
func (mock *RepositoryMock) DeleteUserCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockDeleteUser.RLock()
	calls = mock.calls.DeleteUser
	mock.lockDeleteUser.RUnlock()
	return calls
}

// GetIDByUsername calls GetIDByUsernameFunc.
func (mock *RepositoryMock) GetIDByUsername(ctx context.Context, username string) (string, error) {
	if mock.GetIDByUsernameFunc == nil {
		panic("RepositoryMock.GetIDByUsernameFunc: method is nil but Repository.GetIDByUsername was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Username string
	}{
		Ctx:      ctx,
		Username: username,
	}
	mock.lockGetIDByUsername.Lock()
	mock.calls.GetIDByUsername = append(mock.calls.GetIDByUsername, callInfo)
	mock.lockGetIDByUsername.Unlock()
	return mock.GetIDByUsernameFunc(ctx, username)
}

// GetIDByUsernameCalls gets all the calls that were made to GetIDByUsername.
// Check the length with:
//     len(mockedRepository.GetIDByUsernameCalls())
func (mock *RepositoryMock) GetIDByUsernameCalls() []struct {
	Ctx      context.Context
	Username string
} {
	var calls []struct {
		Ctx      context.Context
		Username string
	}
	mock.lockGetIDByUsername.RLock()
	calls = mock.calls.GetIDByUsername
	mock.lockGetIDByUsername.RUnlock()
	return calls
}

// GetPasswordByUsername calls GetPasswordByUsernameFunc.
func (mock *RepositoryMock) GetPasswordByUsername(ctx context.Context, username string) (string, error) {
	if mock.GetPasswordByUsernameFunc == nil {
		panic("RepositoryMock.GetPasswordByUsernameFunc: method is nil but Repository.GetPasswordByUsername was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Username string
	}{
		Ctx:      ctx,
		Username: username,
	}
	mock.lockGetPasswordByUsername.Lock()
	mock.calls.GetPasswordByUsername = append(mock.calls.GetPasswordByUsername, callInfo)
	mock.lockGetPasswordByUsername.Unlock()
	return mock.GetPasswordByUsernameFunc(ctx, username)
}

// GetPasswordByUsernameCalls gets all the calls that were made to GetPasswordByUsername.
// Check the length with:
//     len(mockedRepository.GetPasswordByUsernameCalls())
func (mock *RepositoryMock) GetPasswordByUsernameCalls() []struct {
	Ctx      context.Context
	Username string
} {
	var calls []struct {
		Ctx      context.Context
		Username string
	}
	mock.lockGetPasswordByUsername.RLock()
	calls = mock.calls.GetPasswordByUsername
	mock.lockGetPasswordByUsername.RUnlock()
	return calls
}

// GetUser calls GetUserFunc.
func (mock *RepositoryMock) GetUser(ctx context.Context) ([]uModel.User, error) {
	if mock.GetUserFunc == nil {
		panic("RepositoryMock.GetUserFunc: method is nil but Repository.GetUser was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetUser.Lock()
	mock.calls.GetUser = append(mock.calls.GetUser, callInfo)
	mock.lockGetUser.Unlock()
	return mock.GetUserFunc(ctx)
}

// GetUserCalls gets all the calls that were made to GetUser.
// Check the length with:
//     len(mockedRepository.GetUserCalls())
func (mock *RepositoryMock) GetUserCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockGetUser.RLock()
	calls = mock.calls.GetUser
	mock.lockGetUser.RUnlock()
	return calls
}

// GetUserByID calls GetUserByIDFunc.
func (mock *RepositoryMock) GetUserByID(ctx context.Context, id string) (uModel.User, error) {
	if mock.GetUserByIDFunc == nil {
		panic("RepositoryMock.GetUserByIDFunc: method is nil but Repository.GetUserByID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGetUserByID.Lock()
	mock.calls.GetUserByID = append(mock.calls.GetUserByID, callInfo)
	mock.lockGetUserByID.Unlock()
	return mock.GetUserByIDFunc(ctx, id)
}

// GetUserByIDCalls gets all the calls that were made to GetUserByID.
// Check the length with:
//     len(mockedRepository.GetUserByIDCalls())
func (mock *RepositoryMock) GetUserByIDCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockGetUserByID.RLock()
	calls = mock.calls.GetUserByID
	mock.lockGetUserByID.RUnlock()
	return calls
}

// GetUsernameIsAvailable calls GetUsernameIsAvailableFunc.
func (mock *RepositoryMock) GetUsernameIsAvailable(ctx context.Context, username string) error {
	if mock.GetUsernameIsAvailableFunc == nil {
		panic("RepositoryMock.GetUsernameIsAvailableFunc: method is nil but Repository.GetUsernameIsAvailable was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Username string
	}{
		Ctx:      ctx,
		Username: username,
	}
	mock.lockGetUsernameIsAvailable.Lock()
	mock.calls.GetUsernameIsAvailable = append(mock.calls.GetUsernameIsAvailable, callInfo)
	mock.lockGetUsernameIsAvailable.Unlock()
	return mock.GetUsernameIsAvailableFunc(ctx, username)
}

// GetUsernameIsAvailableCalls gets all the calls that were made to GetUsernameIsAvailable.
// Check the length with:
//     len(mockedRepository.GetUsernameIsAvailableCalls())
func (mock *RepositoryMock) GetUsernameIsAvailableCalls() []struct {
	Ctx      context.Context
	Username string
} {
	var calls []struct {
		Ctx      context.Context
		Username string
	}
	mock.lockGetUsernameIsAvailable.RLock()
	calls = mock.calls.GetUsernameIsAvailable
	mock.lockGetUsernameIsAvailable.RUnlock()
	return calls
}

// InsertUser calls InsertUserFunc.
func (mock *RepositoryMock) InsertUser(ctx context.Context, user uModel.User) (string, error) {
	if mock.InsertUserFunc == nil {
		panic("RepositoryMock.InsertUserFunc: method is nil but Repository.InsertUser was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		User uModel.User
	}{
		Ctx:  ctx,
		User: user,
	}
	mock.lockInsertUser.Lock()
	mock.calls.InsertUser = append(mock.calls.InsertUser, callInfo)
	mock.lockInsertUser.Unlock()
	return mock.InsertUserFunc(ctx, user)
}

// InsertUserCalls gets all the calls that were made to InsertUser.
// Check the length with:
//     len(mockedRepository.InsertUserCalls())
func (mock *RepositoryMock) InsertUserCalls() []struct {
	Ctx  context.Context
	User uModel.User
} {
	var calls []struct {
		Ctx  context.Context
		User uModel.User
	}
	mock.lockInsertUser.RLock()
	calls = mock.calls.InsertUser
	mock.lockInsertUser.RUnlock()
	return calls
}
