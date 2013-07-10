package gofractions

import (
	"fmt"
	"math/big"
	"testing"
)

// TODO: Convert all functions from feature tests to actual tests

func Test1(t *testing.T) {
	f1, _ := NewFraction(20, 2)
	f2 := NewFractionFromInt(20)
	//f3 := NewFraction(20, 0)
	f3, _ := NewFraction(20, 2)
	f3.reduce()
	if f1.String() != f3.String() {
		t.Errorf("Should be equal:", f1.String(), f3.String())
	}
	if f1.String() == f2.String() {
		t.Errorf("Should be different:", f1.String(), f2.String())
	}
}

func Test2(t *testing.T) {
	f, _ := NewFraction(22, 2)
	fmt.Println(f)
	f, _ = NewFraction(33, 3)
	fmt.Println(f)
}

func Test3(t *testing.T) {
	var f *Fraction
	f, _ = NewFraction(16, -10)
	fmt.Println(f)
	f = NewFractionFromInt(123)
	fmt.Println(f)
	f = NewZeroFraction()
	fmt.Println(f)
	f, _ = NewFractionFromString("3/7")
	fmt.Println(f)
	f, _ = NewFractionFromString("6/-14")
	fmt.Println(f)
	f, _ = NewFractionFromString("-3/7")
	fmt.Println(f)
}

func Test4(t *testing.T) {
	x, _ := NewFractionFromString("1/3")
	y, _ := NewFractionFromString("2/4")
	x.Multiply(y)
	fmt.Println(x.String(), "looks nicer than", (1.0/3.0)*(2.0/4.0))
	y.MultiplyInt(2)
	fmt.Println("y is", y.String())
	z := x
	z.Multiply(y)
	fmt.Println("z is", z.String(), z.Round(), "(", (1.0/3.0)*(2.0/4.0)*2*(2.0/4.0), ")")
}

func Test5(t *testing.T) {
	x := NewFractionFromInt(3)
	y := NewFractionFromInt(2)
	x.Divide(y)
	fmt.Println(x.String(), x.Round())
	x.DivideInt(2)
	fmt.Println(x.String(), x.Round())
}

func Test6(t *testing.T) {
	var pi float64 = 3.14159265359
	fmt.Println("num dom i", "\t\t", "fraction", "\t", "float", "\t\t", "rounded")
	f := NewFractionFromFloat64(0.5, -1)
	exact := f.ExactFloat64()
	fmt.Println(f, "\t\t", f.String(), "\t\t", f.Float64(), "\t\texact:", exact, "\t\trounded:", f.Round())
	f = NewFractionFromFloat64(pi, 10000)
	exact = f.ExactFloat64()
	fmt.Println(f, "\t", f.String(), "\t", f.Float64(), "\texact:", exact, "\t\trounded:", f.Round())
}

func Test7(t *testing.T) {
	x := NewFractionFromFloat64(0.7, -1)
	y := NewFractionFromFloat64(0.5, -1)
	x.AddInt(2)
	fmt.Println("0.7 + 2 =", x.String(), x.Round(), x.Float64(), 0.7+2)
	y.SubInt(4)
	fmt.Println("0.5 - 4 =", y.String(), y.Round(), y.Float64(), 0.5-4)
}

func Test8(t *testing.T) {
	x, _ := NewFraction(1, 3)
	y, _ := NewFraction(1, 2)
	fmt.Println(" ", x.String(), x.Float64())
	fmt.Println("+", y.String(), y.Float64())
	x.Add(y)
	fmt.Println("=", x.String(), x.Float64())
}

func Test9(t *testing.T) {
	x, _ := NewFraction(1, 2)
	y, _ := NewFraction(1, 3)
	fmt.Println(" ", x.String(), x.Float64())
	fmt.Println("-", y.String(), y.Float64())
	x.Sub(y)
	fmt.Println("=", x.String(), x.Float64())
}

func Test10(t *testing.T) {
	x, _ := NewFraction(3, 2)
	i, f := x.Splitup()
	fmt.Printf("%s is also %d+%s\n", x, i, f)
}

func Test11(t *testing.T) {
	r := big.NewRat(3, 9)
	f := NewFractionFromRat(r)
	fmt.Println("One third:", r, f)
	rfloat, exact := r.Float64()
	ffloat := f.Float64()
	fmt.Println("One third:", rfloat, "(exact:", exact, ")", ffloat)
	f.SetMaxReduceIterations(4)
	fmt.Println(f.Float64())
	f = NewFractionFromRat(r)
	fmt.Println(f.ExactFloat64())

	r = big.NewRat(2, 4)
	f = NewFractionFromRat(r)
	//f.SetMaxReduceIterations(-1)
	rfloat, exact = r.Float64()
	if exact != f.ExactFloat64() {
		t.Errorf("Both should be exact: %s %s (%v %v)", r, f, exact, f.ExactFloat64())
	}
}
