package num

import (
	"fmt"
	"math/big"
	"testing"
)

func Test1(t *testing.T) {
	f1, _ := New(20, 2)
	f2 := NewFromInt(20)
	f3, _ := New(20, 2)
	f3.reduce()
	if f1.String() != f3.String() {
		t.Errorf("Should be equal: %s %s\n", f1.String(), f3.String())
	}
	if f1.String() == f2.String() {
		t.Errorf("Should be different: %s %s \n", f1.String(), f2.String())
	}
}

func Test2(t *testing.T) {
	f, _ := New(22, 2)
	fmt.Println(f)
	f, _ = New(33, 3)
	fmt.Println(f)
}

func Test3(t *testing.T) {
	var f *Frac
	f, _ = New(16, -10)
	fmt.Println(f)
	f = NewFromInt(123)
	fmt.Println(f)

	fmt.Println(f)
	f, _ = NewFromString("3/7")
	fmt.Println(f)
	f, _ = NewFromString("6/-14")
	fmt.Println(f)
	f, _ = NewFromString("-3/7")
	fmt.Println(f)
}

func Test4(t *testing.T) {
	x, _ := NewFromString("1/3")
	y, _ := NewFromString("2/4")
	x.Mul(y)
	fmt.Println(x.String(), "looks nicer than", (1.0/3.0)*(2.0/4.0))
	y.MulInt(2)
	fmt.Println("y is", y.String())
	z := x
	z.Mul(y)
	fmt.Println("z is", z.String(), z.Round(), "(", (1.0/3.0)*(2.0/4.0)*2*(2.0/4.0), ")")
}

func Test5(t *testing.T) {
	x := NewFromInt(3)
	y := NewFromInt(2)
	x.Div(y)
	fmt.Println(x.String(), x.Round())
	x.DivInt(2)
	fmt.Println(x.String(), x.Round())
}

func Test6(t *testing.T) {
	var pi = 3.14159265359
	fmt.Println("num dom i", "\t\t", "fraction", "\t", "float", "\t\t", "rounded")
	f := NewFromFloat64(0.5, -1)
	exact := f.ExactFloat64()
	fmt.Println(f, "\t\t", f.String(), "\t\t", f.Float64(), "\t\texact:", exact, "\t\trounded:", f.Round())
	f = NewFromFloat64(pi, 10000)
	exact = f.ExactFloat64()
	fmt.Println(f, "\t", f.String(), "\t", f.Float64(), "\texact:", exact, "\t\trounded:", f.Round())
}

func Test7(t *testing.T) {
	x := NewFromFloat64(0.7, -1)
	y := NewFromFloat64(0.5, -1)
	x.AddInt(2)
	fmt.Println("0.7 + 2 =", x.String(), x.Round(), x.Float64(), 0.7+2)
	y.SubInt(4)
	fmt.Println("0.5 - 4 =", y.String(), y.Round(), y.Float64(), 0.5-4)
}

func Test8(t *testing.T) {
	x, _ := New(1, 3)
	y, _ := New(1, 2)
	fmt.Println(" ", x.String(), x.Float64())
	fmt.Println("+", y.String(), y.Float64())
	x.Add(y)
	fmt.Println("=", x.String(), x.Float64())
}

func Test9(t *testing.T) {
	x, _ := New(1, 2)
	y, _ := New(1, 3)
	fmt.Println(" ", x.String(), x.Float64())
	fmt.Println("-", y.String(), y.Float64())
	x.Sub(y)
	fmt.Println("=", x.String(), x.Float64())
}

func Test10(t *testing.T) {
	x, _ := New(3, 2)
	i, f := x.Splitup()
	fmt.Printf("%s is also %d+%s\n", x, i, f)
}

func Test11(t *testing.T) {
	r := big.NewRat(3, 9)
	f := NewFromRat(r)
	fmt.Println("One third:", r, f)
	rfloat, exact := r.Float64()
	ffloat := f.Float64()
	fmt.Println("One third:", rfloat, "(exact:", exact, ")", ffloat)
	f.SetMaxReduceIterations(4)
	fmt.Println(f.Float64())
	f = NewFromRat(r)
	fmt.Println(f.ExactFloat64())

	r = big.NewRat(2, 4)
	f = NewFromRat(r)
	//f.SetMaxReduceIterations(-1)
	_, exact = r.Float64()
	if exact != f.ExactFloat64() {
		t.Errorf("Both should be exact: %s %s (%v %v)", r, f, exact, f.ExactFloat64())
	}
}
