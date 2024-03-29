package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime/debug"

	"cirello.io/moq/internal/moq"
)

func init() {
	// Necessary hack to get `go/type` to report type aliases. Necessary to
	// get generics support to work correctly when paired with type aliases.
	godebug := os.Getenv("GODEBUG")
	if godebug == "" {
		godebug += ","
	}
	godebug += "gotypesalias=1"
	_ = os.Setenv("GODEBUG", godebug)
}

type userFlags struct {
	moq.Config

	outFile string

	remove    bool
	namePairs []string
}

func main() {
	log.SetPrefix("moq: ")
	log.SetFlags(0)

	flagset := flag.NewFlagSet("moq", flag.ExitOnError)
	var flags userFlags
	flagset.StringVar(&flags.outFile, "out", "", "output file (default stdout)")
	flagset.StringVar(&flags.Config.PkgName, "pkg", "", "package name (default will infer)")
	flagset.StringVar(&flags.Config.Formatter, "fmt", "", "go pretty-printer: gofmt, goimports or noop (default gofmt)")
	flagset.BoolVar(&flags.Config.StubImpl, "stub", false, "return zero values when no mock implementation is provided, do not panic")
	flagset.BoolVar(&flags.Config.SkipEnsure, "skip-ensure", false, "suppress mock implementation check, avoid import cycle if mocks generated outside of the tested package")
	flagset.BoolVar(&flags.remove, "rm", false, "first remove output file, if it exists")
	flagset.BoolVar(&flags.Config.WithResets, "with-resets", false, "generate functions to facilitate resetting calls made to a mock")
	printVersion := flagset.Bool("version", false, "show the version for moq")

	flagset.Usage = func() {
		fmt.Fprintln(flagset.Output(), `moq [flags] source-dir interface [interface2 [interface3 [...]]]`)
		flagset.PrintDefaults()
		fmt.Fprintln(flagset.Output(), `Specifying an alias for the mock is also supported with the format 'interface:alias'`)
		fmt.Fprintln(flagset.Output(), `Ex: moq -pkg different . MyInterface:MyMock`)
	}

	if err := flagset.Parse(os.Args[1:]); err != nil {
		log.Fatal("cannot parse flags")
	}

	if *printVersion {
		info, ok := debug.ReadBuildInfo()
		if !ok {
			log.Fatal("could not read build info")
		}
		version := "dev"
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				version = setting.Value
			}
		}
		fmt.Printf("moq version %s\n", version)
		os.Exit(0)
	}

	if len(flagset.Args()) < 2 {
		log.Fatal("not enough arguments")
	}

	flags.Config.SrcDir = flagset.Arg(0)
	flags.namePairs = flagset.Args()[1:]

	buf := new(bytes.Buffer)
	m, err := moq.New(flags.Config)
	if err != nil {
		log.Fatalf("cannot begin mock generation: %v", err)
	}
	if err := m.Mock(buf, flags.namePairs...); err != nil {
		log.Fatalf("cannot render mock: %v", err)
	}
	if flags.outFile == "" {
		io.Copy(os.Stdout, buf)
		return
	}
	if flags.remove {
		if err := os.Remove(flags.outFile); err != nil && !errors.Is(err, os.ErrNotExist) {
			log.Fatalf("cannot remove %q: %v", flags.outFile, err)
		}
	}
	if err := os.MkdirAll(filepath.Dir(flags.outFile), 0o750); err != nil {
		log.Fatalf("cannot create base directory: %v", err)
	}
	if err := os.WriteFile(flags.outFile, buf.Bytes(), 0o600); err != nil {
		log.Fatalf("cannot store generated mock %q: %v", flags.outFile, err)
	}
}
