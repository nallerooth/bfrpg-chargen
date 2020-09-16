package dice

import (
	"fmt"
	"math/rand"
	"regexp"
	"sort"
	"strconv"
	"time"
)

type DiceRoller struct {
	rng *rand.Rand
}

type roll struct {
	rolled   []int
	numSides int
	numDice  int
	modifier int
	drop     int
}

func (r *roll) Sum() uint {
	var sum int

	for _, d := range r.rolled {
		sum += d
	}

	return uint(sum + r.modifier)
}

func (r *roll) Execute(rng *rand.Rand) {
	for i := 0; i < r.numDice; i++ {
		r.rolled = append(r.rolled, rng.Intn(r.numSides)+1)
	}

	sort.Ints(r.rolled)

	if r.drop > 0 {
		r.rolled = r.rolled[r.drop:]
	}
}

func New() *DiceRoller {
	d := DiceRoller{}
	source := rand.NewSource(time.Now().UnixNano())
	d.rng = rand.New(source)

	return &d
}

func (d *DiceRoller) rollDie(sides int) uint {
	return uint(d.rng.Intn(sides))
}

func parseDiceString(dicestr string) (*roll, error) {
	var err error

	//pattern := regexp.MustCompile(`^(\d+)d(\d+)([+-]?)(\d*)$`)
	pattern := regexp.MustCompile(`^(\d+)d(\d+)(?:([+-]?)(\d*)?)(?:(d?)(\d)?)$`)
	matches := pattern.FindStringSubmatch(dicestr)

	r := roll{}

	if r.numDice, err = strconv.Atoi(matches[1]); err != nil {
		return nil, err
	}

	if r.numSides, err = strconv.Atoi(matches[2]); err != nil {
		return nil, err
	}

	if matches[4] != "" {
		if r.modifier, err = strconv.Atoi(matches[4]); err != nil {
			return nil, err
		}
		if matches[3] == "-" {
			r.modifier *= -1
		}
	}

	if matches[5] != "" { // Drop
		if r.drop, err = strconv.Atoi(matches[6]); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

func (d *DiceRoller) Roll(dicestr string) uint {
	var r *roll
	var err error

	if r, err = parseDiceString(dicestr); err != nil {
		fmt.Printf("DiceRoller: Error: %s\n", err)
	}

	r.Execute(d.rng)

	//fmt.Printf("%s >> amount: %d, sides: %d, modifier: %+d :: %v :: %d\n",
	//dicestr, r.numDice, r.numSides, r.modifier, r.rolled, r.Sum())

	return r.Sum()
}
