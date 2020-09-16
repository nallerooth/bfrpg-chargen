package bf

import (
	"fmt"
	"io"
	"log"
	"text/template"
)

type saves struct {
}

var textTmpl = template.Must(template.ParseFiles("src/bf/templates/cli.tmpl"))

func (c *Character) Text(out io.Writer) {
	presentation := struct {
		Name       string
		Class      string
		Level      uint
		Race       string
		StrBase    string
		StrMod     string
		DexBase    string
		DexMod     string
		ConBase    string
		ConMod     string
		IntBase    string
		IntMod     string
		WisBase    string
		WisMod     string
		ChaBase    string
		ChaMod     string
		Saves      saves
		Skills     *ThiefSkills
		SpellSlots *spellSlots
	}{
		Name:       fmt.Sprintf("%s", c.name),
		Class:      c.class.name,
		Level:      c.level,
		Race:       c.race.name,
		StrBase:    fmt.Sprintf("%2d", c.attributes[AttrStr].value),
		StrMod:     fmt.Sprintf("%+d", c.attributes[AttrStr].mod),
		DexBase:    fmt.Sprintf("%2d", c.attributes[AttrDex].value),
		DexMod:     fmt.Sprintf("%+d", c.attributes[AttrDex].mod),
		ConBase:    fmt.Sprintf("%2d", c.attributes[AttrCon].value),
		ConMod:     fmt.Sprintf("%+d", c.attributes[AttrCon].mod),
		IntBase:    fmt.Sprintf("%2d", c.attributes[AttrInt].value),
		IntMod:     fmt.Sprintf("%+d", c.attributes[AttrInt].mod),
		WisBase:    fmt.Sprintf("%2d", c.attributes[AttrWis].value),
		WisMod:     fmt.Sprintf("%+d", c.attributes[AttrWis].mod),
		ChaBase:    fmt.Sprintf("%2d", c.attributes[AttrCha].value),
		ChaMod:     fmt.Sprintf("%+d", c.attributes[AttrCha].mod),
		Saves:      saves{},
		Skills:     nil,
		SpellSlots: nil,
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
