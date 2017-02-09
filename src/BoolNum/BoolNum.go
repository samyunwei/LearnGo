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

func main() {
	fmt.Println(BitFlag(0), Active, Send, flag, Receive, flag|Receive)
}
