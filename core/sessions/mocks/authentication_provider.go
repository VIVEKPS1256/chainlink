// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	auth "github.com/smartcontractkit/chainlink/v2/core/auth"
	bridges "github.com/smartcontractkit/chainlink/v2/core/bridges"

	mock "github.com/stretchr/testify/mock"

	sessions "github.com/smartcontractkit/chainlink/v2/core/sessions"
)

// AuthenticationProvider is an autogenerated mock type for the AuthenticationProvider type
type AuthenticationProvider struct {
	mock.Mock
}

// AuthorizedUserWithSession provides a mock function with given fields: sessionID
func (_m *AuthenticationProvider) AuthorizedUserWithSession(sessionID string) (sessions.User, error) {
	ret := _m.Called(sessionID)

	if len(ret) == 0 {
		panic("no return value specified for AuthorizedUserWithSession")
	}

	var r0 sessions.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (sessions.User, error)); ok {
		return rf(sessionID)
	}
	if rf, ok := ret.Get(0).(func(string) sessions.User); ok {
		r0 = rf(sessionID)
	} else {
		r0 = ret.Get(0).(sessions.User)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(sessionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClearNonCurrentSessions provides a mock function with given fields: sessionID
func (_m *AuthenticationProvider) ClearNonCurrentSessions(sessionID string) error {
	ret := _m.Called(sessionID)

	if len(ret) == 0 {
		panic("no return value specified for ClearNonCurrentSessions")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(sessionID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateAndSetAuthToken provides a mock function with given fields: user
func (_m *AuthenticationProvider) CreateAndSetAuthToken(user *sessions.User) (*auth.Token, error) {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for CreateAndSetAuthToken")
	}

	var r0 *auth.Token
	var r1 error
	if rf, ok := ret.Get(0).(func(*sessions.User) (*auth.Token, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(*sessions.User) *auth.Token); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*auth.Token)
		}
	}

	if rf, ok := ret.Get(1).(func(*sessions.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSession provides a mock function with given fields: sr
func (_m *AuthenticationProvider) CreateSession(sr sessions.SessionRequest) (string, error) {
	ret := _m.Called(sr)

	if len(ret) == 0 {
		panic("no return value specified for CreateSession")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(sessions.SessionRequest) (string, error)); ok {
		return rf(sr)
	}
	if rf, ok := ret.Get(0).(func(sessions.SessionRequest) string); ok {
		r0 = rf(sr)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(sessions.SessionRequest) error); ok {
		r1 = rf(sr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUser provides a mock function with given fields: user
func (_m *AuthenticationProvider) CreateUser(user *sessions.User) error {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*sessions.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAuthToken provides a mock function with given fields: user
func (_m *AuthenticationProvider) DeleteAuthToken(user *sessions.User) error {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAuthToken")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*sessions.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: email
func (_m *AuthenticationProvider) DeleteUser(email string) error {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUserSession provides a mock function with given fields: sessionID
func (_m *AuthenticationProvider) DeleteUserSession(sessionID string) error {
	ret := _m.Called(sessionID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUserSession")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(sessionID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindExternalInitiator provides a mock function with given fields: eia
func (_m *AuthenticationProvider) FindExternalInitiator(eia *auth.Token) (*bridges.ExternalInitiator, error) {
	ret := _m.Called(eia)

	if len(ret) == 0 {
		panic("no return value specified for FindExternalInitiator")
	}

	var r0 *bridges.ExternalInitiator
	var r1 error
	if rf, ok := ret.Get(0).(func(*auth.Token) (*bridges.ExternalInitiator, error)); ok {
		return rf(eia)
	}
	if rf, ok := ret.Get(0).(func(*auth.Token) *bridges.ExternalInitiator); ok {
		r0 = rf(eia)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bridges.ExternalInitiator)
		}
	}

	if rf, ok := ret.Get(1).(func(*auth.Token) error); ok {
		r1 = rf(eia)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUser provides a mock function with given fields: email
func (_m *AuthenticationProvider) FindUser(email string) (sessions.User, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for FindUser")
	}

	var r0 sessions.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (sessions.User, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) sessions.User); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(sessions.User)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUserByAPIToken provides a mock function with given fields: apiToken
func (_m *AuthenticationProvider) FindUserByAPIToken(apiToken string) (sessions.User, error) {
	ret := _m.Called(apiToken)

	if len(ret) == 0 {
		panic("no return value specified for FindUserByAPIToken")
	}

	var r0 sessions.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (sessions.User, error)); ok {
		return rf(apiToken)
	}
	if rf, ok := ret.Get(0).(func(string) sessions.User); ok {
		r0 = rf(apiToken)
	} else {
		r0 = ret.Get(0).(sessions.User)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(apiToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserWebAuthn provides a mock function with given fields: email
func (_m *AuthenticationProvider) GetUserWebAuthn(email string) ([]sessions.WebAuthn, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for GetUserWebAuthn")
	}

	var r0 []sessions.WebAuthn
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]sessions.WebAuthn, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) []sessions.WebAuthn); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]sessions.WebAuthn)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListUsers provides a mock function with given fields:
func (_m *AuthenticationProvider) ListUsers() ([]sessions.User, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ListUsers")
	}

	var r0 []sessions.User
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]sessions.User, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []sessions.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]sessions.User)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveWebAuthn provides a mock function with given fields: token
func (_m *AuthenticationProvider) SaveWebAuthn(token *sessions.WebAuthn) error {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for SaveWebAuthn")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*sessions.WebAuthn) error); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Sessions provides a mock function with given fields: offset, limit
func (_m *AuthenticationProvider) Sessions(offset int, limit int) ([]sessions.Session, error) {
	ret := _m.Called(offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for Sessions")
	}

	var r0 []sessions.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) ([]sessions.Session, error)); ok {
		return rf(offset, limit)
	}
	if rf, ok := ret.Get(0).(func(int, int) []sessions.Session); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]sessions.Session)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetAuthToken provides a mock function with given fields: user, token
func (_m *AuthenticationProvider) SetAuthToken(user *sessions.User, token *auth.Token) error {
	ret := _m.Called(user, token)

	if len(ret) == 0 {
		panic("no return value specified for SetAuthToken")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*sessions.User, *auth.Token) error); ok {
		r0 = rf(user, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetPassword provides a mock function with given fields: user, newPassword
func (_m *AuthenticationProvider) SetPassword(user *sessions.User, newPassword string) error {
	ret := _m.Called(user, newPassword)

	if len(ret) == 0 {
		panic("no return value specified for SetPassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*sessions.User, string) error); ok {
		r0 = rf(user, newPassword)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TestPassword provides a mock function with given fields: email, password
func (_m *AuthenticationProvider) TestPassword(email string, password string) error {
	ret := _m.Called(email, password)

	if len(ret) == 0 {
		panic("no return value specified for TestPassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(email, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateRole provides a mock function with given fields: email, newRole
func (_m *AuthenticationProvider) UpdateRole(email string, newRole string) (sessions.User, error) {
	ret := _m.Called(email, newRole)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRole")
	}

	var r0 sessions.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (sessions.User, error)); ok {
		return rf(email, newRole)
	}
	if rf, ok := ret.Get(0).(func(string, string) sessions.User); ok {
		r0 = rf(email, newRole)
	} else {
		r0 = ret.Get(0).(sessions.User)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, newRole)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAuthenticationProvider creates a new instance of AuthenticationProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthenticationProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthenticationProvider {
	mock := &AuthenticationProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
