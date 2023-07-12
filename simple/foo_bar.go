package simple

type Foo struct {
}

func NewFoo() *Foo {
	return &Foo{}
}

type Bar struct {
}

func NewBar() *Bar {
	return &Bar{}
}

type FooBar struct {
	*Foo
	*Bar
}

// implementasi struct provider (walapun jarang sekali digunakan, biasanya gunakan func)
