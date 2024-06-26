// Code generated by moq; DO NOT EDIT.
// cirello.io/moq

package typealias

import (
	"sync"
)

// Ensure, that DoerMock does implement Doer.
// If this is not the case, regenerate this file with moq.
var _ Doer = &DoerMock{}

// DoerMock is a mock implementation of Doer.
//
//	func TestSomethingThatUsesDoer(t *testing.T) {
//
//		// make and configure a mocked Doer
//		mockedDoer := &DoerMock{
//			DoFunc: func(v AliasToA)  {
//				panic("mock out the Do method")
//			},
//		}
//
//		// use mockedDoer in code that requires Doer
//		// and then make assertions.
//
//	}
type DoerMock struct {
	// DoFunc mocks the Do method.
	DoFunc func(v AliasToA)

	// calls tracks calls to the methods.
	calls struct {
		// Do holds details about calls to the Do method.
		Do []struct {
			// V is the v argument value.
			V AliasToA
		}
	}
	lockDo sync.RWMutex
}

// Do calls DoFunc.
func (mock *DoerMock) Do(v AliasToA) {
	if mock.DoFunc == nil {
		panic("DoerMock.DoFunc: method is nil but Doer.Do was just called")
	}
	callInfo := struct {
		V AliasToA
	}{
		V: v,
	}
	mock.lockDo.Lock()
	mock.calls.Do = append(mock.calls.Do, callInfo)
	mock.lockDo.Unlock()
	mock.DoFunc(v)
}

// DoCalls gets all the calls that were made to Do.
// Check the length with:
//
//	len(mockedDoer.DoCalls())
func (mock *DoerMock) DoCalls() []struct {
	V AliasToA
} {
	var calls []struct {
		V AliasToA
	}
	mock.lockDo.RLock()
	calls = mock.calls.Do
	mock.lockDo.RUnlock()
	return calls
}
