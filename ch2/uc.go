package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/k4rtik/gopl/ch2/tempconv"
	"github.com/k4rtik/gopl/ch2/unitconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unitconv: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		ft := unitconv.Feet(t)
		m := unitconv.Meters(t)
		p := unitconv.Pounds(t)
		k := unitconv.Kilograms(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
		fmt.Printf("%s = %s, %s = %s\n", ft, unitconv.FToM(ft), m, unitconv.MToF(m))
		fmt.Printf("%s = %s, %s = %s\n", p, unitconv.PToK(p), k, unitconv.KToP(k))
	}
}
