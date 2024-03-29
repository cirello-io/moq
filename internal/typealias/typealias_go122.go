//go:build go1.22
// +build go1.22

package typealias

import "os"

func ConfigureGoDebug() {
	// Necessary hack to get `go/type` to report type aliases. Necessary to
	// get generics support to work correctly when paired with type aliases.
	godebug := os.Getenv("GODEBUG")
	if godebug == "" {
		godebug += ","
	}
	godebug += "gotypesalias=1"
	_ = os.Setenv("GODEBUG", godebug)
}
