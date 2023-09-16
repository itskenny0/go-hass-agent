// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package tracker

import (
	"sync"
)

// Ensure, that agentConfigMock does implement agentConfig.
// If this is not the case, regenerate this file with moq.
var _ agentConfig = &agentConfigMock{}

// agentConfigMock is a mock implementation of agentConfig.
//
//	func TestSomethingThatUsesagentConfig(t *testing.T) {
//
//		// make and configure a mocked agentConfig
//		mockedagentConfig := &agentConfigMock{
//			GetFunc: func(s string, ifaceVal interface{}) error {
//				panic("mock out the Get method")
//			},
//			StoragePathFunc: func(s string) (string, error) {
//				panic("mock out the StoragePath method")
//			},
//		}
//
//		// use mockedagentConfig in code that requires agentConfig
//		// and then make assertions.
//
//	}
type agentConfigMock struct {
	// GetFunc mocks the Get method.
	GetFunc func(s string, ifaceVal interface{}) error

	// StoragePathFunc mocks the StoragePath method.
	StoragePathFunc func(s string) (string, error)

	// calls tracks calls to the methods.
	calls struct {
		// Get holds details about calls to the Get method.
		Get []struct {
			// S is the s argument value.
			S string
			// IfaceVal is the ifaceVal argument value.
			IfaceVal interface{}
		}
		// StoragePath holds details about calls to the StoragePath method.
		StoragePath []struct {
			// S is the s argument value.
			S string
		}
	}
	lockGet         sync.RWMutex
	lockStoragePath sync.RWMutex
}

// Get calls GetFunc.
func (mock *agentConfigMock) Get(s string, ifaceVal interface{}) error {
	if mock.GetFunc == nil {
		panic("agentConfigMock.GetFunc: method is nil but agentConfig.Get was just called")
	}
	callInfo := struct {
		S        string
		IfaceVal interface{}
	}{
		S:        s,
		IfaceVal: ifaceVal,
	}
	mock.lockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	mock.lockGet.Unlock()
	return mock.GetFunc(s, ifaceVal)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//
//	len(mockedagentConfig.GetCalls())
func (mock *agentConfigMock) GetCalls() []struct {
	S        string
	IfaceVal interface{}
} {
	var calls []struct {
		S        string
		IfaceVal interface{}
	}
	mock.lockGet.RLock()
	calls = mock.calls.Get
	mock.lockGet.RUnlock()
	return calls
}

// StoragePath calls StoragePathFunc.
func (mock *agentConfigMock) StoragePath(s string) (string, error) {
	if mock.StoragePathFunc == nil {
		panic("agentConfigMock.StoragePathFunc: method is nil but agentConfig.StoragePath was just called")
	}
	callInfo := struct {
		S string
	}{
		S: s,
	}
	mock.lockStoragePath.Lock()
	mock.calls.StoragePath = append(mock.calls.StoragePath, callInfo)
	mock.lockStoragePath.Unlock()
	return mock.StoragePathFunc(s)
}

// StoragePathCalls gets all the calls that were made to StoragePath.
// Check the length with:
//
//	len(mockedagentConfig.StoragePathCalls())
func (mock *agentConfigMock) StoragePathCalls() []struct {
	S string
} {
	var calls []struct {
		S string
	}
	mock.lockStoragePath.RLock()
	calls = mock.calls.StoragePath
	mock.lockStoragePath.RUnlock()
	return calls
}