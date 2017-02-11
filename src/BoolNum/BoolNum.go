package main

import (
	"fmt"
	"strings"
	"math"
)

type BitFlag int

const (
	Active  BitFlag = 1 << iota
	Send
	Receive
)

var flag = Active | Send

func (flag BitFlag) String() string {
	var flags []string
	if flag&Active == Active {
		flags = append(flags, "Active")
	}
	if flag&Send == Send {
		flags = append(flags, "Send")
	}
	if flag&Receive == Receive {
		flags = append(flags, "Receive")
	}
	if len(flags) > 0 {
		return fmt.Sprintf("%d(%s)", int(flag), strings.Join(flags, "|"))
	}
	return "0()"

}

func Uint8FromInt(x int) (uint8, error) {
	if 0 <= x && x <= math.MaxUint8 {
		return uint8(x), nil
	}
	return 0, fmt.Errorf("%d is out of the uint8 range", x)
}

func EqualFloat(x, y, limit float64) bool {
	if limit <= 0.0 {
		limit = math.SmallestNonzeroFloat64
	}
	return math.Abs(x-y) <= (limit * math.Min(math.Abs(x), math.Abs(y)))
}

func EqualFloatPrec(x, y float64, decimals int) bool {
	a := fmt.Sprintf("%.*f", decimals, x)
	b := fmt.Sprintf("%.*f", decimals, y)
	return len(a) == len(b) && a == b
}

func IntFromFloat64(x float64) int {
	if math.MinInt32 <= x && x <= math.MaxFloat32 {
		whole, fraction := math.Modf(x)
		if fraction >= 0.5 {
			whole++
		}
		return int(whole)
	}
	panic(fmt.Sprintf("%g is out of the int32 range", x))
}


func main() {
	//fmt.Println(BitFlag(0), Active, Send, flag, Receive, flag|Receive)

	x, y := 0.0, 0.0
	for i := 0; i < 10; i++ {
		x += 0.1
		if i%2 == 0 {
			y += 0.2
		} else {
			fmt.Printf("%-5t %-5t %-5t %-5t", x == y, EqualFloat(x, y, -1), EqualFloat(x, y, 0.000000000001), EqualFloatPrec(x, y, 6))
			fmt.Println(x, y)
		}
	}
}
