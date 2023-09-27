// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package tracker

import (
	"github.com/joshuar/go-hass-agent/internal/hass/api"
	"sync"
)

// Ensure, that apiResponseMock does implement apiResponse.
// If this is not the case, regenerate this file with moq.
var _ apiResponse = &apiResponseMock{}

// apiResponseMock is a mock implementation of apiResponse.
//
//	func TestSomethingThatUsesapiResponse(t *testing.T) {
//
//		// make and configure a mocked apiResponse
//		mockedapiResponse := &apiResponseMock{
//			DisabledFunc: func() bool {
//				panic("mock out the Disabled method")
//			},
//			RegisteredFunc: func() bool {
//				panic("mock out the Registered method")
//			},
//			TypeFunc: func() api.ResponseType {
//				panic("mock out the Type method")
//			},
//		}
//
//		// use mockedapiResponse in code that requires apiResponse
//		// and then make assertions.
//
//	}
type apiResponseMock struct {
	// DisabledFunc mocks the Disabled method.
	DisabledFunc func() bool

	// RegisteredFunc mocks the Registered method.
	RegisteredFunc func() bool

	// TypeFunc mocks the Type method.
	TypeFunc func() api.ResponseType

	// calls tracks calls to the methods.
	calls struct {
		// Disabled holds details about calls to the Disabled method.
		Disabled []struct {
		}
		// Registered holds details about calls to the Registered method.
		Registered []struct {
		}
		// Type holds details about calls to the Type method.
		Type []struct {
		}
	}
	lockDisabled   sync.RWMutex
	lockRegistered sync.RWMutex
	lockType       sync.RWMutex
}

// Disabled calls DisabledFunc.
func (mock *apiResponseMock) Disabled() bool {
	if mock.DisabledFunc == nil {
		panic("apiResponseMock.DisabledFunc: method is nil but apiResponse.Disabled was just called")
	}
	callInfo := struct {
	}{}
	mock.lockDisabled.Lock()
	mock.calls.Disabled = append(mock.calls.Disabled, callInfo)
	mock.lockDisabled.Unlock()
	return mock.DisabledFunc()
}

// DisabledCalls gets all the calls that were made to Disabled.
// Check the length with:
//
//	len(mockedapiResponse.DisabledCalls())
func (mock *apiResponseMock) DisabledCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockDisabled.RLock()
	calls = mock.calls.Disabled
	mock.lockDisabled.RUnlock()
	return calls
}

// Registered calls RegisteredFunc.
func (mock *apiResponseMock) Registered() bool {
	if mock.RegisteredFunc == nil {
		panic("apiResponseMock.RegisteredFunc: method is nil but apiResponse.Registered was just called")
	}
	callInfo := struct {
	}{}
	mock.lockRegistered.Lock()
	mock.calls.Registered = append(mock.calls.Registered, callInfo)
	mock.lockRegistered.Unlock()
	return mock.RegisteredFunc()
}

// RegisteredCalls gets all the calls that were made to Registered.
// Check the length with:
//
//	len(mockedapiResponse.RegisteredCalls())
func (mock *apiResponseMock) RegisteredCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockRegistered.RLock()
	calls = mock.calls.Registered
	mock.lockRegistered.RUnlock()
	return calls
}

// Type calls TypeFunc.
func (mock *apiResponseMock) Type() api.ResponseType {
	if mock.TypeFunc == nil {
		panic("apiResponseMock.TypeFunc: method is nil but apiResponse.Type was just called")
	}
	callInfo := struct {
	}{}
	mock.lockType.Lock()
	mock.calls.Type = append(mock.calls.Type, callInfo)
	mock.lockType.Unlock()
	return mock.TypeFunc()
}

// TypeCalls gets all the calls that were made to Type.
// Check the length with:
//
//	len(mockedapiResponse.TypeCalls())
func (mock *apiResponseMock) TypeCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockType.RLock()
	calls = mock.calls.Type
	mock.lockType.RUnlock()
	return calls
}