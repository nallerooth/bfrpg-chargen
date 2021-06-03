package bf

const (
	weaponSmall = iota
	weaponMedium
	weaponLarge
)

type armor struct {
	name              string
	ac                int
	acMod             int
	weight            uint
	restrictedClasses int
}

type weapon struct {
	name              string
	damage            string
	size              int
	weight            int
	restrictedClasses int
}

type rangedWeapon struct {
	weapon
	rangeShort  int
	rangeMedium int
	rangeLong   int
}

func AllArmors() []armor {
	return []armor{
		armor{"None", 11, 0, 0, 0},
		armor{"Leather Armor", 13, 0, 15, ClassMagicUser},
		armor{"Chain Mail", 15, 0, 40, ClassMagicUser | ClassThief},
		armor{"Plate Mail", 17, 0, 50, ClassMagicUser | ClassThief},
		armor{"Shield", 0, 1, 5, ClassMagicUser | ClassThief},
	}
}

func AllMeleeWeapons() []weapon {
	return []weapon{
		weapon{"Hand Axe", "1d6", weaponSmall, 5, ClassCleric | ClassMagicUser},
		weapon{"Battle Axe", "1d8", weaponMedium, 7, ClassCleric | ClassMagicUser},
		weapon{"Great Axe", "1d10", weaponLarge, 15, ClassCleric | ClassMagicUser},
		weapon{"Dagger", "1d4", weaponSmall, 1, ClassCleric},
		weapon{"Shortsword", "1d6", weaponSmall, 3, ClassCleric | ClassMagicUser},
		weapon{"Longsword", "1d8", weaponMedium, 4, ClassCleric | ClassMagicUser},
		weapon{"Two-Handed Sword", "1d10", weaponLarge, 4, ClassCleric | ClassMagicUser},
		weapon{"Warhammer", "1d6", weaponSmall, 4, ClassMagicUser},
		weapon{"Mace", "1d8", weaponMedium, 10, ClassMagicUser},
		weapon{"Maul", "1d10", weaponLarge, 16, ClassMagicUser},
		weapon{"Club", "1d4", weaponMedium, 2, 0},
		weapon{"Quarterstaff", "1d6", weaponLarge, 4, ClassMagicUser},
		weapon{"Pole Arm", "1d10", weaponLarge, 15, ClassCleric | ClassMagicUser},
		weapon{"Spear, one handed", "1d6", weaponMedium, 5, ClassCleric | ClassMagicUser},
		weapon{"Spear, two handed", "1d8", weaponMedium, 5, ClassCleric | ClassMagicUser},
	}
}

func AllRangedWeapons() []rangedWeapon {
	return []rangedWeapon{
		rangedWeapon{weapon{"Shortbow", "1d6", weaponMedium, 2, ClassCleric | ClassMagicUser}, 50, 100, 150},
		rangedWeapon{weapon{"Longbow", "1d8", weaponLarge, 3, ClassCleric | ClassMagicUser}, 70, 140, 210},
		rangedWeapon{weapon{"Light Crossbow", "1d6", weaponMedium, 7, ClassCleric | ClassMagicUser}, 60, 120, 180},
		rangedWeapon{weapon{"Heavy Crossbow", "1d8", weaponLarge, 14, ClassCleric | ClassMagicUser}, 80, 160, 240},
		rangedWeapon{weapon{"Sling", "1d4", weaponSmall, 0, ClassMagicUser}, 30, 60, 90},
		rangedWeapon{weapon{"Dagger", "1d4", weaponSmall, 1, ClassCleric}, 10, 20, 30},
		rangedWeapon{weapon{"Hand Axe", "1d4", weaponSmall, 5, ClassCleric | ClassMagicUser}, 10, 20, 30},
		rangedWeapon{weapon{"Spear", "1d6", weaponMedium, 5, ClassCleric | ClassMagicUser}, 10, 20, 30},
		rangedWeapon{weapon{"Warhammer", "1d6", weaponSmall, 4, ClassMagicUser}, 10, 20, 30},
	}
}
