package bf

import "fmt"

func isBetween(min, max int, value uint) bool {
	return (int(value) >= min && int(value) <= max)
}

func calculateBonus(value uint) (bonus int, err error) {
	if !isBetween(3, 18, value) {
		return bonus, fmt.Errorf("Error setting attribute value of %d", value)
	}

	if value == 3 {
		bonus = -3
	} else if isBetween(4, 5, value) {
		bonus = -2
	} else if isBetween(6, 8, value) {
		bonus = -1
	} else if isBetween(9, 12, value) {
		bonus = 0
	} else if isBetween(13, 15, value) {
		bonus = 1
	} else if isBetween(16, 17, value) {
		bonus = 2
	} else if value == 18 {
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
