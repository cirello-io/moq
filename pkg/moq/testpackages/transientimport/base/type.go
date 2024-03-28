package base

import (
	four "cirello.io/moq/pkg/moq/testpackages/transientimport/four/app/v1"
	one "cirello.io/moq/pkg/moq/testpackages/transientimport/one/v1"
	"cirello.io/moq/pkg/moq/testpackages/transientimport/onev1"
	three "cirello.io/moq/pkg/moq/testpackages/transientimport/three/v1"
	two "cirello.io/moq/pkg/moq/testpackages/transientimport/two/app/v1"
)

type Transient interface {
	DoSomething(onev1.Zero, one.One, two.Two, three.Three, four.Four)
}
