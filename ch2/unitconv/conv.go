package unitconv

import "fmt"

type Feet float64
type Meters float64
type Pounds float64
type Kilograms float64

func (f Feet) String() string      { return fmt.Sprintf("%gft", f) }
func (m Meters) String() string    { return fmt.Sprintf("%gm", m) }
func (p Pounds) String() string    { return fmt.Sprintf("%glb", p) }
func (k Kilograms) String() string { return fmt.Sprintf("%gkg", k) }

func FToM(f Feet) Meters      { return Meters(f * 0.3048) }
func MToF(m Meters) Feet      { return Feet(m * 3.28084) }
func KToP(k Kilograms) Pounds { return Pounds(k * 2.20462) }
func PToK(p Pounds) Kilograms { return Kilograms(p * 0.453592) }
