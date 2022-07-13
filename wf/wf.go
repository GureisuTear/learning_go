package wf

import (
	"fmt"
	"os"
	"strconv"

	"github.com/GureisuTear/learning_go/General-Purpose_Unit-Conversion/weightconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		i := weightconv.Inch(t)
		f := weightconv.Feet(t)
		m := weightconv.Meter(t)
		fmt.Printf("%s = %s,\n %s = %s,\n %s = %s\n", i, weightconv.IToF(i), f, weightconv.FToI(f), m, weightconv.MToI(m))
	}
}
