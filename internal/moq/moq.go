package moq

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"go/token"
	"go/types"
	"io"
	"strings"

	"cirello.io/moq/internal/registry"
	"cirello.io/moq/internal/template"
)

// Mocker can generate mock structs.
type Mocker struct {
	cfg Config

	registry *registry.Registry
	tmpl     template.Template
}

// Config specifies details about how interfaces should be mocked.
// SrcDir is the only field which needs be specified.
type Config struct {
	SrcDir     string
	PkgName    string
	Formatter  string
	StubImpl   bool
	SkipEnsure bool
	WithResets bool
}

// New makes a new Mocker for the specified package directory.
func New(cfg Config) (*Mocker, error) {
	reg, err := registry.New(cfg.SrcDir, cfg.PkgName)
	if err != nil {
		return nil, err
	}
	return &Mocker{
		cfg:      cfg,
		registry: reg,
		tmpl:     template.New(),
	}, nil
}

// Mock generates a mock for the specified interface name.
func (m *Mocker) Mock(w io.Writer, namePairs ...string) error {
	if len(namePairs) == 0 {
		return errors.New("must specify one interface")
	}

	mocks := make([]template.MockData, len(namePairs))
	for i, np := range namePairs {
		interfaceName, mockName := parseInterfaceName(np)
		interfaceType, typeParams, err := m.registry.LookupInterface(interfaceName)
		if err != nil {
			return err
		}
		methods := make([]template.MethodData, interfaceType.NumMethods())
		for j := 0; j < interfaceType.NumMethods(); j++ {
			methods[j] = m.methodData(interfaceType.Method(j))
		}
		mocks[i] = template.MockData{
			InterfaceName: interfaceName,
			MockName:      mockName,
			Methods:       methods,
			TypeParams:    m.typeParams(typeParams),
		}
	}

	data := template.Data{
		PkgName:    m.mockPkgName(),
		Mocks:      mocks,
		StubImpl:   m.cfg.StubImpl,
		SkipEnsure: m.cfg.SkipEnsure,
		WithResets: m.cfg.WithResets,
	}

	if data.MocksHaveMethod() {
		m.registry.AddImport(types.NewPackage("sync", "sync"))
	}
	if m.registry.SrcPkgName() != m.mockPkgName() {
		data.SrcPkgQualifier = m.registry.SrcPkgName() + "."
		if !m.cfg.SkipEnsure {
			imprt := m.registry.AddImport(m.registry.SrcPkg())
			data.SrcPkgQualifier = imprt.Qualifier() + "."
		}
	}

	data.Imports = m.registry.Imports()

	var buf bytes.Buffer
	if err := m.tmpl.Execute(&buf, data); err != nil {
		return fmt.Errorf("cannot render mock template after code analysis: %w", err)
	}

	formatted, err := m.formatSource(buf.Bytes())
	if err != nil {
		return fmt.Errorf("cannot run pretty printer on generated code: %w", err)
	}

	if _, err := w.Write(formatted); err != nil {
		return err
	}
	return nil
}

func (m *Mocker) typeParams(typeParamsList *types.TypeParamList) []template.TypeParamData {
	var tpd []template.TypeParamData
	if typeParamsList == nil {
		return tpd
	}

	tpd = make([]template.TypeParamData, typeParamsList.Len())
	scope := m.registry.MethodScope()
	for i := 0; i < len(tpd); i++ {
		tp := typeParamsList.At(i)
		typeParam := types.NewParam(token.Pos(i), tp.Obj().Pkg(), tp.Obj().Name(), tp.Constraint())
		tpd[i] = template.TypeParamData{
			ParamData:  template.ParamData{Var: scope.AddVar(typeParam, "")},
			Constraint: explicitConstraintType(typeParam),
		}
	}

	return tpd
}

func explicitConstraintType(typeParam *types.Var) (t types.Type) {
	underlying := typeParam.Type().Underlying().(*types.Interface)
	// check if any of the embedded types is either a basic type or a union,
	// because the generic type has to be an alias for one of those types then
	for j := 0; j < underlying.NumEmbeddeds(); j++ {
		t := underlying.EmbeddedType(j)
		switch t := t.(type) {
		case *types.Basic:
			return t
		case *types.Union: // only unions of basic types are allowed, so just take the first one as a valid type constraint
			return t.Term(0).Type()
		}
	}
	return nil
}

func (m *Mocker) methodData(f *types.Func) template.MethodData {
	sig := f.Type().(*types.Signature)
	scope := m.registry.MethodScope()
	paramsLen := sig.Params().Len()
	resultsLen := sig.Results().Len()
	params := make([]template.ParamData, paramsLen)
	for i := 0; i < paramsLen; i++ {
		p := template.ParamData{
			Var: scope.AddVar(sig.Params().At(i), ""),
		}
		p.Variadic = sig.Variadic() && i == paramsLen-1 && p.Var.IsSlice() // check for final variadic argument
		params[i] = p
	}
	results := make([]template.ParamData, resultsLen)
	for i := 0; i < resultsLen; i++ {
		results[i] = template.ParamData{
			Var: scope.AddVar(sig.Results().At(i), "Out"),
		}
	}
	return template.MethodData{
		Name:    f.Name(),
		Params:  params,
		Returns: results,
	}
}

func (m *Mocker) mockPkgName() string {
	if m.cfg.PkgName != "" {
		return m.cfg.PkgName
	}
	return m.registry.SrcPkgName()
}

func (m *Mocker) formatSource(src []byte) ([]byte, error) {
	if m.cfg.Formatter == "disabled" {
		return src, nil
	}
	formatted, err := format.Source(src)
	if err != nil {
		return nil, fmt.Errorf("go/format: %w", err)
	}
	return formatted, nil
}

func parseInterfaceName(namePair string) (interfaceName, mockName string) {
	interfaceName, mockName, found := strings.Cut(namePair, ":")
	if !found {
		return interfaceName, interfaceName + "Mock"
	}
	return interfaceName, mockName
}
