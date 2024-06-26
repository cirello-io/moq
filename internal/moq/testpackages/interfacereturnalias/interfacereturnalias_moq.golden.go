// Code generated by moq; DO NOT EDIT.
// cirello.io/moq

package interfacereturnalias

import (
	"cirello.io/moq/pkg/moq/testpackages/interfacereturnalias/internal"
	"sync"
)

// Ensure, that DoerMock does implement internal.Doer.
// If this is not the case, regenerate this file with moq.
var _ internal.Doer = &DoerMock{}

// DoerMock is a mock implementation of internal.Doer.
//
//	func TestSomethingThatUsesDoer(t *testing.T) {
//
//		// make and configure a mocked internal.Doer
//		mockedDoer := &DoerMock{
//			DoFunc: func() internal.Writer {
//				panic("mock out the Do method")
//			},
//		}
//
//		// use mockedDoer in code that requires internal.Doer
//		// and then make assertions.
//
//	}
type DoerMock struct {
	// DoFunc mocks the Do method.
	DoFunc func() internal.Writer

	// calls tracks calls to the methods.
	calls struct {
		// Do holds details about calls to the Do method.
		Do []struct {
		}
	}
	lockDo sync.RWMutex
}

// Do calls DoFunc.
func (mock *DoerMock) Do() internal.Writer {
	if mock.DoFunc == nil {
		panic("DoerMock.DoFunc: method is nil but Doer.Do was just called")
	}
	callInfo := struct {
	}{}
	mock.lockDo.Lock()
	mock.calls.Do = append(mock.calls.Do, callInfo)
	mock.lockDo.Unlock()
	return mock.DoFunc()
}

// DoCalls gets all the calls that were made to Do.
// Check the length with:
//
//	len(mockedDoer.DoCalls())
func (mock *DoerMock) DoCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockDo.RLock()
	calls = mock.calls.Do
	mock.lockDo.RUnlock()
	return calls
}
