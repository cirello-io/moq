package generate

// In a terminal, run `go generate` in this directory to have
// it generates the generated.go file.

//go:generate go run cirello.io/moq -out generated.go . MyInterfaceString

// MyInterface is a test interface.
type MyInterface[T any] interface {
	One(T) bool
	Two() int
	Three() string
}

type MyInterfaceString = MyInterface[string]
