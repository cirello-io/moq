// Code generated by moq; DO NOT EDIT.
// cirello.io/moq

package transientimport

import (
	fourappv1 "cirello.io/moq/pkg/moq/testpackages/transientimport/four/app/v1"
	transientimportonev1 "cirello.io/moq/pkg/moq/testpackages/transientimport/one/v1"
	testpackagestransientimportonev1 "cirello.io/moq/pkg/moq/testpackages/transientimport/onev1"
	threev1 "cirello.io/moq/pkg/moq/testpackages/transientimport/three/v1"
	twoappv1 "cirello.io/moq/pkg/moq/testpackages/transientimport/two/app/v1"
	"sync"
)

// Ensure, that TransientMock does implement Transient.
// If this is not the case, regenerate this file with moq.
var _ Transient = &TransientMock{}

// TransientMock is a mock implementation of Transient.
//
//	func TestSomethingThatUsesTransient(t *testing.T) {
//
//		// make and configure a mocked Transient
//		mockedTransient := &TransientMock{
//			DoSomethingFunc: func(zero testpackagestransientimportonev1.Zero, one transientimportonev1.One, two twoappv1.Two, three threev1.Three, four fourappv1.Four)  {
//				panic("mock out the DoSomething method")
//			},
//		}
//
//		// use mockedTransient in code that requires Transient
//		// and then make assertions.
//
//	}
type TransientMock struct {
	// DoSomethingFunc mocks the DoSomething method.
	DoSomethingFunc func(zero testpackagestransientimportonev1.Zero, one transientimportonev1.One, two twoappv1.Two, three threev1.Three, four fourappv1.Four)

	// calls tracks calls to the methods.
	calls struct {
		// DoSomething holds details about calls to the DoSomething method.
		DoSomething []struct {
			// Zero is the zero argument value.
			Zero testpackagestransientimportonev1.Zero
			// One is the one argument value.
			One transientimportonev1.One
			// Two is the two argument value.
			Two twoappv1.Two
			// Three is the three argument value.
			Three threev1.Three
			// Four is the four argument value.
			Four fourappv1.Four
		}
	}
	lockDoSomething sync.RWMutex
}

// DoSomething calls DoSomethingFunc.
func (mock *TransientMock) DoSomething(zero testpackagestransientimportonev1.Zero, one transientimportonev1.One, two twoappv1.Two, three threev1.Three, four fourappv1.Four) {
	if mock.DoSomethingFunc == nil {
		panic("TransientMock.DoSomethingFunc: method is nil but Transient.DoSomething was just called")
	}
	callInfo := struct {
		Zero  testpackagestransientimportonev1.Zero
		One   transientimportonev1.One
		Two   twoappv1.Two
		Three threev1.Three
		Four  fourappv1.Four
	}{
		Zero:  zero,
		One:   one,
		Two:   two,
		Three: three,
		Four:  four,
	}
	mock.lockDoSomething.Lock()
	mock.calls.DoSomething = append(mock.calls.DoSomething, callInfo)
	mock.lockDoSomething.Unlock()
	mock.DoSomethingFunc(zero, one, two, three, four)
}

// DoSomethingCalls gets all the calls that were made to DoSomething.
// Check the length with:
//
//	len(mockedTransient.DoSomethingCalls())
func (mock *TransientMock) DoSomethingCalls() []struct {
	Zero  testpackagestransientimportonev1.Zero
	One   transientimportonev1.One
	Two   twoappv1.Two
	Three threev1.Three
	Four  fourappv1.Four
} {
	var calls []struct {
		Zero  testpackagestransientimportonev1.Zero
		One   transientimportonev1.One
		Two   twoappv1.Two
		Three threev1.Three
		Four  fourappv1.Four
	}
	mock.lockDoSomething.RLock()
	calls = mock.calls.DoSomething
	mock.lockDoSomething.RUnlock()
	return calls
}
