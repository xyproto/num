package main

import (
	"fmt"
	"strings"
	"strconv"
)

type Fraction struct {
	num int // numerator
	dum int // denominator
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (my *Fraction) reduce() {
	var maxtry int = 300
	for trydiv := min(maxtry, min(abs(my.num), abs(my.dum))); trydiv >= 2; trydiv-- {
		for (my.num/trydiv)*trydiv == my.num && (my.dum/trydiv)*trydiv == my.dum {
			my.num /= trydiv
			my.dum /= trydiv
		}
	}
	my.prettyNegative()
}

func Multiply(x, y Fraction) Fraction {
	return NewFraction(x.num * y.num, x.dum * y.dum)
}

func (my *Fraction) Float64() float64 {
	return float64(my.num) / float64(my.dum)
}

func (my *Fraction) Int() int {
	return int(my.Float64())
}

func (my *Fraction) Round() int {
	return int(my.Float64() + 0.5)
}

func (my *Fraction) String() string {
	if my.num == 1 && my.dum == 1 {
		return "1"
	}
	return strconv.Itoa(my.num) + "/" + strconv.Itoa(my.dum)
}

func (my *Fraction) prettyNegative() {
	if (my.num < 0 && my.dum < 0) || (my.num > 0 && my.dum < 0) {
		my.num = -my.num
		my.dum = -my.dum
	}
}

func (my *Fraction) Multiply(x Fraction) {
	my.num *= x.num
	my.dum *= x.dum
	my.reduce()
}

func (my *Fraction) Divide(x Fraction) {
	my.num *= x.dum
	my.dum *= x.num
	my.reduce()
}

func (my *Fraction) MultiplyInt(x int) {
	my.num *= x
	my.reduce()
}

func (my *Fraction) DivideInt(x int) {
	my.dum *= x
	my.reduce()
}

func (my *Fraction) AddInt(x int) {
	my.num += my.dum * x
	my.reduce()
}

func (my *Fraction) SubInt(x int) {
	my.num -= my.dum * x
	my.reduce()
}

func NewFraction(num int, dum int) Fraction {
	var f Fraction
	f.num = num
	if dum == 0 {
		panic(fmt.Sprintf("Can't divide %v on 0", num))
	}
	f.dum = dum
	f.reduce()
	return f
}

// Tries to convert a float to a fraction
// Takes a float and a maximum number of iterations (can be -1)
func NewFractionFromFloat64(f float64, maxIterations int) Fraction {
	// stackoverflow.com/questions/95727/how-to-convert-floats-to-human-readable-fractions
	num := 1
	dum := 1
	result := float64(num) / float64(dum)
	counter := 0
	for (result != f) {
		if result < f {
			num++
		} else {
			dum++
			num = int(f * float64(dum))
		}
		result = float64(num) / float64(dum)
		if counter == maxIterations {
			break
		}
		counter++
	}
	return NewFraction(num, dum)
}

func NewFractionFromInt(num int) Fraction {
	return NewFraction(num, 1)
}

func NewFractionFromVoid() Fraction {
	return NewFraction(0, 1)
}

func NewFractionFromString(exp string) Fraction {
	num := 0
	dum := 1
	parts := strings.Split(exp, "/", -1)
	if len(parts) == 2 {
		if value, err := strconv.Atoi(parts[0]); err == nil {
			num = value
		} else {
			panic(fmt.Sprintf("Invalid first part of the fraction: %s", parts[0]))
		}
		if value, err := strconv.Atoi(parts[1]); err == nil {
			dum = value
		} else {
			panic(fmt.Sprintf("Invalid second part of the fraction: %s", parts[1]))
		}
	} else {
		panic(fmt.Sprintf("This does not look like a fraction: %s", exp))
	}
	return NewFraction(num, dum)
}

func test1() {
	f1 := NewFraction(20, 2)
	f2 := NewFractionFromInt(20)
	//f3 := NewFraction(20, 0)
	f3 := NewFraction(20, 2)
	f3.reduce()
	fmt.Println(f1, f2, f3)
}

func test2() {
	f := NewFraction(22, 2)
	fmt.Println(f)
	f = NewFraction(33, 3)
	fmt.Println(f)
}

func test3() {
	var f Fraction
	f = NewFraction(16, -10)
	fmt.Println(f)
	f = NewFractionFromInt(123)
	fmt.Println(f)
	f = NewFractionFromVoid()
	fmt.Println(f)
	f = NewFractionFromString("3/7")
	fmt.Println(f)
	f = NewFractionFromString("6/-14")
	fmt.Println(f)
	f = NewFractionFromString("-3/7")
	fmt.Println(f)
}

func test5() {
	var x, y, z Fraction
	x = NewFractionFromString("1/3")
	y = NewFractionFromString("2/4")
	x.Multiply(y)
	fmt.Println(x.String(), "looks nicer than", (1.0/3.0) * (2.0/4.0))
	y.MultiplyInt(2)
	fmt.Println("y is", y.String())
	z = Multiply(x, y)
	fmt.Println("z is", z.String(), z.Round(), "(", (1.0/3.0) * (2.0/4.0) * 2 * (2.0/4.0), ")")
}

func test6() {
	var x, y Fraction
	x = NewFractionFromInt(3)
	y = NewFractionFromInt(2)
	x.Divide(y)
	fmt.Println(x.String(), x.Round())
	x.DivideInt(2)
	fmt.Println(x.String(), x.Round())
}

func test7() {
	var pi float64 = 3.14159265359
	f := NewFractionFromFloat64(0.5, -1)
	fmt.Println(f, "\t\t", f.String(), "\t\t", f.Float64(), "\t\t", f.Round())
	f = NewFractionFromFloat64(pi, 10000)
	fmt.Println(f, "\t", f.String(), "\t", f.Float64(), "\t", f.Round())
}

func test8() {
	x := NewFractionFromFloat64(0.7, -1)
	y := NewFractionFromFloat64(0.5, -1)
	x.AddInt(2)
	fmt.Println(x, x.String(), x.Round(), x.Float64(), 0.7+2)
	y.SubInt(4)
	fmt.Println(y, y.String(), y.Round(), y.Float64(), 0.5-4)	
}

func main() {
	test1()
	fmt.Println("---")
	test2()
	fmt.Println("---")
	test3()
	fmt.Println("---")
	test5()
	fmt.Println("---")
	test6()
	fmt.Println("---")	
	test7()
	fmt.Println("---")	
	test8()
}

