package main

import (
	"fmt"
)

func printColors(n interface{}, c interface{}) {
	if n.(names).nameN == "Sandra" {
		fmt.Println("Sandra likes color " + c.(colors).nameC)
	} else {
		fmt.Println("Arturo likes color " + c.(colors).nameC)
	}
}

type names struct {
	nameN string
}

type colors struct {
	nameC string
}

func main() {
	printColors(names{"Sandra"}, colors{"Pink"})
	printColors(names{"Arturo"}, colors{"Red"})
}
