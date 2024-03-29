// Code generated by moq; DO NOT EDIT.
// cirello.io/moq

package generate

import (
	"sync"
)

// Ensure, that MyInterfaceStringMock does implement MyInterfaceString.
// If this is not the case, regenerate this file with moq.
var _ MyInterfaceString = &MyInterfaceStringMock{}

// MyInterfaceStringMock is a mock implementation of MyInterfaceString.
//
//	func TestSomethingThatUsesMyInterfaceString(t *testing.T) {
//
//		// make and configure a mocked MyInterfaceString
//		mockedMyInterfaceString := &MyInterfaceStringMock{
//			OneFunc: func(s string) bool {
//				panic("mock out the One method")
//			},
//			ThreeFunc: func() string {
//				panic("mock out the Three method")
//			},
//			TwoFunc: func() int {
//				panic("mock out the Two method")
//			},
//		}
//
//		// use mockedMyInterfaceString in code that requires MyInterfaceString
//		// and then make assertions.
//
//	}
type MyInterfaceStringMock struct {
	// OneFunc mocks the One method.
	OneFunc func(s string) bool

	// ThreeFunc mocks the Three method.
	ThreeFunc func() string

	// TwoFunc mocks the Two method.
	TwoFunc func() int

	// calls tracks calls to the methods.
	calls struct {
		// One holds details about calls to the One method.
		One []struct {
			// S is the s argument value.
			S string
		}
		// Three holds details about calls to the Three method.
		Three []struct {
		}
		// Two holds details about calls to the Two method.
		Two []struct {
		}
	}
	lockOne   sync.RWMutex
	lockThree sync.RWMutex
	lockTwo   sync.RWMutex
}

// One calls OneFunc.
func (mock *MyInterfaceStringMock) One(s string) bool {
	if mock.OneFunc == nil {
		panic("MyInterfaceStringMock.OneFunc: method is nil but MyInterfaceString.One was just called")
	}
	callInfo := struct {
		S string
	}{
		S: s,
	}
	mock.lockOne.Lock()
	mock.calls.One = append(mock.calls.One, callInfo)
	mock.lockOne.Unlock()
	return mock.OneFunc(s)
}

// OneCalls gets all the calls that were made to One.
// Check the length with:
//
//	len(mockedMyInterfaceString.OneCalls())
func (mock *MyInterfaceStringMock) OneCalls() []struct {
	S string
} {
	var calls []struct {
		S string
	}
	mock.lockOne.RLock()
	calls = mock.calls.One
	mock.lockOne.RUnlock()
	return calls
}

// Three calls ThreeFunc.
func (mock *MyInterfaceStringMock) Three() string {
	if mock.ThreeFunc == nil {
		panic("MyInterfaceStringMock.ThreeFunc: method is nil but MyInterfaceString.Three was just called")
	}
	callInfo := struct {
	}{}
	mock.lockThree.Lock()
	mock.calls.Three = append(mock.calls.Three, callInfo)
	mock.lockThree.Unlock()
	return mock.ThreeFunc()
}

// ThreeCalls gets all the calls that were made to Three.
// Check the length with:
//
//	len(mockedMyInterfaceString.ThreeCalls())
func (mock *MyInterfaceStringMock) ThreeCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockThree.RLock()
	calls = mock.calls.Three
	mock.lockThree.RUnlock()
	return calls
}

// Two calls TwoFunc.
func (mock *MyInterfaceStringMock) Two() int {
	if mock.TwoFunc == nil {
		panic("MyInterfaceStringMock.TwoFunc: method is nil but MyInterfaceString.Two was just called")
	}
	callInfo := struct {
	}{}
	mock.lockTwo.Lock()
	mock.calls.Two = append(mock.calls.Two, callInfo)
	mock.lockTwo.Unlock()
	return mock.TwoFunc()
}

// TwoCalls gets all the calls that were made to Two.
// Check the length with:
//
//	len(mockedMyInterfaceString.TwoCalls())
func (mock *MyInterfaceStringMock) TwoCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockTwo.RLock()
	calls = mock.calls.Two
	mock.lockTwo.RUnlock()
	return calls
}
