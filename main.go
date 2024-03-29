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
	"slices"
	"strings"

	"cirello.io/moq/internal/moq"
	"cirello.io/moq/internal/typealias"
)

func main() {
	typealias.ConfigureGoDebug()
	log.SetPrefix("moq: ")
	log.SetFlags(0)
	debugFlags := strings.Split(",", strings.ToLower(os.Getenv("MOQ_DEBUG")))

	flagset := flag.NewFlagSet("moq", flag.ExitOnError)
	var moqCfg moq.Config
	outfile := flagset.String("out", "", "output file (default stdout)")
	flagset.StringVar(&moqCfg.PkgName, "pkg", "", "package name (default will infer)")
	if slices.Contains(debugFlags, "disableFormat") {
		moqCfg.Formatter = "disabled"
	}
	flagset.BoolVar(&moqCfg.StubImpl, "stub", false, "return zero values when no mock implementation is provided, do not panic")
	flagset.BoolVar(&moqCfg.SkipEnsure, "skip-ensure", false, "suppress check that confirms a mock implements an interface, avoid import cycle if mocks generated outside of the tested package")
	remove := flagset.Bool("rm", false, "first remove output file, if it exists")
	flagset.BoolVar(&moqCfg.WithResets, "with-resets", false, "generate functions to facilitate resetting calls made to a mock")
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

	moqCfg.SrcDir = flagset.Arg(0)
	namePairs := flagset.Args()[1:]

	buf := new(bytes.Buffer)
	m, err := moq.New(moqCfg)
	if err != nil {
		log.Fatalf("cannot begin mock generation: %v", err)
	}
	if err := m.Mock(buf, namePairs...); err != nil {
		log.Fatalf("cannot render mock: %v", err)
	}
	if *outfile == "" {
		io.Copy(os.Stdout, buf)
		return
	}
	if *remove {
		if err := os.Remove(*outfile); err != nil && !errors.Is(err, os.ErrNotExist) {
			log.Fatalf("cannot remove %q: %v", *outfile, err)
		}
	}
	if err := os.MkdirAll(filepath.Dir(*outfile), 0o750); err != nil {
		log.Fatalf("cannot create base directory: %v", err)
	}
	if err := os.WriteFile(*outfile, buf.Bytes(), 0o600); err != nil {
		log.Fatalf("cannot store generated mock %q: %v", *outfile, err)
	}
}
