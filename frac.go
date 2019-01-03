package num

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

const defaultMaxReduceIterations = 400

type Frac struct {
	top                 int64 // numerator
	bot                 int64 // denominator
	maxReduceIterations int64 // maximum number of iterations for reducing the fraction
	exactfloat          bool  // if the float64 representation will be exact
}

var (
	Zero = &Frac{0, 1, defaultMaxReduceIterations, true}
	One  = &Frac{1, 1, defaultMaxReduceIterations, true}

	ErrDivByZero = errors.New("division by zero")
)

// New creates a new fractional number.
// Takes a numinator and a denominator.
// The maximum number of iterations that should be used for reducing the
// fraction during calculations is set to the default value.
func New(num, dom int64) (*Frac, error) {
	if dom == 0 {
		return nil, ErrDivByZero
	}
	frac := &Frac{
		top:                 num,
		bot:                 dom,
		maxReduceIterations: defaultMaxReduceIterations,
		exactfloat:          true,
	}
	frac.reduce()
	return frac, nil
}

// MustNew must create a new fractional number.
// If it is not possible, no error will be returned and it will panic.
func MustNew(num, dom int64) *Frac {
	n, err := New(num, dom)
	if err != nil {
		panic(err)
	}
	return n
}

// Try to convert a float to a fraction
// Takes a float and a maximum number of iterations to find the fraction
// The maximum number of iterations can be -1 to iterate as much as necessary
// Returns a bool that is True if the maximum number of iterations has not been reached
func NewFromFloat64(f float64, maxIterations int64) *Frac {
	// Thanks stackoverflow.com/questions/95727/how-to-convert-floats-to-human-readable-fractions
	var (
		num     int64   = 1
		dom     int64   = 1
		result  float64 = 1
		counter int64
		exact   = true
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
	frac, _ := New(num, dom)
	frac.SetExact(exact)
	return frac
}

func (my *Frac) SetExact(exact bool) {
	my.exactfloat = exact
}

// Create a new fraction that is "N/1"
func NewFromInt(num int64) *Frac {
	// Will never divide on 0, so it's safe to ignore the error
	frac, _ := New(num, 1)
	return frac
}

// NewZero returns a fraction that is "0/1"
func NewZero() *Frac {
	return Zero
}

// Creates a new fraction from a string on the form "N/D", where N is the
// numerator and D is the denominator. For example: "1/2" or "3/8".
func NewFromString(exp string) (*Frac, error) {
	var (
		top int64
		bot int64 = 1
	)
	if !strings.Contains(exp, "/") {
		return &Frac{}, errors.New("This doesn't look like a fraction: " + exp)
	}
	parts := strings.Split(exp, "/")
	if len(parts) != 2 {
		return &Frac{}, errors.New("This doesn't look like a fraction: " + exp)
	}
	if value, err := strconv.Atoi(parts[0]); err == nil {
		top = int64(value)
	} else {
		return &Frac{}, errors.New("Invalid numerator: " + parts[0])
	}
	if value, err := strconv.Atoi(parts[1]); err == nil {
		bot = int64(value)
	} else {
		return &Frac{}, errors.New("Invalid denominator: " + parts[1])
	}
	return New(top, bot)
}

// Creates a new fraction from a rational number (big.Rat)
func NewFromRat(rat *big.Rat) *Frac {
	// Ignore error since *big.Rat denom can't be 0
	frac, _ := New(rat.Num().Int64(), rat.Denom().Int64())
	return frac
}

// Returns a rational number (big.Rat)
func (my *Frac) Rat() *big.Rat {
	return big.NewRat(my.top, my.bot)
}

// Try reducing the fraction up to a maximum number of iterations which
// is stored in the fraction itself
func (my *Frac) reduce() {
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
func (my *Frac) Float64() float64 {
	return float64(my.top) / float64(my.bot)
}

func (my *Frac) ExactFloat64() bool {
	return my.exactfloat
}

// Return the fraction as an int, not rounded
func (my *Frac) Int64() int64 {
	return int64(my.Float64())
}

// Round of the fraction to an int
func (my *Frac) Round() int64 {
	return int64(my.Float64() + 0.5)
}

// Return the fraction as a string
func (my *Frac) String() string {
	return fmt.Sprintf("%d/%d", my.top, my.bot)
}

// If both the numinator and denuminator are negative, make them positive
func (my *Frac) prettyNegative() {
	if (my.bot < 0) && (my.top != 0) {
		my.top = -my.top
		my.bot = -my.bot
	}
}

// Multiply by another fraction, don't return anything
func (my *Frac) Multiply(x *Frac) {
	my.top *= x.top
	my.bot *= x.bot
	my.reduce()
}

// Divide by another fraction, don't return anything
func (my *Frac) Divide(x *Frac) {
	my.top *= x.bot
	my.bot *= x.top
	my.reduce()
}

// Add another fraction, don't return anything
func (my *Frac) Add(x *Frac) {
	my.top = my.top*x.bot + x.top*my.bot
	my.bot = my.bot * x.bot
	my.reduce()
}

// Subtract another fraction, don't return anything
func (my *Frac) Sub(x *Frac) {
	my.top = my.top*x.bot - x.top*my.bot
	my.bot = my.bot * x.bot
	my.reduce()
}

// Multiply with an integer and reduce the result
func (my *Frac) MultiplyInt(x int64) {
	my.top *= x
	my.reduce()
}

// Divide by an integer and reduce the result
func (my *Frac) DivideInt(x int64) {
	my.bot *= x
	my.reduce()
}

// Add an integer and reduce the result
func (my *Frac) AddInt(x int64) {
	my.top += my.bot * x
	my.reduce()
}

// Subtract an integer and reduce the result
func (my *Frac) SubInt(x int64) {
	my.top -= my.bot * x
	my.reduce()
}

// Change the maximum number of iterations that should be used for reductions
func (my *Frac) SetMaxReduceIterations(maxReduceIterations int64) {
	my.maxReduceIterations = maxReduceIterations
}

// Split up a fraction into an integer part, and the rest as another fraction
func (my *Frac) Splitup() (int64, *Frac) {
	i := my.Int64()
	clone := *my
	clone.SubInt(i)
	return i, &clone
}
