package bf

import (
	"log"
	"math/rand"

	"../dice"
)

// RollAttributes performes the standard 3D6 attribute rolls, for all attributes
func (c *Character) RollAttributes(dice *dice.DiceRoller, method string) {
	var err error

	switch method {
	case "4d6":
		method = "4d6d1"
	case "2d6":
		method = "2d6+6"
	case "default":
		fallthrough
	default:
		method = "3d6"
	}

	for i := 0; i < 6; i++ {
		if err = c.attributes[i].Set(dice.Roll(method)); err != nil {
			log.Fatalln(err)
		}
	}
}

// SetLevel sets a character's level - without updating hitpoints, learned
// spells, etc.
func (c *Character) SetLevel(level uint) {
	if level < 1 {
		c.level = 1
	} else if level > 20 {
		c.level = 20
	} else {
		c.level = level
	}
}

// SelectClass will set the class best suited to the character's attributes
func (c *Character) SelectClass() { // TODO: Add desired class?
	c.class = NewCleric()
}

func (c *Character) SelectRace() {
	races := AllRaces()
	availableRaces := make([]*Race, 0, len(races))

	for _, r := range races {
		valid := true
		for a := 0; a < len(r.minAttributes); a++ {
			val := c.attributes[a].value
			if val < r.minAttributes[a] || val > r.maxAttributes[a] {
				valid = false
				break
			}
		}

		if valid {
			availableRaces = append(availableRaces, r)
		}
	}

	if len(availableRaces) == 0 {
		c.SelectRace()
	} else if len(availableRaces) > 1 {
		c.race = availableRaces[rand.Intn(len(availableRaces))]
	} else {
		c.race = availableRaces[0]
	}
}
