package parse

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/vetcher/go-astra/types"
)

func prepareTest() *File {
	path := "./testdata/tester.go"
	f := NewFileParser(path)
	if err := f.Parse(); err != nil {
		panic(err)
	}
	return f
}

func TestFile_GetFunction(t *testing.T) {
	file := prepareTest()
	g := NewGomegaWithT(t)

	fixtures := []struct {
		input  string
		expect types.Function
	}{
		{
			input: "TestFunction",
			expect: types.Function{
				Base: types.Base{
					Name: "TestFunction",
					Docs: nil,
				},
				Args: nil,
				Results: []types.Variable{
					{
						Base: types.Base{},
						Type: types.TName{TypeName: "error"},
					},
				},
			},
		},
		{
			input: "testFunction",
			expect: types.Function{
				Base: types.Base{
					Name: "TestFunction",
					Docs: nil,
				},
				Args: nil,
				Results: []types.Variable{
					{
						Base: types.Base{},
						Type: types.TName{TypeName: "error"},
					},
				},
			},
		},
		{
			input:  "TestStruct",
			expect: types.Function{},
		},
	}

	for _, tt := range fixtures {
		out := file.GetFunction(tt.input)
		g.Expect(out).To(Equal(tt.expect))
	}
}

func TestFile_GetMethod(t *testing.T) {
	file := prepareTest()
	g := NewGomegaWithT(t)

	fixtures := []struct {
		input  string
		expect types.Method
	}{
		{
			input: "TestMethod",
			expect: types.Method{
				Function: types.Function{
					Base: types.Base{
						Name: "TestMethod",
						Docs: nil,
					},
					Args: nil,
				},
				Receiver: types.Variable{
					Base: types.Base{
						Name: "ts",
						Docs: nil,
					},
					Type: types.TPointer{
						NumberOfPointers: 1,
						Next:             types.TName{TypeName: "TestStruct"},
					},
				},
			},
		},
		{
			input: "testmethod",
			expect: types.Method{
				Function: types.Function{
					Base: types.Base{
						Name: "TestMethod",
						Docs: nil,
					},
					Args: nil,
				},
				Receiver: types.Variable{
					Base: types.Base{
						Name: "ts",
						Docs: nil,
					},
					Type: types.TPointer{
						NumberOfPointers: 1,
						Next:             types.TName{TypeName: "TestStruct"},
					},
				},
			},
		},
		{
			input:  "TestStruct",
			expect: types.Method{},
		},
		{
			input:  "",
			expect: types.Method{},
		},
		{
			input:  "TestFunction",
			expect: types.Method{},
		},
	}

	for _, tt := range fixtures {
		out := file.GetMethod(tt.input)
		g.Expect(out).To(Equal(tt.expect))
	}
}


func TestFile_GetInterface(t *testing.T) {
	file := prepareTest()
	g := NewGomegaWithT(t)

	fixtures := []struct {
		input  string
		expect types.Interface
	}{
		{
			input: "TestInterface",
			expect: types.Interface{
			},
		},
		{
			input:  "TestStruct",
			expect: types.Interface{},
		},
	}

	for _, tt := range fixtures {
		out := file.GetInterface(tt.input)
		g.Expect(out).To(Equal(tt.expect))
	}
}


func TestFile_HasFunction(t *testing.T) {
	file := prepareTest()
	g := NewGomegaWithT(t)

	fixtures := []struct {
		input     string
		expectHas bool
	}{
		{
			input:     "TestFunction",
			expectHas: true,
		},
		{
			input:     "testfunction",
			expectHas: true,
		},
		{
			input:     "NotExisting",
			expectHas: false,
		},
		{
			input:     "TestMethod",
			expectHas: false,
		},
	}

	for _, tt := range fixtures {
		out := file.HasFunction(tt.input)
		g.Expect(out).To(Equal(tt.expectHas))
	}
}

func TestFile_HasMethod(t *testing.T) {
	file := prepareTest()
	g := NewGomegaWithT(t)

	fixtures := []struct {
		input     string
		expectHas bool
	}{
		{
			input:     "TestFunction",
			expectHas: false,
		},
		{
			input:     "testfunction",
			expectHas: false,
		},
		{
			input:     "NotExisting",
			expectHas: false,
		},
		{
			input:     "TestMethod",
			expectHas: true,
		},
		{
			input:     "testmethod",
			expectHas: true,
		},
	}

	for _, tt := range fixtures {
		out := file.HasMethod(tt.input)
		g.Expect(out).To(Equal(tt.expectHas))
	}
}

func TestFile_HasInterface(t *testing.T) {
	file := prepareTest()
	g := NewGomegaWithT(t)

	fixtures := []struct {
		input     string
		expectHas bool
	}{
		{
			input:     "TestInterface",
			expectHas: true,
		},
		{
			input:     "testinterface",
			expectHas: true,
		},
		{
			input:     "notExisting",
			expectHas: false,
		},
		{
			input:     "TestStruct",
			expectHas: false,
		},
	}

	for _, tt := range fixtures {
		out := file.HasInterface(tt.input)
		g.Expect(out).To(Equal(tt.expectHas))
	}
}
