package args

import (
	"fmt"
	"os"
)

func Echo1() {
	var s, spc string
	for i := 1; i < len(os.Args); i++ {
		s += spc + os.Args[i]
		spc = "\n"
	}
	fmt.Println(s)
}

func Exec1() {
	var s, spc string
	for _, arg := range os.Args {
		s += spc + arg
		spc = "\n"
	}
	fmt.Println(s)
}
