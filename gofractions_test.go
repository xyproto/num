package gofractions

import (
	"testing"
	"fmt"
)

// TODO: Convert all functions from feature tests to actual tests

func Test1(t *testing.T) {
	f1 := NewFraction(20, 2)
	f2 := NewFractionFromInt(20)
	//f3 := NewFraction(20, 0)
	f3 := NewFraction(20, 2)
	f3.reduce()
	if f1.String() != f3.String() {
		t.Errorf("Should be equal:", f1.String(), f3.String())
	}
	if f1.String() == f2.String() {
		t.Errorf("Should be different:", f1.String(), f2.String())
	}
}

func Test2(t *testing.T) {
	f := NewFraction(22, 2)
	fmt.Println(f.String())
	f = NewFraction(33, 3)
	fmt.Println(f.String())
}

func Test3(t *testing.T) {
	var f Fraction
	f = NewFraction(16, -10)
	fmt.Println(f.String())
	f = NewFractionFromInt(123)
	fmt.Println(f.String())
	f = NewFractionFromVoid()
	fmt.Println(f.String())
	f = NewFractionFromString("3/7")
	fmt.Println(f.String())
	f = NewFractionFromString("6/-14")
	fmt.Println(f.String())
	f = NewFractionFromString("-3/7")
	fmt.Println(f.String())
}

func Test4(t *testing.T) {
	var x, y, z Fraction
	x = NewFractionFromString("1/3")
	y = NewFractionFromString("2/4")
	x.Multiply(y)
	fmt.Println(x.String(), "looks nicer than", (1.0/3.0)*(2.0/4.0))
	y.MultiplyInt(2)
	fmt.Println("y is", y.String())
	z = x
	z.Multiply(y)
	fmt.Println("z is", z.String(), z.Round(), "(", (1.0/3.0)*(2.0/4.0)*2*(2.0/4.0), ")")
}

func Test5(t *testing.T) {
	var x, y Fraction
	x = NewFractionFromInt(3)
	y = NewFractionFromInt(2)
	x.Divide(y)
	fmt.Println(x.String(), x.Round())
	x.DivideInt(2)
	fmt.Println(x.String(), x.Round())
}

func Test6(t *testing.T) {
	var pi float64 = 3.14159265359
	fmt.Println("num dom i", "\t\t", "fraction", "\t", "float", "\t\t", "rounded")
	f := NewFractionFromFloat64(0.5, -1)
	fmt.Println(f, "\t\t", f.String(), "\t\t", f.Float64(), "\t\t", f.Round())
	f = NewFractionFromFloat64(pi, 10000)
	fmt.Println(f, "\t", f.String(), "\t", f.Float64(), "\t", f.Round())
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
	x := NewFraction(1, 3)
	y := NewFraction(1, 2)
	fmt.Println(" ", x.String(), x.Float64())
	fmt.Println("+", y.String(), y.Float64())
	x.Add(y)
	fmt.Println("=", x.String(), x.Float64())
}

func Test9(t *testing.T) {
	x := NewFraction(1, 2)
	y := NewFraction(1, 3)
	fmt.Println(" ", x.String(), x.Float64())
	fmt.Println("-", y.String(), y.Float64())
	x.Sub(y)
	fmt.Println("=", x.String(), x.Float64())
}

func Test10(t *testing.T) {
	x := NewFraction(3, 2)
	i, f := x.Splitup()
	fmt.Println(x.String(), "is also", i, "and", f.String())
}
