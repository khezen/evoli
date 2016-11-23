package individual

// Phenotyped is an interface for any entity gifted with a Phenotype
type Phenotyped interface {
	Resilience() float32
	SetResilience(float32)
}
