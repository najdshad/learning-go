package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// working with os.Args
func cmdArgs() {
	fArgs := make([]float64, 0, 2)
	for _, arg := range os.Args[1:3] {
		v, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Println("error")
			fmt.Println(err)
		}
		fArgs = append(fArgs, v)
	}
	fmt.Println(math.Pow(fArgs[0], fArgs[1]))
}
