package main

import (
	"fmt"
	"github.com/jongyunha/celeritas"
)

func main() {
	result := celeritas.TestFunc(1, 1)
	fmt.Println(result)

	result2 := celeritas.TestFund2(10, 2)
	fmt.Println(result2)
}
