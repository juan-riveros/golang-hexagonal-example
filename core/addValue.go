package core

// AddValue adds a value to a category
func (g *Generator) AddValue(cat string, val string) error {
	return g.repo.AddValue(cat, val)
}
