package two

import (
	"cirello.io/moq/pkg/moq/testpackages/imports/one"
)

// DoSomething does something.
type DoSomething interface {
	Do(thing one.Thing) error
	Another(thing one.Thing) error
}
