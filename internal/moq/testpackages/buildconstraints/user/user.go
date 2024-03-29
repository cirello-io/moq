package user

import "github.com/cirello-io/moq-test-pkgs/buildconstraints"

// Service does something good with computers.
type Service interface {
	DoSomething(buildconstraints.SomeType) error
}
