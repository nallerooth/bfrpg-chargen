package main

import (
	"fmt"
	"os"

	"./bf"
	"./dice"
)

func main() {
	method := ""
	if len(os.Args) == 2 {
		// TODO: Usage info
		method = os.Args[1]
	}

	dice := dice.New()

	character := bf.Character{}
	for {
		character.RollAttributes(dice, method)
		character.SelectRace()

		if err := character.SelectClass(); err == nil {
			break
		}

	}
	character.SetLevel(8)
	//character.Print()
	character.Text(os.Stdout)

	fmt.Println("")
}
