package main

import "fmt"

type Foo struct {
	A int
	B string
}

// method declaration cf function declaration
// a method receiver declared between func keyword and
// the name of method(String)
func (f Foo) String() string {
	// no virtual overriting or virtual method dispatch
	// so f.fieldCount() is not overwrite by Bar.fieldCount()
	return fmt.Sprintf("%d fields: A: %d and B: %s", f.fieldCount(), f.A, f.B)
}

func (f Foo) fieldCount() int {
	return 2
}

// pointer to Foo. called reference receiver
func (f *Foo) Double() {
	f.A = f.A * 2
}

type Bar struct {
	C   bool
	Foo //embedded struct
}

func (b Bar) String() string {
	return fmt.Sprintf("%s and C: %v", b.Foo.String(), b.C)
}

// this will never be called from inside of Foo
// struct embedding is not inheritance
func (b Bar) fieldCount() int {
	return 3
}

func main() {
	f := Foo{
		A: 10,
		B: "Hello",
	}
	b := Bar{
		C:   true,
		Foo: f,
	}
	fmt.Println(b.String())
	b.Double()
	fmt.Println(b.String())
}
