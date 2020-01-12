package testdata

func TestFunction() error {
	return nil
}

type TestStruct struct {
}

func (ts *TestStruct) TestMethod() {
}

type TestInterface interface {
	Foo(bar string) error
}