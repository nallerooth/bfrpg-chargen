package bf

const (
	ClassFighter = iota
	ClassCleric
	ClassMagicUser
	ClassThief
)

const (
	AttrStr = iota
	AttrDex
	AttrCon
	AttrInt
	AttrWis
	AttrCha
)

type attribute struct {
	value uint
	mod   int
}

type HitDie struct {
	base    uint
	extraHP uint
}

type Class struct {
	id               int
	name             string
	desc             string
	hitDie           HitDie
	reqAttributes    [6]uint
	primaryReqAttr   int
	experienceLevels [20]int
	numSpells        [20][6]int
	thiefSkills      [20]ThiefSkills
}

type Race struct {
	name             string
	availableClasses [4]int
	minAttributes    [6]uint
	maxAttributes    [6]uint
	maxHitDieBase    uint
	// TODO: Handle weapon sizes
	specialAbilities []string // TODO: Add special abilities
}

type Character struct {
	name       [20]byte
	class      *Class
	race       *Race
	attributes [6]attribute
	level      uint
	Saves      [5]uint
}

type ThiefSkills struct {
	OpenLocks    uint
	RemoveTraps  uint
	PickPockets  uint
	MoveSilently uint
	ClimbWalls   uint
	Hide         uint
	Listen       uint
}

type spellSlots [6]int
