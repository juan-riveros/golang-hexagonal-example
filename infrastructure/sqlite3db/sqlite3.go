package sqlite3db

import (
	"database/sql"
	"fmt"

	// All about the hex architecture
	_ "github.com/mattn/go-sqlite3"

	util "../../util"
)

type (
	// DB represents a JSON datastore backend
	DB struct {
		database *sql.DB
		filepath string
	}
)

// NewSQLDB creates a new DB object
func NewSQLDB(filepath string) *DB {
	db := DB{}
	sql, err := sql.Open("sqlite3", filepath)
	util.CheckErr(true, err, "Could not open sqlite3 DB (%s)", filepath)
	db.database = sql
	db.filepath = filepath
	return &db
}

// GetValue retrieves a random value from a category
func (db *DB) GetValue(cat string) string {
	var value string
	rows, err := db.database.Query("select value from keyvalues where key = ? order by random() limit 1", cat)
	util.CheckErr(false, err, "Unable to get row from sqlite3 DB")
	for rows.Next() {
		rows.Scan(&value)
	}
	return value

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
	stmt, err := db.database.Prepare("insert into keyvalues (key, value) values (?, ?)")
	util.CheckErr(false, err, "Failed to prepare sqlite3DB")

	res, err := stmt.Exec(cat, value)
	util.CheckErr(false, err, "Failed to insert {%s: %s} into sqlite3 DB", cat, value)
	id, err := res.LastInsertId()
	util.CheckErr(false, err, "Failed to insert {%s: %s} into sqlite3 DB", cat, value)
	fmt.Printf("sqlite3 insert result: %+v\n", id)

	return err
}
