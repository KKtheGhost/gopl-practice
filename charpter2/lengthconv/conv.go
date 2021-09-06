package lengthconv

// CToF converts a Celsius temperature to Fahrenheit.
func MToI(m Meter) Inch { return Inch(m * 3.2808) }

// FToC converts a Fahrenheit temperature to Celsius.
func IToM(i Inch) Meter { return Meter(i / 3.2808) }
