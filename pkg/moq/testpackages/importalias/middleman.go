package importalias

import (
	srcclient "cirello.io/moq/pkg/moq/testpackages/importalias/source/client"
	tgtclient "cirello.io/moq/pkg/moq/testpackages/importalias/target/client"
)

type MiddleMan interface {
	Connect(src srcclient.Client, tgt tgtclient.Client)
}
