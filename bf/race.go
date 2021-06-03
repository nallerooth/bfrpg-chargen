package bf

func AllRaces() []*Race {
	r := []*Race{
		&Race{
			name:             "Dwarf",
			availableClasses: ClassCleric | ClassFighter | ClassThief,
			minAttributes: [6]uint{
				3, 3, 9, 3, 3, 3, // Con 9
			},
			maxAttributes: [6]uint{
				18, 18, 18, 18, 18, 17, // Cha 17
			},
			maxHitDieBase: 8,
			savesMod:      [5]uint{4, 4, 4, 3, 4},
		},
		&Race{
			name:             "Elf",
			availableClasses: ClassCleric | ClassFighter | ClassMagicUser | ClassThief, // TODO: add multiclass
			minAttributes: [6]uint{
				3, 3, 3, 9, 3, 3, // Int 9
			},
			maxAttributes: [6]uint{
				18, 18, 17, 18, 18, 18, // Con 17
			},
			maxHitDieBase: 6,
			savesMod:      [5]uint{0, 2, 1, 0, 2},
		},
		&Race{
			name:             "Halfling",
			availableClasses: ClassCleric | ClassFighter | ClassThief,
			minAttributes: [6]uint{
				3, 9, 3, 3, 3, 3, // Dex 9
			},
			maxAttributes: [6]uint{
				17, 18, 18, 18, 18, 18, // Str 17
			},
			maxHitDieBase: 6,
			savesMod:      [5]uint{4, 4, 4, 3, 4},
		},
		&Race{
			name:             "Human",
			availableClasses: ClassCleric | ClassFighter | ClassMagicUser | ClassThief,
			minAttributes: [6]uint{
				3, 3, 3, 3, 3, 3, // None
			},
			maxAttributes: [6]uint{
				18, 18, 18, 18, 18, 18, // None
			},
			maxHitDieBase: 8,
			savesMod:      [5]uint{0, 0, 0, 0, 0},
		},
	}

	return r
}
