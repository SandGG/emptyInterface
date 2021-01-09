package main

import (
	"fmt"
)

func printColors(n interface{}, c interface{}) {
	var name = n.(names).nameN
	if name == "Sandra" {
		fmt.Println(name + " likes color " + c.(colors).nameC)
	} else {
		fmt.Println(name + " likes color " + c.(colors).nameC)
	}
}

type names struct {
	nameN string
}

type colors struct {
	nameC string
}

func main() {
	var nam = names{"Sandra"}
	var col = colors{"Pink"}
	printColors(nam, col)

	nam = names{"Arturo"}
	col = colors{"Red"}
	printColors(nam, col)
}
