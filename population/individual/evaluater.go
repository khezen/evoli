package individual

// Evaluater computes and set individual resilience
type Evaluater interface {
	Evaluate(Interface) (resilience float32)
}
