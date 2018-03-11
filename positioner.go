package evoli

// Positioner produces a new individual from
// its previous version(indiv),
// its best version since it exists(pBest),
// the best of the poopulation(gBest)  or local neighborhood,
// and the learning coefficients(c1,c2). typical values are c1=c2=2
type Positioner interface {
	Position(indiv, pBest, gBest Individual, c1, c2 float64) (Individual, error)
}
