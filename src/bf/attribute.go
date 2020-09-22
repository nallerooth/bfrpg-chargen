package bf

import "fmt"

func isBetween(min, max int, value uint) bool {
	return (int(value) >= min && int(value) <= max)
}

func calculateBonus(value uint) (bonus int, err error) {
	if !isBetween(3, 18, value) {
		return bonus, fmt.Errorf("Error setting attribute value of %d", value)
	}

	switch {
	case value == 3:
		bonus = -3
	case isBetween(4, 5, value):
		bonus = -2
	case isBetween(6, 8, value):
		bonus = -1
	case isBetween(9, 12, value):
		bonus = 0
	case isBetween(13, 15, value):
		bonus = 1
	case isBetween(16, 17, value):
		bonus = 2
	case value == 18:
		bonus = 3
	}

	return bonus, nil
}

func (a *attribute) Set(value uint) error {
	var bonus int
	var err error

	if bonus, err = calculateBonus(value); err != nil {
		return err
	}

	a.value = value
	a.mod = int(bonus)

	return nil
}
