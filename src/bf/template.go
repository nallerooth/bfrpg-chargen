package bf

import (
	"fmt"
	"io"
	"log"
	"strconv"
	"text/template"
)

var textTmpl *template.Template

func initTextTemplate(funcs *template.FuncMap) {
	if textTmpl == nil {
		textTmpl = template.Must(template.New("cli.tmpl").Funcs(*funcs).ParseFiles("src/bf/templates/cli.tmpl"))
	}

}

func (c *Character) Text(out io.Writer) {
	funcs := template.FuncMap{
		"add": func(strA, strB string) string {
			var a, b int
			var err error
			if a, err = strconv.Atoi(strA); err != nil {
				log.Fatalln("Error parsing string value to int:", strA)
			}
			if b, err = strconv.Atoi(strB); err != nil {
				log.Fatalln("Error parsing string value to int:", strB)
			}
			return fmt.Sprintf("%d", a+b)
		},
		"addInt": func(a, b int) int {
			return a + b
		},
	}
	initTextTemplate(&funcs)

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
		Spells      [6]*[]string
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
		AttackBonus: fmt.Sprintf("%+d", c.class.attackBonus[c.level-1]),
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
		presentation.SecretDoors++
	}
	if c.attributes[AttrInt].value >= 15 {
		presentation.SecretDoors++
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

		spells := [6]*[]string{}
		for i, amount := range slots {
			if amount > 0 {
				lvlSpells, err := c.class.GetRandomSpellsOfLevel(uint(amount), uint(i+1))
				if err != nil {
					log.Fatalln(err)
				}
				spells[i] = &lvlSpells
			}
		}
		presentation.Spells = spells
	}

	err := textTmpl.Execute(out, presentation)
	if err != nil {
		log.Fatalln(err)
	}

}
