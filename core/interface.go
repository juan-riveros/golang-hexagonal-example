package core

type (
	// DataRepository represents some datastore backend
	DataRepository interface {
		GetValue(string) string
		GetValues([]string) map[string][]string
		AddValue(string, string) error
	}

	// SceneGenerator represents the API that UIs can access
	SceneGenerator interface {
		GenerateScene() Scene
		AddValue(string, string) error
	}
)

type (
	// Generator represents the generator
	Generator struct {
		repo DataRepository
	}

	// Scene represents a scene
	Scene struct {
		Locale     string
		Characters []Character
		TimeOfDay  string
		TimePeriod string
		Mood       string
	}

	// Character represents a character
	Character struct {
		Profession string
		AgeGroup   string
		Gender     string
	}
)

// NewGenerator creates a new Generator Object
func NewGenerator(repo *DataRepository) *Generator {
	return &Generator{*repo}
}
