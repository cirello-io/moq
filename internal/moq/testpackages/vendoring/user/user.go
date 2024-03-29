package user

import "github.com/cirello-io/moq-test-pkgs/somerepo"

// Service does something good with computers.
type Service interface {
	DoSomething(somerepo.SomeType) error
}
