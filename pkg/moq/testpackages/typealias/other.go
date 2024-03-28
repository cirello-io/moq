package typealias

import (
	"cirello.io/moq/pkg/moq/testpackages/typealias/internal"
)

type AliasToA = internal.A

type Doer interface {
	Do(AliasToA)
}
