package cmd

import (
	"fmt"

	core "../../core"
)

type (
	// TUI represents a Text User Interface
	TUI struct {
		generator core.SceneGenerator
	}
)

// NewTUI creates a TUI object
func NewTUI(generator core.SceneGenerator) *TUI {
	return &TUI{generator}
}

// DisplayScene outputs a representation of the Scene to STDOUT
func (T *TUI) DisplayScene() {
	scn := T.generator.GenerateScene()
	fmt.Println("Locale:", scn.Locale)
	fmt.Println("Mood:", scn.Mood)
	fmt.Println("Time of Day:", scn.TimeOfDay)
	fmt.Println("Time Period:", scn.TimePeriod)
	fmt.Println("Characters:")
	for _, char := range scn.Characters {
		fmt.Printf("\t%+v\n", T.displayCharacter(char))
	}
	fmt.Println()
}

func (T *TUI) displayCharacter(chara core.Character) string {
	return fmt.Sprintf("%s %s %s", chara.AgeGroup, chara.Gender, chara.Profession)
}

// Run initializes anything necessary to run the UI (i.e. here is just calls the display func.)
func (T *TUI) Run() {
	T.DisplayScene()
}
