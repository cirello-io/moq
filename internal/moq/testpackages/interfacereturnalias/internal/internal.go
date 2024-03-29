package internal

import "io"

type A struct{}

type Writer = io.Writer

type Doer interface {
	Do() Writer
}
