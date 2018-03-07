package darwin

// Evolution -
type Evolution interface {
	Next(pop Population) (Population, error)
}
