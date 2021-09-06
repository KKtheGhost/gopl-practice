package main

import "fmt"

func main() {
	const FreezingF, BoilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", FreezingF, fToC(FreezingF))
	fmt.Printf("%g°F = %g°C\n", BoilingF, fToC(BoilingF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
