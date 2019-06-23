package mockdb

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type (
	// DB represents a in-memory datastore "backend"
	DB struct {
		database map[string][]string
	}
)

// NewMockDB creates a DB object
func NewMockDB(db map[string][]string) *DB {
	return &DB{db}
}

// GetValue retrieves a random value from the datastore backend
func (m *DB) GetValue(cat string) string {
	length := len(m.database[cat])
	return m.database[cat][rand.Intn(length)]
}

// GetValues retrieves a list of values from the datastore backend
func (m *DB) GetValues(cats []string) map[string][]string {
	results := make(map[string][]string)
	for _, cat := range cats {
		results[cat] = append(results[cat], m.GetValue(cat))
	}
	return results
}

// AddValue adds a value to the backend. Here it's ephemeral (i.e. it's only around until the memory is released)
func (m *DB) AddValue(cat string, value string) error {
	m.database[cat] = append(m.database[cat], value)
	log.Println(time.Now().UnixNano(), "[WARN]", fmt.Sprintf("Value '%s' was added to the %s datastore ephemerally.", value, cat))
	return nil
}
