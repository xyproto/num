package num

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

const defaultMaxReduceIterations = 400

type Frac struct {
	top                 int64 // numerator
	bot                 int64 // denominator
	maxReduceIterations int   // maximum number of iterations for reducing the fraction
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
func NewFromFloat64(f float64, maxIterations int) *Frac {
	// Thanks stackoverflow.com/questions/95727/how-to-convert-floats-to-human-readable-fractions
	var (
		num     int64   = 1
		dom     int64   = 1
		result  float64 = 1
		counter int
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

func (f *Frac) SetExact(exact bool) {
	f.exactfloat = exact
}

// Create a new fraction that is "N/1"
func NewFromInt(num int) *Frac {
	// Will never divide on 0, so it's safe to ignore the error
	frac, _ := New(int64(num), 1)
	return frac
}

// Create a new fraction that is "N/1"
func NewFromInt64(num int64) *Frac {
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
func (f *Frac) Rat() *big.Rat {
	return big.NewRat(f.top, f.bot)
}

// Try reducing the fraction up to a maximum number of iterations which
// is stored in the fraction itself
func (f *Frac) reduce() {
	// Equal above and below are 1
	if f.top == f.bot {
		f.top = 1
		f.bot = 1
		f.exactfloat = true
		return
	}
	var counter int
	for trydiv := min(abs(f.top), abs(f.bot)); trydiv >= 2; trydiv-- {
		if (f.top/trydiv)*trydiv == f.top && (f.bot/trydiv)*trydiv == f.bot {
			f.top /= trydiv
			f.bot /= trydiv
		}
		if counter == f.maxReduceIterations {
			f.exactfloat = false
			break
		}
		counter++
	}
	f.prettyNegative()
}

// Return the fraction as a float64. Some precision may be lost.
func (f *Frac) Float64() float64 {
	return float64(f.top) / float64(f.bot)
}

func (f *Frac) ExactFloat64() bool {
	return f.exactfloat
}

// Return the fraction as an int, not rounded
func (f *Frac) Int() int {
	return int(f.Float64())
}

// Return the fraction as an int, not rounded
func (f *Frac) Int64() int64 {
	return int64(f.Float64())
}

// Round of the fraction to an int
func (f *Frac) Round() int64 {
	return int64(f.Float64() + 0.5)
}

// Return the fraction as a string
func (f *Frac) String() string {
	return fmt.Sprintf("%d/%d", f.top, f.bot)
}

// If both the numinator and denuminator are negative, make them positive
func (f *Frac) prettyNegative() {
	if (f.bot < 0) && (f.top != 0) {
		f.top = -f.top
		f.bot = -f.bot
	}
}

// Multiply by another fraction, don't return anything
func (f *Frac) Mul(x *Frac) {
	f.top *= x.top
	f.bot *= x.bot
	f.reduce()
}

// Multiply two fractions and return the result
func Mul(a, b *Frac) *Frac {
	top := a.top * b.top
	bot := a.bot * b.bot
	if bot == 0 {
		panic("Multiplying with a number that has 0 as the denominator!")
	}
	return MustNew(top, bot)
}

// Divide by another fraction, don't return anything
func (f *Frac) Div(x *Frac) {
	f.top *= x.bot
	f.bot *= x.top
	f.reduce()
}

// Divide two fractions and return the result
func Div(a, b *Frac) (*Frac, error) {
	top := a.top * b.bot
	bot := a.bot * b.top
	return New(top, bot)
}

// Add another fraction, don't return anything
func (f *Frac) Add(x *Frac) {
	f.top = f.top*x.bot + x.top*f.bot
	f.bot = f.bot * x.bot
	f.reduce()
}

// Add two fractions and return the result
func Add(a, b *Frac) *Frac {
	top := a.top*b.bot + b.top*a.bot
	bot := a.bot * b.bot
	return MustNew(top, bot)
}

// Subtract another fraction, don't return anything
func (f *Frac) Sub(x *Frac) {
	f.top = f.top*x.bot - x.top*f.bot
	f.bot = f.bot * x.bot
	f.reduce()
}

// Subtract two fractions and return the result
func Sub(a, b *Frac) *Frac {
	top := a.top*b.bot - b.top*a.bot
	bot := a.bot * b.bot
	return MustNew(top, bot)
}

// Multiply with an integer and reduce the result
func (f *Frac) MulInt(x int) {
	f.top *= int64(x)
	f.reduce()
}

// Multiply with an integer and reduce the result
func MulInt(f *Frac, x int) (*Frac, error) {
	return New(f.top*int64(x), f.bot)
}

// Divide by an integer and reduce the result
func (f *Frac) DivInt(x int) {
	f.bot *= int64(x)
	f.reduce()
}

// Divide by an integer and reduce the result
func DivInt(f *Frac, x int) (*Frac, error) {
	return New(f.top, f.bot*int64(x))
}

// Add an integer and reduce the result
func (f *Frac) AddInt(x int) {
	f.top += f.bot * int64(x)
	f.reduce()
}

// Add an int64 and reduce the result
func (f *Frac) AddInt64(x int64) {
	f.top += f.bot * x
	f.reduce()
}

// Add an integer and reduce the result
func AddInt(f *Frac, x int) *Frac {
	return MustNew(f.top+f.bot*int64(x), f.bot)
}

// Subtract an integer and reduce the result
func (f *Frac) SubInt(x int) {
	f.top -= f.bot * int64(x)
	f.reduce()
}

// Subtract an int64 and reduce the result
func (f *Frac) SubInt64(x int64) {
	f.top -= f.bot * x
	f.reduce()
}

// Subtract an integer and reduce the result
func SubInt(f *Frac, x int) *Frac {
	return MustNew(f.top-f.bot*int64(x), f.bot)
}

// IsZero checks if this fraction is 0
func (f *Frac) IsZero() bool {
	return f.top == 0
}

// Sqrt returns the square root of the number
func Sqrt(f *Frac) *Frac {
	// TODO: Use a numeric algorithm instead
	return NewFromFloat64(math.Sqrt(float64(f.top))/math.Sqrt(float64(f.bot)), f.maxReduceIterations)
}

// Take the square root of this number
func (f *Frac) Sqrt() {
	// TODO: Use a numeric algorithm instead
	x := math.Sqrt(float64(f.top)) / math.Sqrt(float64(f.bot))
	f = NewFromFloat64(x, f.maxReduceIterations)
}

// Multiply this number by itself
func (f *Frac) Square() {
	f.top *= f.top
	f.reduce()
}

// Multiply this number by itself
func Square(f *Frac) *Frac {
	x := f.Copy()
	x.top *= x.top
	x.reduce()
	return x
}

func (f *Frac) Abs() {
	if f.top < 0 {
		f.top = -f.top
	}
	if f.bot < 0 {
		f.bot = -f.bot
	}
}

// Return the absolute value
func Abs(f *Frac) *Frac {
	x := f.Copy()
	if f.top < 0 {
		x.top = -x.top
	}
	if f.bot < 0 {
		x.bot = -x.bot
	}
	x.reduce()
	return x
}

// Change the maximum number of iterations that should be used for reductions
func (f *Frac) SetMaxReduceIterations(maxReduceIterations int) {
	f.maxReduceIterations = maxReduceIterations
}

// Split up a fraction into an integer part, and the rest as another fraction
func (f *Frac) Splitup() (int64, *Frac) {
	i64 := f.Int64()
	clone := *f
	clone.SubInt64(i64)
	return i64, &clone
}

// Copy creates a copy
func (f *Frac) Copy() *Frac {
	return &Frac{
		top:                 f.top,
		bot:                 f.bot,
		maxReduceIterations: f.maxReduceIterations,
		exactfloat:          f.exactfloat,
	}
}
