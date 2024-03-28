package syncimport

import (
	stdsync "sync"

	"cirello.io/moq/pkg/moq/testpackages/syncimport/sync"
)

type Syncer interface {
	Blah(s sync.Thing, wg *stdsync.WaitGroup)
}
