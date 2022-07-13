package weightconv

import "fmt"

type Feet float64
type Inch float64
type Meter float64

const ()

func (f Feet) String() string  { return fmt.Sprintf("%g", f) }
func (i Inch) String() string  { return fmt.Sprintf("%g", i) }
func (m Meter) String() string { return fmt.Sprintf("%g", m) }
