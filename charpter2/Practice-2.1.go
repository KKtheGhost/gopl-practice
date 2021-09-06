package main

type Celsuis float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsuis = -273.15
	FreezingC     Celsuis = 0
	BoilingC      Celsuis = 100
)

func CToF(c Celsuis) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsuis { return Celsuis((f - 32) * 5 / 9) }
func KToC(k Kelvin) Celsuis     { return Celsuis(k - 273.15) }
func CToK(c Celsuis) Kelvin     { return Kelvin(c + 273.15) }
