package evoli

// Evolution -
type Evolution interface {
	Next(pop Population) (Population, error)
}
