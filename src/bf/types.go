package bf

const (
	ClassFighter   = 1
	ClassCleric    = 2
	ClassMagicUser = 4
	ClassThief     = 8
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
	saves            [20][5]uint
	thiefSkills      [20]ThiefSkills
}

type Race struct {
	name             string
	availableClasses int
	minAttributes    [6]uint
	maxAttributes    [6]uint
	maxHitDieBase    uint
	savesMod         [5]uint
	// TODO: Handle weapon sizes
	specialAbilities []string // TODO: Add special abilities
}

type Character struct {
	name       [20]byte
	class      *Class
	race       *Race
	attributes [6]attribute
	level      uint
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
