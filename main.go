package main

import (
	"flag"
	"fmt"
	"os"

	cmd "./application/cmd"
	web "./application/web"
	core "./core"
	jsondb "./infrastructure/jsondb"
	mockdb "./infrastructure/mockdb"
	sqlite3db "./infrastructure/sqlite3db"

	"math/rand"
	"time"
)

// UI represents some form of user interface
type UI interface {
	Run()
}

func init() {
	rand.Seed(time.Now().UnixNano())
	m := rand.Int63n(6) + rand.Int63n(6) + rand.Int63n(6)
	rand.Seed(time.Now().UnixNano() + m)
}

func selectDB(dbflag *string, db *core.DataRepository) {
	switch *dbflag {
	case "json":
		*db = jsondb.NewJSONDB("./test.json")
	case "mock":
		*db = mockdb.NewMockDB(map[string][]string{
			"locale":     {"Abandoned Ship", "Abundant Shire", "Aetherflux Reservour", "Altar to the Moon", "Ancient Tomb", "Azure Bay", "Dessert Desert"},
			"profession": {"teacher", "officer", "maid", "official", "ranger", "sales associate", "driver"},
			"agegroup":   {"young", "older", "ancient"},
			"gender":     {"male", "female"},
			"timeofday":  {"dawn", "morning", "late morning", "midday", "lunch hour", "afternoon", "evening", "night", "midnight", "late night", "twilight"},
			"timeperiod": {"1800s", "1900", "1910", "1920", "1930", "1940", "1950", "1960", "1970", "1980", "1990", "2000", "2010", "2020"},
			"mood":       {"somber", "hectic", "frantic", "quiet", "joyful", "comfy", "tense", "dinosaur"},
		})
	case "sqlite3":
		*db = sqlite3db.NewSQLDB("./sqlite3.db")
	default:
		fmt.Println("-database [json|mock|sqlite3]")
		os.Exit(1)
	}
}

func selectUI(uiflag *string, ui *UI, generator core.SceneGenerator) {
	switch *uiflag {
	case "cmd":
		*ui = cmd.NewTUI(generator)
	case "web":
		*ui = web.NewWebUI(generator)
	default:
		fmt.Println("-ui [cmd|web]")
		os.Exit(1)
	}
}

func main() {
	var db core.DataRepository
	var gener core.SceneGenerator
	var ui UI

	var dbflag *string
	var uiflag *string

	dbflag = flag.String("database", "sqlite3", "-database [json|mock|sqlite3]")
	uiflag = flag.String("ui", "web", "-ui [cmd|web]")

	flag.Parse()

	selectDB(dbflag, &db)
	gener = core.NewGenerator(&db)
	selectUI(uiflag, &ui, gener)
	ui.Run()

}
