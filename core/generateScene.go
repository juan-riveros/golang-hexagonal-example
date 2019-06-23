package core

import "math/rand"

// GenerateScene creates a random scene
func (g *Generator) GenerateScene() Scene {
	var scene Scene
	numCharacters := rand.Intn(3) + rand.Intn(2) + 2

	for i := 0; i < numCharacters; i++ {
		scene.Characters = append(scene.Characters, g.generateCharacter())
	}
	scene.Locale = g.repo.GetValue("locale")
	scene.TimeOfDay = g.repo.GetValue("timeofday")
	scene.TimePeriod = g.repo.GetValue("timeperiod")
	scene.Mood = g.repo.GetValue("mood")
	return scene
}

func (g *Generator) generateCharacter() Character {
	var chara Character

	chara.Profession = g.repo.GetValue("profession")
	chara.AgeGroup = g.repo.GetValue("agegroup")
	chara.Gender = g.repo.GetValue("gender")
	return chara
}
