package main

import (
	"os"

	"github.com/nallerooth/bfrpg-chargen/bf"
	"github.com/nallerooth/bfrpg-chargen/dice"
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
	character.RollHealth(false) // Don't max health for lvl 1
	character.Text(os.Stdout)

}
