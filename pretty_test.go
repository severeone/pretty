package pretty

import "bytes"
import "testing"

type Bag map[string]interface{}

type Struct struct {
	N int
	S string
	B bool
	A []int
	Z []int
}

var (
	ch chan string

	s = struct {
		n int
		s string
	}{
		42,
		"hello world",
	}

	x = struct{}{}

	arry = []Bag{bag, bag, bag}

	strutty = Struct{N: 42, S: "Hello", B: true, A: []int{1, 2, 3}}

	bag = Bag{
		"a": 1,
		"b": false,
		"c": "some stuff",
		"d": []float64{0.0, 0.1, 1.2, 1.23, 1.23456, 999999999999},
		"e": Bag{
			"e1": "here",
			"e2": []int{1, 2, 3, 4},
			"e3": nil,
		},
		"s":   s,
		"x":   x,
		"z":   []int{},
		"bad": ch,
	}
)

func TestPrettyPrint(test *testing.T) {
	PrettyPrint(arry)
}

func TestPrettyFormat(test *testing.T) {
	test.Log(PrettyFormat(bag))
}

func TestStruct(test *testing.T) {
	test.Log(PrettyFormat(strutty))
}

func TestPretty(test *testing.T) {
	var out bytes.Buffer
	p := Pretty{Indent: "", Out: &out, NilString: "nil"}
	p.Print(strutty)
	test.Log(out.String())
}

func TestPrettyCompact(test *testing.T) {
	var out bytes.Buffer
	p := Pretty{Indent: "", Out: &out, NilString: "nil", Compact: true}
	p.Print(strutty)
	test.Log(out.String())
}

func TestTabPrint(test *testing.T) {
	tp := NewTabPrinter(8)

	for i := 0; i < 33; i++ {
		tp.Print(i)
	}

	tp.Println()

	for _, v := range []string{"one", "two", "three", "four", "five", "six"} {
		tp.Print(v)
	}

	tp.Println()
}
