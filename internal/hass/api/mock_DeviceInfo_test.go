// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package api

import (
	"sync"
)

// Ensure, that DeviceInfoMock does implement DeviceInfo.
// If this is not the case, regenerate this file with moq.
var _ DeviceInfo = &DeviceInfoMock{}

// DeviceInfoMock is a mock implementation of DeviceInfo.
//
//	func TestSomethingThatUsesDeviceInfo(t *testing.T) {
//
//		// make and configure a mocked DeviceInfo
//		mockedDeviceInfo := &DeviceInfoMock{
//			AppDataFunc: func() interface{} {
//				panic("mock out the AppData method")
//			},
//			AppIDFunc: func() string {
//				panic("mock out the AppID method")
//			},
//			AppNameFunc: func() string {
//				panic("mock out the AppName method")
//			},
//			AppVersionFunc: func() string {
//				panic("mock out the AppVersion method")
//			},
//			DeviceIDFunc: func() string {
//				panic("mock out the DeviceID method")
//			},
//			DeviceNameFunc: func() string {
//				panic("mock out the DeviceName method")
//			},
//			ManufacturerFunc: func() string {
//				panic("mock out the Manufacturer method")
//			},
//			MarshalJSONFunc: func() ([]byte, error) {
//				panic("mock out the MarshalJSON method")
//			},
//			ModelFunc: func() string {
//				panic("mock out the Model method")
//			},
//			OsNameFunc: func() string {
//				panic("mock out the OsName method")
//			},
//			OsVersionFunc: func() string {
//				panic("mock out the OsVersion method")
//			},
//			SupportsEncryptionFunc: func() bool {
//				panic("mock out the SupportsEncryption method")
//			},
//		}
//
//		// use mockedDeviceInfo in code that requires DeviceInfo
//		// and then make assertions.
//
//	}
type DeviceInfoMock struct {
	// AppDataFunc mocks the AppData method.
	AppDataFunc func() interface{}

	// AppIDFunc mocks the AppID method.
	AppIDFunc func() string

	// AppNameFunc mocks the AppName method.
	AppNameFunc func() string

	// AppVersionFunc mocks the AppVersion method.
	AppVersionFunc func() string

	// DeviceIDFunc mocks the DeviceID method.
	DeviceIDFunc func() string

	// DeviceNameFunc mocks the DeviceName method.
	DeviceNameFunc func() string

	// ManufacturerFunc mocks the Manufacturer method.
	ManufacturerFunc func() string

	// MarshalJSONFunc mocks the MarshalJSON method.
	MarshalJSONFunc func() ([]byte, error)

	// ModelFunc mocks the Model method.
	ModelFunc func() string

	// OsNameFunc mocks the OsName method.
	OsNameFunc func() string

	// OsVersionFunc mocks the OsVersion method.
	OsVersionFunc func() string

	// SupportsEncryptionFunc mocks the SupportsEncryption method.
	SupportsEncryptionFunc func() bool

	// calls tracks calls to the methods.
	calls struct {
		// AppData holds details about calls to the AppData method.
		AppData []struct {
		}
		// AppID holds details about calls to the AppID method.
		AppID []struct {
		}
		// AppName holds details about calls to the AppName method.
		AppName []struct {
		}
		// AppVersion holds details about calls to the AppVersion method.
		AppVersion []struct {
		}
		// DeviceID holds details about calls to the DeviceID method.
		DeviceID []struct {
		}
		// DeviceName holds details about calls to the DeviceName method.
		DeviceName []struct {
		}
		// Manufacturer holds details about calls to the Manufacturer method.
		Manufacturer []struct {
		}
		// MarshalJSON holds details about calls to the MarshalJSON method.
		MarshalJSON []struct {
		}
		// Model holds details about calls to the Model method.
		Model []struct {
		}
		// OsName holds details about calls to the OsName method.
		OsName []struct {
		}
		// OsVersion holds details about calls to the OsVersion method.
		OsVersion []struct {
		}
		// SupportsEncryption holds details about calls to the SupportsEncryption method.
		SupportsEncryption []struct {
		}
	}
	lockAppData            sync.RWMutex
	lockAppID              sync.RWMutex
	lockAppName            sync.RWMutex
	lockAppVersion         sync.RWMutex
	lockDeviceID           sync.RWMutex
	lockDeviceName         sync.RWMutex
	lockManufacturer       sync.RWMutex
	lockMarshalJSON        sync.RWMutex
	lockModel              sync.RWMutex
	lockOsName             sync.RWMutex
	lockOsVersion          sync.RWMutex
	lockSupportsEncryption sync.RWMutex
}

// AppData calls AppDataFunc.
func (mock *DeviceInfoMock) AppData() interface{} {
	if mock.AppDataFunc == nil {
		panic("DeviceInfoMock.AppDataFunc: method is nil but DeviceInfo.AppData was just called")
	}
	callInfo := struct {
	}{}
	mock.lockAppData.Lock()
	mock.calls.AppData = append(mock.calls.AppData, callInfo)
	mock.lockAppData.Unlock()
	return mock.AppDataFunc()
}

// AppDataCalls gets all the calls that were made to AppData.
// Check the length with:
//
//	len(mockedDeviceInfo.AppDataCalls())
func (mock *DeviceInfoMock) AppDataCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockAppData.RLock()
	calls = mock.calls.AppData
	mock.lockAppData.RUnlock()
	return calls
}

// AppID calls AppIDFunc.
func (mock *DeviceInfoMock) AppID() string {
	if mock.AppIDFunc == nil {
		panic("DeviceInfoMock.AppIDFunc: method is nil but DeviceInfo.AppID was just called")
	}
	callInfo := struct {
	}{}
	mock.lockAppID.Lock()
	mock.calls.AppID = append(mock.calls.AppID, callInfo)
	mock.lockAppID.Unlock()
	return mock.AppIDFunc()
}

// AppIDCalls gets all the calls that were made to AppID.
// Check the length with:
//
//	len(mockedDeviceInfo.AppIDCalls())
func (mock *DeviceInfoMock) AppIDCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockAppID.RLock()
	calls = mock.calls.AppID
	mock.lockAppID.RUnlock()
	return calls
}

// AppName calls AppNameFunc.
func (mock *DeviceInfoMock) AppName() string {
	if mock.AppNameFunc == nil {
		panic("DeviceInfoMock.AppNameFunc: method is nil but DeviceInfo.AppName was just called")
	}
	callInfo := struct {
	}{}
	mock.lockAppName.Lock()
	mock.calls.AppName = append(mock.calls.AppName, callInfo)
	mock.lockAppName.Unlock()
	return mock.AppNameFunc()
}

// AppNameCalls gets all the calls that were made to AppName.
// Check the length with:
//
//	len(mockedDeviceInfo.AppNameCalls())
func (mock *DeviceInfoMock) AppNameCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockAppName.RLock()
	calls = mock.calls.AppName
	mock.lockAppName.RUnlock()
	return calls
}

// AppVersion calls AppVersionFunc.
func (mock *DeviceInfoMock) AppVersion() string {
	if mock.AppVersionFunc == nil {
		panic("DeviceInfoMock.AppVersionFunc: method is nil but DeviceInfo.AppVersion was just called")
	}
	callInfo := struct {
	}{}
	mock.lockAppVersion.Lock()
	mock.calls.AppVersion = append(mock.calls.AppVersion, callInfo)
	mock.lockAppVersion.Unlock()
	return mock.AppVersionFunc()
}

// AppVersionCalls gets all the calls that were made to AppVersion.
// Check the length with:
//
//	len(mockedDeviceInfo.AppVersionCalls())
func (mock *DeviceInfoMock) AppVersionCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockAppVersion.RLock()
	calls = mock.calls.AppVersion
	mock.lockAppVersion.RUnlock()
	return calls
}

// DeviceID calls DeviceIDFunc.
func (mock *DeviceInfoMock) DeviceID() string {
	if mock.DeviceIDFunc == nil {
		panic("DeviceInfoMock.DeviceIDFunc: method is nil but DeviceInfo.DeviceID was just called")
	}
	callInfo := struct {
	}{}
	mock.lockDeviceID.Lock()
	mock.calls.DeviceID = append(mock.calls.DeviceID, callInfo)
	mock.lockDeviceID.Unlock()
	return mock.DeviceIDFunc()
}

// DeviceIDCalls gets all the calls that were made to DeviceID.
// Check the length with:
//
//	len(mockedDeviceInfo.DeviceIDCalls())
func (mock *DeviceInfoMock) DeviceIDCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockDeviceID.RLock()
	calls = mock.calls.DeviceID
	mock.lockDeviceID.RUnlock()
	return calls
}

// DeviceName calls DeviceNameFunc.
func (mock *DeviceInfoMock) DeviceName() string {
	if mock.DeviceNameFunc == nil {
		panic("DeviceInfoMock.DeviceNameFunc: method is nil but DeviceInfo.DeviceName was just called")
	}
	callInfo := struct {
	}{}
	mock.lockDeviceName.Lock()
	mock.calls.DeviceName = append(mock.calls.DeviceName, callInfo)
	mock.lockDeviceName.Unlock()
	return mock.DeviceNameFunc()
}

// DeviceNameCalls gets all the calls that were made to DeviceName.
// Check the length with:
//
//	len(mockedDeviceInfo.DeviceNameCalls())
func (mock *DeviceInfoMock) DeviceNameCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockDeviceName.RLock()
	calls = mock.calls.DeviceName
	mock.lockDeviceName.RUnlock()
	return calls
}

// Manufacturer calls ManufacturerFunc.
func (mock *DeviceInfoMock) Manufacturer() string {
	if mock.ManufacturerFunc == nil {
		panic("DeviceInfoMock.ManufacturerFunc: method is nil but DeviceInfo.Manufacturer was just called")
	}
	callInfo := struct {
	}{}
	mock.lockManufacturer.Lock()
	mock.calls.Manufacturer = append(mock.calls.Manufacturer, callInfo)
	mock.lockManufacturer.Unlock()
	return mock.ManufacturerFunc()
}

// ManufacturerCalls gets all the calls that were made to Manufacturer.
// Check the length with:
//
//	len(mockedDeviceInfo.ManufacturerCalls())
func (mock *DeviceInfoMock) ManufacturerCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockManufacturer.RLock()
	calls = mock.calls.Manufacturer
	mock.lockManufacturer.RUnlock()
	return calls
}

// MarshalJSON calls MarshalJSONFunc.
func (mock *DeviceInfoMock) MarshalJSON() ([]byte, error) {
	if mock.MarshalJSONFunc == nil {
		panic("DeviceInfoMock.MarshalJSONFunc: method is nil but DeviceInfo.MarshalJSON was just called")
	}
	callInfo := struct {
	}{}
	mock.lockMarshalJSON.Lock()
	mock.calls.MarshalJSON = append(mock.calls.MarshalJSON, callInfo)
	mock.lockMarshalJSON.Unlock()
	return mock.MarshalJSONFunc()
}

// MarshalJSONCalls gets all the calls that were made to MarshalJSON.
// Check the length with:
//
//	len(mockedDeviceInfo.MarshalJSONCalls())
func (mock *DeviceInfoMock) MarshalJSONCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockMarshalJSON.RLock()
	calls = mock.calls.MarshalJSON
	mock.lockMarshalJSON.RUnlock()
	return calls
}

// Model calls ModelFunc.
func (mock *DeviceInfoMock) Model() string {
	if mock.ModelFunc == nil {
		panic("DeviceInfoMock.ModelFunc: method is nil but DeviceInfo.Model was just called")
	}
	callInfo := struct {
	}{}
	mock.lockModel.Lock()
	mock.calls.Model = append(mock.calls.Model, callInfo)
	mock.lockModel.Unlock()
	return mock.ModelFunc()
}

// ModelCalls gets all the calls that were made to Model.
// Check the length with:
//
//	len(mockedDeviceInfo.ModelCalls())
func (mock *DeviceInfoMock) ModelCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockModel.RLock()
	calls = mock.calls.Model
	mock.lockModel.RUnlock()
	return calls
}

// OsName calls OsNameFunc.
func (mock *DeviceInfoMock) OsName() string {
	if mock.OsNameFunc == nil {
		panic("DeviceInfoMock.OsNameFunc: method is nil but DeviceInfo.OsName was just called")
	}
	callInfo := struct {
	}{}
	mock.lockOsName.Lock()
	mock.calls.OsName = append(mock.calls.OsName, callInfo)
	mock.lockOsName.Unlock()
	return mock.OsNameFunc()
}

// OsNameCalls gets all the calls that were made to OsName.
// Check the length with:
//
//	len(mockedDeviceInfo.OsNameCalls())
func (mock *DeviceInfoMock) OsNameCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockOsName.RLock()
	calls = mock.calls.OsName
	mock.lockOsName.RUnlock()
	return calls
}

// OsVersion calls OsVersionFunc.
func (mock *DeviceInfoMock) OsVersion() string {
	if mock.OsVersionFunc == nil {
		panic("DeviceInfoMock.OsVersionFunc: method is nil but DeviceInfo.OsVersion was just called")
	}
	callInfo := struct {
	}{}
	mock.lockOsVersion.Lock()
	mock.calls.OsVersion = append(mock.calls.OsVersion, callInfo)
	mock.lockOsVersion.Unlock()
	return mock.OsVersionFunc()
}

// OsVersionCalls gets all the calls that were made to OsVersion.
// Check the length with:
//
//	len(mockedDeviceInfo.OsVersionCalls())
func (mock *DeviceInfoMock) OsVersionCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockOsVersion.RLock()
	calls = mock.calls.OsVersion
	mock.lockOsVersion.RUnlock()
	return calls
}

// SupportsEncryption calls SupportsEncryptionFunc.
func (mock *DeviceInfoMock) SupportsEncryption() bool {
	if mock.SupportsEncryptionFunc == nil {
		panic("DeviceInfoMock.SupportsEncryptionFunc: method is nil but DeviceInfo.SupportsEncryption was just called")
	}
	callInfo := struct {
	}{}
	mock.lockSupportsEncryption.Lock()
	mock.calls.SupportsEncryption = append(mock.calls.SupportsEncryption, callInfo)
	mock.lockSupportsEncryption.Unlock()
	return mock.SupportsEncryptionFunc()
}

// SupportsEncryptionCalls gets all the calls that were made to SupportsEncryption.
// Check the length with:
//
//	len(mockedDeviceInfo.SupportsEncryptionCalls())
func (mock *DeviceInfoMock) SupportsEncryptionCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockSupportsEncryption.RLock()
	calls = mock.calls.SupportsEncryption
	mock.lockSupportsEncryption.RUnlock()
	return calls
}