package moq

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func copyFile(srcPath, destPath string, item os.FileInfo) error {
	if item.IsDir() {
		if err := os.MkdirAll(destPath, os.FileMode(0750)); err != nil {
			return err
		}
		items, err := os.ReadDir(srcPath)
		if err != nil {
			return err
		}
		for _, item := range items {
			src := filepath.Join(srcPath, item.Name())
			dest := filepath.Join(destPath, item.Name())
			info, err := item.Info()
			if err != nil {
				return err
			}
			if err := copyFile(src, dest, info); err != nil {
				return err
			}
		}
		return nil
	}
	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer src.Close()

	dest, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer dest.Close()

	_, err = io.Copy(dest, src)
	if err != nil {
		return err
	}
	return nil
}

func copyTestPackage(srcPath string) (string, error) {
	tmpDir, err := os.MkdirTemp("", "moq-tests")
	if err != nil {
		return "", err
	}

	info, err := os.Lstat(srcPath)
	if err != nil {
		return tmpDir, err
	}
	return tmpDir, copyFile(srcPath, tmpDir, info)
}

func TestModulesSamePackage(t *testing.T) {
	tmpDir, err := copyTestPackage("testpackages/modules")
	defer os.RemoveAll(tmpDir)
	if err != nil {
		t.Fatalf("Test package copy error: %s", err)
	}

	m, err := New(Config{SrcDir: tmpDir})
	if err != nil {
		t.Fatalf("moq.New: %s", err)
	}
	var buf bytes.Buffer
	err = m.Mock(&buf, "Foo")
	if err != nil {
		t.Errorf("m.Mock: %s", err)
	}
	s := buf.String()
	if strings.Contains(s, `cirello.io/moq/modules`) {
		t.Errorf("should not have cyclic dependency")
	}
	// assertions of things that should be mentioned
	var strs = []string{
		"package simple",
		"var _ Foo = &FooMock{}",
		"type FooMock struct",
	}
	for _, str := range strs {
		if !strings.Contains(s, str) {
			t.Errorf("expected but missing: \"%s\"", str)
		}
	}
}
func TestModulesNestedPackage(t *testing.T) {
	tmpDir, err := copyTestPackage("testpackages/modules")
	defer os.RemoveAll(tmpDir)
	if err != nil {
		t.Fatalf("Test package copy error: %s", err)
	}

	m, err := New(Config{SrcDir: tmpDir, PkgName: "nested"})
	if err != nil {
		t.Fatalf("moq.New: %s", err)
	}
	var buf bytes.Buffer
	err = m.Mock(&buf, "Foo")
	if err != nil {
		t.Errorf("m.Mock: %s", err)
	}
	s := buf.String()
	// assertions of things that should be mentioned
	var strs = []string{
		"package nested",
		"cirello.io/moq/modules",
		"var _ simple.Foo = &FooMock{}",
		"type FooMock struct",
		"func (mock *FooMock) FooIt(bar *simple.Bar) {",
	}
	for _, str := range strs {
		if !strings.Contains(s, str) {
			t.Errorf("expected but missing: \"%s\"", str)
		}
	}
}
