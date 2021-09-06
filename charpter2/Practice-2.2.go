// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"os"
	"strconv"

	"charpter2/lengthconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		l, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		i := lengthconv.Inch(l)
		m := lengthconv.Meter(l)
		fmt.Printf("%s = %s, %f = %s\n",
			i, lengthconv.IToM(i), m, lengthconv.MToI(m))
	}
}
