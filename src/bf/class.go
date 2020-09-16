package bf

func AllClasses() []*Class {
	return []*Class{
		NewFighter(),
		NewCleric(),
		NewMagicUser(),
		NewThief(),
	}
}

func NewCleric() *Class {
	c := Class{}
	c.id = ClassCleric
	c.name = "Cleric"
	c.primaryReqAttr = AttrWis
	c.hitDie = HitDie{6, 1}
	c.reqAttributes[AttrWis] = 9
	c.experienceLevels = [20]int{
		0,
		1_500,
		3_000,
		6_000,
		12_000,

		24_000,
		48_000,
		90_000,
		180_000,
		270_000,

		360_000,
		450_000,
		540_000,
		630_000,
		720_000,

		810_000,
		900_000,
		990_000,
		1_080_000,
		1_170_000,
	}
	c.numSpells = [20][6]int{
		{0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0},
		{2, 0, 0, 0, 0, 0},
		{2, 1, 0, 0, 0, 0},
		{2, 2, 0, 0, 0, 0},

		{2, 2, 1, 0, 0, 0},
		{3, 2, 2, 0, 0, 0},
		{3, 2, 2, 1, 0, 0},
		{3, 3, 2, 2, 0, 0},
		{3, 3, 2, 2, 1, 0},

		{4, 3, 3, 2, 2, 0},
		{4, 4, 3, 2, 2, 1},
		{4, 4, 3, 3, 2, 2},
		{4, 4, 4, 3, 2, 2},
		{4, 4, 4, 3, 3, 2},

		{5, 4, 4, 3, 3, 2},
		{5, 5, 4, 3, 3, 2},
		{5, 5, 4, 4, 3, 3},
		{6, 5, 4, 4, 3, 3},
		{6, 5, 5, 4, 3, 3},
	}

	return &c
}

func NewFighter() *Class {
	c := Class{}
	c.id = ClassFighter
	c.name = "Fighter"
	c.primaryReqAttr = AttrStr
	c.hitDie = HitDie{8, 2}
	c.reqAttributes[AttrStr] = 9
	c.experienceLevels = [20]int{
		0,
		2_000,
		4_000,
		8_000,
		16_000,

		32_000,
		64_000,
		120_000,
		240_000,
		360_000,

		480_000,
		600_000,
		720_000,
		840_000,
		960_000,

		1_080_000,
		1_200_000,
		1_320_000,
		1_440_000,
		1_560_000,
	}

	return &c
}

func NewMagicUser() *Class {
	c := Class{}
	c.id = ClassMagicUser
	c.name = "Magic-User"
	c.primaryReqAttr = AttrInt
	c.hitDie = HitDie{4, 1}
	c.reqAttributes[AttrInt] = 9
	c.experienceLevels = [20]int{
		0,
		2_500,
		5_000,
		10_000,
		20_000,

		40_000,
		80_000,
		150_000,
		300_000,
		450_000,

		600_000,
		750_000,
		900_000,
		1_050_000,
		1_200_000,

		1_350_000,
		1_500_000,
		1_650_000,
		1_800_000,
		1_950_000,
	}
	c.numSpells = [20][6]int{
		{1, 0, 0, 0, 0, 0},
		{2, 0, 0, 0, 0, 0},
		{2, 1, 0, 0, 0, 0},
		{2, 2, 0, 0, 0, 0},
		{2, 2, 1, 0, 0, 0},

		{3, 2, 2, 0, 0, 0},
		{3, 2, 2, 1, 0, 0},
		{3, 3, 2, 2, 0, 0},
		{3, 3, 2, 2, 1, 0},
		{4, 3, 3, 2, 2, 0},

		{4, 4, 3, 2, 2, 1},
		{4, 4, 3, 3, 2, 2},
		{4, 4, 4, 3, 2, 2},
		{4, 4, 4, 3, 3, 2},
		{5, 4, 4, 3, 3, 2},

		{5, 5, 4, 3, 3, 2},
		{5, 5, 4, 4, 3, 3},
		{6, 5, 4, 4, 3, 3},
		{6, 5, 5, 4, 3, 3},
		{6, 5, 5, 4, 4, 3},
	}

	return &c
}

func NewThief() *Class {
	c := Class{}
	c.id = ClassThief
	c.name = "Thief"
	c.primaryReqAttr = AttrDex
	c.hitDie = HitDie{4, 2}
	c.reqAttributes[AttrDex] = 9
	c.experienceLevels = [20]int{
		0,
		1_250,
		2_500,
		5_000,
		10_000,

		20_000,
		40_000,
		75_000,
		150_000,
		225_000,

		300_000,
		375_000,
		450_000,
		525_000,
		600_000,

		675_000,
		750_000,
		825_000,
		900_000,
		975_000,
	}
	c.thiefSkills = [20]ThiefSkills{
		{25, 20, 30, 25, 80, 10, 30},
		{30, 25, 35, 30, 81, 15, 34},
		{35, 30, 40, 35, 82, 20, 38},
		{40, 35, 45, 40, 83, 25, 42},
		{45, 40, 50, 45, 84, 30, 46},

		{50, 45, 55, 50, 85, 35, 50},
		{55, 50, 60, 55, 86, 40, 54},
		{60, 55, 65, 60, 87, 45, 58},
		{65, 60, 70, 65, 88, 50, 62},
		{68, 63, 74, 68, 89, 53, 65},

		{71, 66, 78, 71, 90, 56, 68},
		{74, 69, 82, 74, 91, 59, 71},
		{77, 72, 86, 77, 92, 62, 74},
		{80, 75, 90, 80, 93, 65, 77},
		{83, 78, 94, 83, 94, 68, 80},

		{84, 79, 95, 85, 95, 69, 83},
		{85, 80, 96, 87, 96, 70, 86},
		{86, 81, 97, 89, 97, 71, 89},
		{87, 82, 98, 91, 98, 72, 92},
		{88, 83, 99, 93, 99, 73, 95},
	}

	return &c
}

func (c *Class) spellSlotsForLevel(level uint) spellSlots {
	// TODO: Limit level range
	return spellSlots(c.numSpells[level-1])
}

func (c *Class) ThiefSkillsForLevel(level uint) *ThiefSkills {
	// TODO: Limit level range
	return &c.thiefSkills[level-1]
}

func (c *Class) PrimaryAttribute() int {
	return c.primaryReqAttr
}
