[![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/khezen/evoli)
[![Build Status](http://img.shields.io/travis/khezen/evoli.svg?style=flat-square)](https://travis-ci.org/khezen/evoli) [![codecov](https://img.shields.io/codecov/c/github/khezen/evoli/master.svg?style=flat-square)](https://codecov.io/gh/khezen/evoli)
[![Go Report Card](https://goreportcard.com/badge/github.com/khezen/evoli?style=flat-square)](https://goreportcard.com/report/github.com/khezen/evoli)

# *evoli*
Genetic Algorithm written in Go
```golang
import "github.com/khezen/evoli"
```

```golang
type Individual interface {
	Fitness() float64
	SetFitness(float64)
}
```

```golang
type Population interface {
	sort.Interface
	Sort()
    	Max() Individual
	Min() Individual

    	Add(...Individual)
   	RemoveAt(int) error
	Remove(...Individual)
	Replace(int, Individual) error
    	Get(int) (Individual, error)
	Has(...Individual) bool
	IndexOf(Individual) (int, error)

	Cap() int
	SetCap(int) error
	Truncate(int) error
}

```

```golang
type Evaluater interface {
	Evaluate(Individual) (Fitness float64, err error)
}
```

```golang
type Selecter interface {
	Select(pop Population, survivorsSize int) (Population, error)
}
```

```golang
type Crosser interface {
	Cross(individual1, individual2 Individual) (Individual, error)
}
```

```golang
type Mutater interface {
	Mutate(Individual) (Individual, error)
}
```

```golang

type Lifecycle interface {
	Next(pop Population) (Population, error)
}

func NewLifecycle(s Selecter, survivorSize int, c Crosser, m Mutater, mutationProbability float64, e Evaluater) Lifecycle

```
