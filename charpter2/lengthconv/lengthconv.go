package lengthconv

import "fmt"

type Meter float64
type Inch float64

func (m Meter) String() string { return fmt.Sprintf("%g m", m) }
func (i Inch) String() string  { return fmt.Sprintf("%g inch", i) }
