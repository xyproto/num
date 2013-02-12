package gofractions

import (
	"fmt"
	"strconv"
	"strings"
)

const defaultMaxReduceIterations = 400

type Fraction struct {
	top                 int // numerator
	bot                 int // denominator
	maxReduceIterations int // maximum number of iterations for reducing the fraction
}

// Return the minimum number of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Return the maximum number of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Return the absolute value of an integer
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// Try reducing the fraction up to a maximum number of iterations which
// is stored in the fraction itself
func (my *Fraction) reduce() {
	// Equal above and below are 1
	if my.top == my.bot {
		my.top = 1
		my.bot = 1
		return
	}
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

// Return the fraction as a float64. Some precision may be lost.
func (my *Fraction) Float64() float64 {
	return float64(my.top) / float64(my.bot)
}

// Return the fraction as an int, not rounded
func (my *Fraction) Int() int {
	return int(my.Float64())
}

// Round of the fraction to an int
func (my *Fraction) Round() int {
	return int(my.Float64() + 0.5)
}

// Return the fraction as a string
func (my *Fraction) String() string {
	return strconv.Itoa(my.top) + "/" + strconv.Itoa(my.bot)
}

// If both the numinator and denuminator are negative, make them positive
func (my *Fraction) prettyNegative() {
	if (my.top < 0 && my.bot < 0) || (my.top > 0 && my.bot < 0) {
		my.top = -my.top
		my.bot = -my.bot
	}
}

// Multiply by another fraction, don't return anything
func (my *Fraction) Multiply(x Fraction) {
	my.top *= x.top
	my.bot *= x.bot
	my.reduce()
}

// Divide by another fraction, don't return anything
func (my *Fraction) Divide(x Fraction) {
	my.top *= x.bot
	my.bot *= x.top
	my.reduce()
}

// Add another fraction, don't return anything
func (my *Fraction) Add(x Fraction) {
	my.top = my.top*x.bot + x.top*my.bot
	my.bot = my.bot * x.bot
	my.reduce()
}

// Subtract another fraction, don't return anything
func (my *Fraction) Sub(x Fraction) {
	my.top = my.top*x.bot - x.top*my.bot
	my.bot = my.bot * x.bot
	my.reduce()
}

// Multiply with an integer and reduce the result
func (my *Fraction) MultiplyInt(x int) {
	my.top *= x
	my.reduce()
}

// Divide by an integer and reduce the result
func (my *Fraction) DivideInt(x int) {
	my.bot *= x
	my.reduce()
}

// Add an integer and reduce the result
func (my *Fraction) AddInt(x int) {
	my.top += my.bot * x
	my.reduce()
}

// Subtract an integer and reduce the result
func (my *Fraction) SubInt(x int) {
	my.top -= my.bot * x
	my.reduce()
}

/*
 * Create a new fraction
 *
 * Takes a numinator and a denumintator 
 *
 * The maximum number of iterations that should be used for reducing the
 * fraction during calculations is set to the default value.
 */
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

// Change the maximum number of iterations that should be used for reductions
func (my *Fraction) SetMaxReduceIterations(maxReduceIterations int) {
	my.maxReduceIterations = maxReduceIterations
}

// Split up a fraction into an integer part, and the rest as another fraction
func (my *Fraction) Splitup() (int, Fraction) {
	i := my.Int()
	clone := *my
	clone.SubInt(i)
	return i, clone
}

// Try to convert a float to a fraction
// Takes a float and a maximum number of iterations to find the fraction
// The maximum number of iterations can be -1 to iterate as much as necessary
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

// Create a new fraction that is "N/1"
func NewFractionFromInt(num int) Fraction {
	return NewFraction(num, 1)
}

// Creates a new fraction that is "0/1"
func NewFractionFromVoid() Fraction {
	return NewFraction(0, 1)
}

// Creates a new fraction from a string on the form "N/D"
// where N is the numinator and D is the denumintator,
// for example: "1/2". Will panic at invalid strings.
func NewFractionFromString(exp string) Fraction {
	top := 0
	bot := 1
	parts := strings.Split(exp, "/")
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
