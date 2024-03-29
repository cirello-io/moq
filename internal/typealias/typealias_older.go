//go:build !go1.22
// +build !go1.22

package typealias

import (
	"fmt"
	"os"
)

func ConfigureGoDebug() {
	fmt.Fprintln(os.Stderr, "only partial generics and type alias support")
}
