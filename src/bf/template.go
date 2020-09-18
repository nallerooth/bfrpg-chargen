package bf

import (
	"fmt"
	"io"
	"log"
	"text/template"
)

var textTmpl = template.Must(template.ParseFiles("src/bf/templates/cli.tmpl"))

func (c *Character) Text(out io.Writer) {
	presentation := struct {
		Name        string
		Class       string
		Level       uint
		Race        string
		StrBase     string
		StrMod      string
		DexBase     string
		DexMod      string
		ConBase     string
		ConMod      string
		IntBase     string
		IntMod      string
		WisBase     string
		WisMod      string
		ChaBase     string
		ChaMod      string
		Saves       *[5]uint
		SavesMod    [5]string
		AttackBonus string
		AC          uint
		HP          uint
		OpenDoors   int
		FindTraps   uint
		SecretDoors uint
		Skills      *ThiefSkills
		SpellSlots  *spellSlots
	}{
		Name:    fmt.Sprintf("%s", c.name),
		Class:   c.class.name,
		Level:   c.level,
		Race:    c.race.name,
		StrBase: fmt.Sprintf("%2d", c.attributes[AttrStr].value),
		StrMod:  fmt.Sprintf("%+d", c.attributes[AttrStr].mod),
		DexBase: fmt.Sprintf("%2d", c.attributes[AttrDex].value),
		DexMod:  fmt.Sprintf("%+d", c.attributes[AttrDex].mod),
		ConBase: fmt.Sprintf("%2d", c.attributes[AttrCon].value),
		ConMod:  fmt.Sprintf("%+d", c.attributes[AttrCon].mod),
		IntBase: fmt.Sprintf("%2d", c.attributes[AttrInt].value),
		IntMod:  fmt.Sprintf("%+d", c.attributes[AttrInt].mod),
		WisBase: fmt.Sprintf("%2d", c.attributes[AttrWis].value),
		WisMod:  fmt.Sprintf("%+d", c.attributes[AttrWis].mod),
		ChaBase: fmt.Sprintf("%2d", c.attributes[AttrCha].value),
		ChaMod:  fmt.Sprintf("%+d", c.attributes[AttrCha].mod),
		Saves:   &c.class.saves[c.level-1],
		SavesMod: [5]string{
			fmt.Sprintf("%+d", c.race.savesMod[0]),
			fmt.Sprintf("%+d", c.race.savesMod[1]),
			fmt.Sprintf("%+d", c.race.savesMod[2]),
			fmt.Sprintf("%+d", c.race.savesMod[3]),
			fmt.Sprintf("%+d", c.race.savesMod[4]),
		},
		AttackBonus: fmt.Sprintf("%+d", c.attackBonus),
		AC:          c.armorClass,
		HP:          c.hitpoints,
		OpenDoors:   1 + c.attributes[AttrStr].mod,
		FindTraps:   1,
		SecretDoors: 1,
		Skills:      nil,
		SpellSlots:  nil,
	}

	if c.race.name == "Dwarf" {
		presentation.FindTraps = 2
	}

	if c.race.name == "Elf" {
		presentation.SecretDoors += 1
	}
	if c.attributes[AttrInt].value >= 15 {
		presentation.SecretDoors += 1
	}

	if presentation.OpenDoors < 1 {
		presentation.OpenDoors = 1
	}

	// Add thief skills if available
	// TODO: Handle multiclass
	if c.class.id == ClassThief {
		presentation.Skills = c.class.ThiefSkillsForLevel(c.level)
	}

	// Add spell slots if available
	// TODO: Add multiclass
	if c.class.id == ClassCleric || c.class.id == ClassMagicUser {
		slots := c.class.spellSlotsForLevel(c.level)
		presentation.SpellSlots = &slots
	}

	err := textTmpl.Execute(out, presentation)
	if err != nil {
		log.Fatalln(err)
	}

}
