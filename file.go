package parse

import (
	"strings"

	"github.com/vetcher/go-astra"
	"github.com/vetcher/go-astra/types"
)

type File struct {
	path   string
	parsed *types.File
}

func NewFileParser(path string) *File {
	return &File{
		path:   path,
		parsed: nil,
	}
}

// Parse will call astra.ParseFile and return any occurred error.
// If the parsing was successful, any registered validators are executed.
func (f *File) Parse() (err error) {
	f.parsed, err = astra.ParseFile(f.path)
	if err != nil {
		return err
	}

	return nil
}

// HasFunction checks whether the file has a function with the given name
func (f *File) HasFunction(name string) bool {
	for _, function := range f.parsed.Functions {
		if strings.ToLower(function.Name) == strings.ToLower(name) {
			return true
		}
	}
	return false
}

// HasMethod checks whether the file has a method with the given name
func (f *File) HasMethod(name string) bool {
	for _, method := range f.parsed.Methods {
		if strings.ToLower(method.Name) == strings.ToLower(name) {
			return true
		}
	}
	return false
}

// HasInterface returns true if the interface exists in the parsed file.
func (f *File) HasInterface(ifaceName string) bool {
	for _, i := range f.parsed.Interfaces {
		if strings.ToLower(i.Name) == strings.ToLower(ifaceName) {
			return true
		}
	}
	return false
}

// GetMethod attempts to return the types.Method with the given name.
// If the method was not found, an empty 'types.Method' struct is returned
func (f *File) GetMethod(name string) types.Method {
	for _, method := range f.parsed.Methods {
		if strings.ToLower(method.Name) == strings.ToLower(name) {
			return method
		}
	}
	return types.Method{}
}

// GetInterface attempts the types.Function with the given name.
// If the method was not found, an empty 'types.Function' struct is returned.
func (f *File) GetFunction(name string) types.Function {
	for _, fun := range f.parsed.Functions {
		if strings.ToLower(fun.Name) == strings.ToLower(name) {
			return fun
		}
	}
	return types.Function{}
}

// GetInterface attempts the types.Function with the given name.
// If the method was not found, an empty 'types.Function' struct is returned.
func (f *File) GetInterface(name string) types.Interface {
	for _, iface := range f.parsed.Interfaces {
		if iface.Name == name {
			return iface
		}
	}
	return types.Interface{}
}

