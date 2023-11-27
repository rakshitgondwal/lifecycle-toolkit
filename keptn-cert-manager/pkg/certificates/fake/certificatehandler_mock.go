// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package fake

import (
	"crypto/x509"
	"encoding/pem"
	"sync"
)

// ICertificateHandlerMock is a mock implementation of certificates.ICertificateHandler.
//
//	func TestSomethingThatUsesICertificateHandler(t *testing.T) {
//
//		// make and configure a mocked certificates.ICertificateHandler
//		mockedICertificateHandler := &ICertificateHandlerMock{
//			DecodeFunc: func(data []byte) (*pem.Block, []byte) {
//				panic("mock out the Decode method")
//			},
//			ParseFunc: func(der []byte) (*x509.Certificate, error) {
//				panic("mock out the Parse method")
//			},
//		}
//
//		// use mockedICertificateHandler in code that requires certificates.ICertificateHandler
//		// and then make assertions.
//
//	}
type ICertificateHandlerMock struct {
	// DecodeFunc mocks the Decode method.
	DecodeFunc func(data []byte) (*pem.Block, []byte)

	// ParseFunc mocks the Parse method.
	ParseFunc func(der []byte) (*x509.Certificate, error)

	// calls tracks calls to the methods.
	calls struct {
		// Decode holds details about calls to the Decode method.
		Decode []struct {
			// Data is the data argument value.
			Data []byte
		}
		// Parse holds details about calls to the Parse method.
		Parse []struct {
			// Der is the der argument value.
			Der []byte
		}
	}
	lockDecode sync.RWMutex
	lockParse  sync.RWMutex
}

// Decode calls DecodeFunc.
func (mock *ICertificateHandlerMock) Decode(data []byte) (*pem.Block, []byte) {
	if mock.DecodeFunc == nil {
		panic("ICertificateHandlerMock.DecodeFunc: method is nil but ICertificateHandler.Decode was just called")
	}
	callInfo := struct {
		Data []byte
	}{
		Data: data,
	}
	mock.lockDecode.Lock()
	mock.calls.Decode = append(mock.calls.Decode, callInfo)
	mock.lockDecode.Unlock()
	return mock.DecodeFunc(data)
}

// DecodeCalls gets all the calls that were made to Decode.
// Check the length with:
//
//	len(mockedICertificateHandler.DecodeCalls())
func (mock *ICertificateHandlerMock) DecodeCalls() []struct {
	Data []byte
} {
	var calls []struct {
		Data []byte
	}
	mock.lockDecode.RLock()
	calls = mock.calls.Decode
	mock.lockDecode.RUnlock()
	return calls
}

// Parse calls ParseFunc.
func (mock *ICertificateHandlerMock) Parse(der []byte) (*x509.Certificate, error) {
	if mock.ParseFunc == nil {
		panic("ICertificateHandlerMock.ParseFunc: method is nil but ICertificateHandler.Parse was just called")
	}
	callInfo := struct {
		Der []byte
	}{
		Der: der,
	}
	mock.lockParse.Lock()
	mock.calls.Parse = append(mock.calls.Parse, callInfo)
	mock.lockParse.Unlock()
	return mock.ParseFunc(der)
}

// ParseCalls gets all the calls that were made to Parse.
// Check the length with:
//
//	len(mockedICertificateHandler.ParseCalls())
func (mock *ICertificateHandlerMock) ParseCalls() []struct {
	Der []byte
} {
	var calls []struct {
		Der []byte
	}
	mock.lockParse.RLock()
	calls = mock.calls.Parse
	mock.lockParse.RUnlock()
	return calls
}