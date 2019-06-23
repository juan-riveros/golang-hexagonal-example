package jsondb

import (
	"math/rand"

	util "../../util"
)

type (
	// DB represents a JSON datastore backend
	DB struct {
		database map[string][]string
		filepath string
	}
)

// NewJSONDB creates a new DB object
func NewJSONDB(filepath string) *DB {
	db := DB{}
	util.ParseJSON(filepath, &db.database)
	db.filepath = filepath
	return &db
}

// GetValue retrieves a random value from a category
func (db *DB) GetValue(cat string) string {
	length := len(db.database[cat])
	return db.database[cat][rand.Intn(length)]
}

// GetValues retrieves a list of values from a category
func (db *DB) GetValues(cats []string) map[string][]string {
	results := make(map[string][]string)
	for _, cat := range cats {
		results[cat] = append(results[cat], db.GetValue(cat))
	}
	return results
}

// AddValue adds a value to a category and returns an error
func (db *DB) AddValue(cat string, value string) error {
	db.database[cat] = append(db.database[cat], value)
	err := util.SaveJSON(db.filepath, db.database)
	return err
}
