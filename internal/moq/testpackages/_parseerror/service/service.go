package service

import "github.com/cirello-io/notexist"

// Service does something good with computers.
type Service interface {
	DoSomething(notexist.SomeType) error
}
