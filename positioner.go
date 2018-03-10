package evoli

// Positioner produces a new individual from
// (indiv) its previous version,
// (pBest) its best version since it exists,
// (gBest) the best of the poopulation or local neighborhood,
// (c1,c2) and the learning coefficients. typical values are c1=c2=2
type Positioner interface {
	Position(indiv, pBest, gBest Individual, c1, c2 float64) (Individual, error)
}
