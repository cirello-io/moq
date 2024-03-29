//go:build !go1.22
// +build !go1.22

package typealias

import (
	"fmt"
	"os"
)

func ConfigureGoDebug() {
	fmt.Fprintln(os.Stderr, "type alias inside generics support only available in Go 1.22+")
	os.Exit(1)
}
