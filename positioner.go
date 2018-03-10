package evoli

// Positioner produces a new individual from
// (indiv) its previous version,
// (pBest) its best version since it exists,
// (gBest) and the best of the poopulation or local neighborhood
type Positioner interface {
	Position(indiv, pBest, gBest Individual, learningCoef1, learningCoef2 float64) (Individual, error)
}
