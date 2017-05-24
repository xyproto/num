package gofractions

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

const defaultMaxReduceIterations = 400

type Fraction struct {
	top                 int64 // numerator
	bot                 int64 // denominator
	maxReduceIterations int64 // maximum number of iterations for reducing the fraction
	exactfloat          bool  // if the float64 representation will be exact
}

/*
 * Create a new fraction
 *
 * Takes a numinator and a denomintator
 *
 * The maximum number of iterations that should be used for reducing the
 * fraction during calculations is set to the default value.
 */
func NewFraction(num int64, dom int64) (*Fraction, error) {
	if dom == 0 {
		return nil, errors.New("Division by zero")
	}
	frac := &Fraction{
		top:                 num,
		bot:                 dom,
		maxReduceIterations: defaultMaxReduceIterations,
		exactfloat:          true,
	}
	frac.reduce()
	return frac, nil
}

var (
	Zero = &Fraction{0, 1, defaultMaxReduceIterations, true}
	One  = &Fraction{1, 1, defaultMaxReduceIterations, true}
)

// Try to convert a float to a fraction
// Takes a float and a maximum number of iterations to find the fraction
// The maximum number of iterations can be -1 to iterate as much as necessary
// Returns a bool that is True if the maximum number of iterations has not been reached
func NewFractionFromFloat64(f float64, maxIterations int64) *Fraction {
	// Thanks stackoverflow.com/questions/95727/how-to-convert-floats-to-human-readable-fractions
	var (
		num     int64   = 1
		dom     int64   = 1
		result  float64 = 1
		counter int64   = 0
		exact   bool    = true
	)
	for result != f {
		if result < f {
			num++
		} else {
			dom++
			num = int64(f * float64(dom))
		}
		result = float64(num) / float64(dom)
		if counter == maxIterations {
			exact = false
			break
		}
		counter++
	}
	// Will never divide on 0, so it's safe to ignore the error
	frac, _ := NewFraction(num, dom)
	frac.SetExact(exact)
	return frac
}

func (my *Fraction) SetExact(exact bool) {
	my.exactfloat = exact
}

// Create a new fraction that is "N/1"
func NewFractionFromInt(num int64) *Fraction {
	// Will never divide on 0, so it's safe to ignore the error
	frac, _ := NewFraction(num, 1)
	return frac
}

// Creates a new fraction that is "0/1"
func NewZeroFraction() *Fraction {
	// Will never divide on 0, so it's safe to ignore the error
	frac, _ := NewFraction(0, 1)
	return frac
}

// Creates a new fraction from a string on the form "N/D", where N is the
// numerator and D is the denominator. For example: "1/2" or "3/8".
func NewFractionFromString(exp string) (*Fraction, error) {
	var (
		top int64 = 0
		bot int64 = 1
	)
	if !strings.Contains(exp, "/") {
		return &Fraction{}, errors.New("This doesn't look like a fraction: " + exp)
	}
	parts := strings.Split(exp, "/")
	if len(parts) != 2 {
		return &Fraction{}, errors.New("This doesn't look like a fraction: " + exp)
	}
	if value, err := strconv.Atoi(parts[0]); err == nil {
		top = int64(value)
	} else {
		return &Fraction{}, errors.New("Invalid numerator: " + parts[0])
	}
	if value, err := strconv.Atoi(parts[1]); err == nil {
		bot = int64(value)
	} else {
		return &Fraction{}, errors.New("Invalid denominator: " + parts[1])
	}
	return NewFraction(top, bot)
}

// Creates a new fraction from a rational number (big.Rat)
func NewFractionFromRat(rat *big.Rat) *Fraction {
	// Ignore error since *big.Rat denom can't be 0
	frac, _ := NewFraction(rat.Num().Int64(), rat.Denom().Int64())
	return frac
}

// Returns a rational number (big.Rat)
func (my *Fraction) Rat() *big.Rat {
	return big.NewRat(my.top, my.bot)
}

// Try reducing the fraction up to a maximum number of iterations which
// is stored in the fraction itself
func (my *Fraction) reduce() {
	// Equal above and below are 1
	if my.top == my.bot {
		my.top = 1
		my.bot = 1
		my.exactfloat = true
		return
	}
	counter := int64(0)
	for trydiv := min(abs(my.top), abs(my.bot)); trydiv >= 2; trydiv-- {
		if (my.top/trydiv)*trydiv == my.top && (my.bot/trydiv)*trydiv == my.bot {
			my.top /= trydiv
			my.bot /= trydiv
		}
		if counter == my.maxReduceIterations {
			my.exactfloat = false
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

func (my *Fraction) ExactFloat64() bool {
	return my.exactfloat
}

// Return the fraction as an int, not rounded
func (my *Fraction) Int64() int64 {
	return int64(my.Float64())
}

// Round of the fraction to an int
func (my *Fraction) Round() int64 {
	return int64(my.Float64() + 0.5)
}

// Return the fraction as a string
func (my *Fraction) String() string {
	return fmt.Sprintf("%v/%v", my.top, my.bot)
}

// If both the numinator and denuminator are negative, make them positive
func (my *Fraction) prettyNegative() {
	if (my.bot < 0) && (my.top != 0) {
		my.top = -my.top
		my.bot = -my.bot
	}
}

// Multiply by another fraction, don't return anything
func (my *Fraction) Multiply(x *Fraction) {
	my.top *= x.top
	my.bot *= x.bot
	my.reduce()
}

// Divide by another fraction, don't return anything
func (my *Fraction) Divide(x *Fraction) {
	my.top *= x.bot
	my.bot *= x.top
	my.reduce()
}

// Add another fraction, don't return anything
func (my *Fraction) Add(x *Fraction) {
	my.top = my.top*x.bot + x.top*my.bot
	my.bot = my.bot * x.bot
	my.reduce()
}

// Subtract another fraction, don't return anything
func (my *Fraction) Sub(x *Fraction) {
	my.top = my.top*x.bot - x.top*my.bot
	my.bot = my.bot * x.bot
	my.reduce()
}

// Multiply with an integer and reduce the result
func (my *Fraction) MultiplyInt(x int64) {
	my.top *= x
	my.reduce()
}

// Divide by an integer and reduce the result
func (my *Fraction) DivideInt(x int64) {
	my.bot *= x
	my.reduce()
}

// Add an integer and reduce the result
func (my *Fraction) AddInt(x int64) {
	my.top += my.bot * x
	my.reduce()
}

// Subtract an integer and reduce the result
func (my *Fraction) SubInt(x int64) {
	my.top -= my.bot * x
	my.reduce()
}

// Change the maximum number of iterations that should be used for reductions
func (my *Fraction) SetMaxReduceIterations(maxReduceIterations int64) {
	my.maxReduceIterations = maxReduceIterations
}

// Split up a fraction into an integer part, and the rest as another fraction
func (my *Fraction) Splitup() (int64, *Fraction) {
	i := my.Int64()
	clone := *my
	clone.SubInt(i)
	return i, &clone
}
