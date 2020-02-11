// Package types performs Celsius and Fahrenheit temperature computations.
package types

import (
	"fmt"
)

// Celsius type definition for celsius temparature representations (float64)
type Celsius float64

// Fahrenheit type definition for fahrenheit temparature representations (float64)
type Fahrenheit float64

// Celsius-related constants
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

var yetanother Fahrenheit = 94

var another Celsius = 94

// CToF conversion from Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC conversion from Celsius to Fahrenheit
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
