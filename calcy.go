package calcy

import "math/rand"

type Calculator interface {
	Calculate(a, b int) int
}

type Addition struct{}
type Subtraction struct{}
type Multiplication struct{}
type Division struct{}
type Bogus struct{}

func (this Addition) Calculate(a, b int) int       { return a + b }
func (this Subtraction) Calculate(a, b int) int    { return a - b }
func (this Multiplication) Calculate(a, b int) int { return a * b }
func (this Division) Calculate(a, b int) int       { return a / b }
func (this Bogus) Calculate(a, b int) int          { return rand.Int() + a + b }
