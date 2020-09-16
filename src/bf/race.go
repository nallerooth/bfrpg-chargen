package bf

func AllRaces() []*Race {
	r := []*Race{
		&Race{
			name: "Dwarf",
			availableClasses: [4]int{
				1, // ClassCleric
				1, // ClassFighter
				0, // ClassMagicUser
				1, // ClassThief
			},
			minAttributes: [6]uint{
				3, 3, 9, 3, 3, 3, // Con 9
			},
			maxAttributes: [6]uint{
				18, 18, 18, 18, 18, 17, // Cha 17
			},
			maxHitDieBase: 8,
		},
		&Race{
			name: "Elf",
			availableClasses: [4]int{
				1, // ClassCleric
				1, // ClassFighter
				1, // ClassMagicUser
				1, // ClassThief
				// TODO: Add multiclass
			},
			minAttributes: [6]uint{
				3, 3, 3, 9, 3, 3, // Int 9
			},
			maxAttributes: [6]uint{
				18, 18, 17, 18, 18, 18, // Con 17
			},
			maxHitDieBase: 6,
		},
		&Race{
			name: "Halfling",
			availableClasses: [4]int{
				1, // ClassCleric
				1, // ClassFighter
				0, // ClassMagicUser
				1, // ClassThief
			},
			minAttributes: [6]uint{
				3, 9, 3, 3, 3, 3, // Dex 9
			},
			maxAttributes: [6]uint{
				17, 18, 18, 18, 18, 18, // Str 17
			},
			maxHitDieBase: 6,
		},
		&Race{
			name: "Human",
			availableClasses: [4]int{
				1, // ClassCleric
				1, // ClassFighter
				1, // ClassMagicUser
				1, // ClassThief
			},
			minAttributes: [6]uint{
				3, 3, 3, 3, 3, 3, // None
			},
			maxAttributes: [6]uint{
				18, 18, 18, 18, 18, 18, // None
			},
			maxHitDieBase: 8,
		},
	}

	return r
}
