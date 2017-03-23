[![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/khezen/darwin)
[![Build Status](http://img.shields.io/travis/Khezen/darwin.svg?style=flat-square)](https://travis-ci.org/Khezen/darwin) [![codecov](https://img.shields.io/codecov/c/github/Khezen/darwin/master.svg?style=flat-square)](https://codecov.io/gh/Khezen/darwin)
[![Go Report Card](https://goreportcard.com/badge/github.com/khezen/darwin?style=flat-square)](https://goreportcard.com/report/github.com/khezen/darwin)

# Darwin
Go implementation of Genetic Algorithm


## IIndividual

```golang
type IIndividual interface {
	Fitness() float32
	SetFitness(float32)
}
```

## Population
```golang
type IPopulation interface {
	sort.Interface
	Sort()
    	Max() IIndividual
	Min() IIndividual
	Extremums() (IIndividual, IIndividual)

    	Append(...IIndividual)
   	RemoveAt(int) error
	Remove(...IIndividual)
	Replace(int, IIndividual) error
    	Get(int) (IIndividual, error)
    	PickCouple() (index1 int, indiv1 IIndividual, index2 int, indiv2 IIndividual, err error)
	Has(...IIndividual) bool
	IndexOf(IIndividual) (int, error)

	Cap() int
	SetCap(int) error
	Truncate(int) error
}

```

## Operators

```golang
type IEvaluater interface {
	Evaluate(IIndividual) (Fitness float32)
}
```

```golang
type ISelecter interface {
	Select(pop IPopulation, survivorsSize int) (IPopulation, error)
}
```

```golang
type ICrosser interface {
	Cross(individual1, individual2 IIndividual) IIndividual
}
```

```golang
type IMutater interface {
	Mutate(IIndividual) IIndividual
}
```

## Lifecycle

```golang

type ILifecycle interface {
	Generation(pop IPopulation, survivorSizeForSelection int, mutationProbability float32) (IPopulation, error)
}

type Lifecycle struct {
	Selecter  ISelecter
	Crosser   ICrosser
	Mutater   IMutater
	Evaluater IEvaluater
}


```
