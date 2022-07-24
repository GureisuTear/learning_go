package weightconv

func IToF(i Inch) Feet { return Feet(i / 12) }

func FToI(f Feet) Inch { return Inch(f * 12) }

func IToM(i Inch) Meter { return Meter(i / 39.37) }

func MToI(m Meter) Inch { return Inch(m * 39.37) }

func MToF(m Meter) Feet { return Feet(m * 3.281) }

func FToM(f Feet) Meter { return Meter(f / 3.281) }
