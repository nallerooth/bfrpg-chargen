package bf

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/nallerooth/bfrpg-chargen/dice"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

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
func (c *Character) SelectClass() error {
	if c.race == nil {
		return fmt.Errorf("Character must have a race before selecing a class is possible")
	}
	classes := AllClasses()
	availableClasses := make([]*Class, 0, len(classes))

	for _, cl := range classes {
		valid := true
		for a := 0; a < len(c.attributes); a++ {
			val := c.attributes[a].value
			if val < cl.reqAttributes[a] {
				valid = false
				break
			}
		}

		if valid {
			// Class is valid given character attributes, check race
			if c.race.availableClasses&cl.id != 0 {
				availableClasses = append(availableClasses, cl)
			}
		}
	}

	if len(availableClasses) == 0 {
		return fmt.Errorf("Impossible character >> no classes available")
	} else if len(availableClasses) > 1 {
		c.class = availableClasses[rng.Intn(len(availableClasses))]
	} else {
		c.class = availableClasses[0]
	}

	return nil
}

// SelectRace filters the available races based on the character's attributes
// and assigns a valid race at random
func (c *Character) SelectRace() {
	races := AllRaces()
	availableRaces := make([]*Race, 0, len(races))

	for _, r := range races {
		valid := true
		for a := 0; a < len(c.attributes); a++ {
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
		c.race = availableRaces[rng.Intn(len(availableRaces))]
	} else {
		c.race = availableRaces[0]
	}
}

func (c *Character) RollHealth(maxHealthFirstLevel bool) {
	var hp uint
	classHD := float64(c.class.hitDie.base)
	raceHD := float64(c.race.maxHitDieBase)
	hitDie := int(math.Min(classHD, raceHD))
	conMod := c.attributes[AttrCon].mod

	// Limit rolls to first 9 levels
	numRolls := int(math.Min(float64(c.level), 9.0))

	if maxHealthFirstLevel {
		hp += uint(hitDie + conMod)
		numRolls--
	}

	for i := 0; i < numRolls; i++ {
		// (Rolled hit die plus con mod) OR 1
		hp += uint(math.Max(float64(rng.Intn(hitDie)+1+conMod), 1.0))
	}

	if c.level > 9 {
		multiplier := c.level - 9
		hp += multiplier * c.class.hitDie.extraHP
	}

	c.hitpoints = hp
}
