package main

import (
	"fmt"
	"strings"
	"strconv"
)

const defaultMaxReduceIterations = 400

type Fraction struct {
	top int // numerator
	bot int // denominator
	maxReduceIterations int // maximum number of iterations used to reduce the fraction
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
	counter := 0
	for trydiv := min(abs(my.top), abs(my.bot)); trydiv >= 2; trydiv-- {
		if (my.top/trydiv)*trydiv == my.top && (my.bot/trydiv)*trydiv == my.bot {
			my.top /= trydiv
			my.bot /= trydiv
		}
		if counter == my.maxReduceIterations {
			break
		}
		counter++
	}
	my.prettyNegative()
}

func (my *Fraction) Float64() float64 {
	return float64(my.top) / float64(my.bot)
}

func (my *Fraction) Int() int {
	return int(my.Float64())
}

func (my *Fraction) Round() int {
	return int(my.Float64() + 0.5)
}

func (my *Fraction) String() string {
	if my.top == my.bot {
		return "1"
	}
	return strconv.Itoa(my.top) + "/" + strconv.Itoa(my.bot)
}

func (my *Fraction) prettyNegative() {
	if (my.top < 0 && my.bot < 0) || (my.top > 0 && my.bot < 0) {
		my.top = -my.top
		my.bot = -my.bot
	}
}

func (my *Fraction) Multiply(x Fraction) {
	my.top *= x.top
	my.bot *= x.bot
	my.reduce()
}

func (my *Fraction) Divide(x Fraction) {
	my.top *= x.bot
	my.bot *= x.top
	my.reduce()
}

func (my *Fraction) Add(x Fraction) {
	my.top = my.top*x.bot + x.top*my.bot
	my.bot = my.bot * x.bot
	my.reduce()
}

func (my *Fraction) Sub(x Fraction) {
	my.top = my.top*x.bot - x.top*my.bot
	my.bot = my.bot * x.bot
	my.reduce()
}

func (my *Fraction) MultiplyInt(x int) {
	my.top *= x
	my.reduce()
}

func (my *Fraction) DivideInt(x int) {
	my.bot *= x
	my.reduce()
}

func (my *Fraction) AddInt(x int) {
	my.top += my.bot * x
	my.reduce()
}

func (my *Fraction) SubInt(x int) {
	my.top -= my.bot * x
	my.reduce()
}

// takes a numinator, denumintator and how many iterations should be used (max)
// to reduce the fraction, during calculations
func NewFraction(num int, dom int) Fraction {
	var frac Fraction
	frac.top = num
	if dom == 0 {
		panic(fmt.Sprintf("Can't divide %v on 0", num))
	}
	frac.bot = dom
	frac.maxReduceIterations = defaultMaxReduceIterations
	frac.reduce()
	return frac
}

func (my *Fraction) SetMaxReduceIterations(maxReduceIterations int) {
	my.maxReduceIterations = maxReduceIterations
}

// Splits up a fraction into an integer part, and the rest as another fraction
func (my *Fraction) Splitup() (int, Fraction) {
	i := my.Int()
	clone := *my
	clone.SubInt(i)
	return i, clone
}

// Tries to convert a float to a fraction
// Takes a float and a maximum number of iterations to find the fraction (can be -1)
func NewFractionFromFloat64(f float64, maxIterations int) Fraction {
	// stackoverflow.com/questions/95727/how-to-convert-floats-to-human-readable-fractions
	num := 1
	dom := 1
	result := float64(num) / float64(dom)
	counter := 0
	for result != f {
		if result < f {
			num++
		} else {
			dom++
			num = int(f * float64(dom))
		}
		result = float64(num) / float64(dom)
		if counter == maxIterations {
			break
		}
		counter++
	}
	return NewFraction(num, dom)
}

func NewFractionFromInt(num int) Fraction {
	return NewFraction(num, 1)
}

func NewFractionFromVoid() Fraction {
	return NewFraction(0, 1)
}

func NewFractionFromString(exp string) Fraction {
	top := 0
	bot := 1
	parts := strings.Split(exp, "/", -1)
	if len(parts) == 2 {
		if value, err := strconv.Atoi(parts[0]); err == nil {
			top = value
		} else {
			panic(fmt.Sprintf("Invalid first part of the fraction: %s", parts[0]))
		}
		if value, err := strconv.Atoi(parts[1]); err == nil {
			bot = value
		} else {
			panic(fmt.Sprintf("Invalid second part of the fraction: %s", parts[1]))
		}
	} else {
		panic(fmt.Sprintf("This does not look like a fraction: %s", exp))
	}
	return NewFraction(top, bot)
}

func test1() {
	f1 := NewFraction(20, 2)
	f2 := NewFractionFromInt(20)
	//f3 := NewFraction(20, 0)
	f3 := NewFraction(20, 2)
	f3.reduce()
	fmt.Println(f1.String(), f2.String(), f3.String())
}

func test2() {
	f := NewFraction(22, 2)
	fmt.Println(f.String())
	f = NewFraction(33, 3)
	fmt.Println(f.String())
}

func test3() {
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

func test4() {
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

func test5() {
	var x, y Fraction
	x = NewFractionFromInt(3)
	y = NewFractionFromInt(2)
	x.Divide(y)
	fmt.Println(x.String(), x.Round())
	x.DivideInt(2)
	fmt.Println(x.String(), x.Round())
}

func test6() {
	var pi float64 = 3.14159265359
	fmt.Println("num dom i", "\t\t", "fraction", "\t", "float", "\t\t", "rounded")
	f := NewFractionFromFloat64(0.5, -1)
	fmt.Println(f, "\t\t", f.String(), "\t\t", f.Float64(), "\t\t", f.Round())
	f = NewFractionFromFloat64(pi, 10000)
	fmt.Println(f, "\t", f.String(), "\t", f.Float64(), "\t", f.Round())
}

func test7() {
	x := NewFractionFromFloat64(0.7, -1)
	y := NewFractionFromFloat64(0.5, -1)
	x.AddInt(2)
	fmt.Println("0.7 + 2 =", x.String(), x.Round(), x.Float64(), 0.7+2)
	y.SubInt(4)
	fmt.Println("0.5 - 4 =", y.String(), y.Round(), y.Float64(), 0.5-4)
}

func test8() {
	x := NewFraction(1, 3)
	y := NewFraction(1, 2)
	fmt.Println(" ", x.String(), x.Float64())
	fmt.Println("+", y.String(), y.Float64())
	x.Add(y)
	fmt.Println("=", x.String(), x.Float64())
}

func test9() {
	x := NewFraction(1, 2)
	y := NewFraction(1, 3)
	fmt.Println(" ", x.String(), x.Float64())
	fmt.Println("-", y.String(), y.Float64())
	x.Sub(y)
	fmt.Println("=", x.String(), x.Float64())
}

func test10() {
	x := NewFraction(3, 2)
	i, f := x.Splitup()
	fmt.Println(x.String(), "is also", i, "and", f.String())
}

func main() {
	test1()
	fmt.Println("---")
	test2()
	fmt.Println("---")
	test3()
	fmt.Println("---")
	test4()
	fmt.Println("---")
	test5()
	fmt.Println("---")
	test6()
	fmt.Println("---")
	test7()
	fmt.Println("---")
	test8()
	fmt.Println("---")
	test9()
	fmt.Println("---")
	test10()
}
