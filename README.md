[![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/khezen/evoli)
[![Build Status](http://img.shields.io/travis/khezen/evoli.svg?style=flat-square)](https://travis-ci.org/khezen/evoli) [![codecov](https://img.shields.io/codecov/c/github/khezen/evoli/master.svg?style=flat-square)](https://codecov.io/gh/khezen/evoli)
[![Go Report Card](https://goreportcard.com/badge/github.com/khezen/evoli?style=flat-square)](https://goreportcard.com/report/github.com/khezen/evoli)

# *evoli*

Genetic Algorithm and Particle Swarm Optimization written in Go

## Example

### Problem

Given `f(x,y) = cos(x^2 * y^2) * 1/(x^2 * y^2 + 1)`

Find `(x,y)` such as `f(x,y)` reaches its maximum 

Answer `f(0,0) = 1`

### Particle Swarm Optimization

```golang
package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/khezen/evoli"
)

// 3d cosine that gets smaller as you move away from 0,0
func f(x, y float64) float64 {
	d := x*x + y*y
	return math.Cos(d) * (1 / (d/10 + 1))
}

type FIndividual struct {
	v       []float64
	x       []float64
	fitness float64
}

func (i *FIndividual) Equal(other evoli.Individual) bool {
	return i == other
}

func (i *FIndividual) Fitness() float64 {
	return i.fitness
}

func (i *FIndividual) SetFitness(newFitness float64) {
	i.fitness = newFitness
}

type FPositioner struct {
}

func (p *FPositioner) Position(indiv, pBest, gBest evoli.Individual, c1, c2 float64) (evoli.Individual, error) {
	fIndiv, ok1 := indiv.(*FIndividual)
	fPBest, ok2 := pBest.(*FIndividual)
	fGBest, ok3 := gBest.(*FIndividual)
	if !ok1 || !ok2 || !ok3 {
		return nil, fmt.Errorf("invalid individual type")
	}
	newIndiv := FIndividual{
		v: make([]float64, len(fIndiv.v)),
		x: make([]float64, len(fIndiv.v)),
	}
	w := 0.9
	for d := range fIndiv.v {
		rp := rand.Float64()
		rg := rand.Float64()
		newIndiv.v[d] = w*fIndiv.v[d] +
			c1*rp*(fPBest.x[d]-fIndiv.x[d]) +
			c2*rg*(fGBest.x[d]-fIndiv.x[d])

		newIndiv.x[d] = fIndiv.x[d] + newIndiv.v[d]
	}
	return &newIndiv, nil
}

type FEvaluater struct {
}

func (e *FEvaluater) Evaluate(indiv evoli.Individual) (Fitness float64, err error) {
	fIndiv, ok := indiv.(*FIndividual)
	if !ok {
		return 0, fmt.Errorf("invalid individual type")
	}
	return f(fIndiv.x[0], fIndiv.x[1]), nil
}

func main() {
	pop := evoli.NewPopulation(50)
	for i := 0; i < pop.Cap(); i++ {
		x := rand.Float64()*20 - 10
		y := rand.Float64()*20 - 10
		vx := rand.Float64()*20 - 10
		vy := rand.Float64()*20 - 10
		pop.Add(&FIndividual{
			x: []float64{x, y},
			v: []float64{vx, vy},
		})
	}
	positioner := &FPositioner{}
	evaluator := &FEvaluater{}

	sw := evoli.NewSwarm(pop, positioner, .2, .2, evaluator)

	for i := 0; i < 100; i++ {
		err := sw.Next()
		if err != nil {
			panic(err)
		}
	}

	// evaluate the latest population
	for _, v := range sw.Population().Slice() {
		f, err := evaluator.Evaluate(v)
		if err != nil {
			panic(err)
		}
		v.SetFitness(f)
	}

	fmt.Printf("Max Value: %.2f\n", sw.Alpha().Fitness())
}
```

```bash
Max Value: 1.00
```

### Gentic Algorithm

```golang
package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/khezen/evoli"
)

// 3d cosine that gets smaller as you move away from 0,0
func h(x, y float64) float64 {
	d := x*x + y*y
	return math.Cos(d) * (1 / (d/10 + 1))
}

type HIndividual struct {
	v       []float64
	x       []float64
	fitness float64
}

func (i *HIndividual) Equal(other evoli.Individual) bool {
	return i == other
}

func (i *HIndividual) Fitness() float64 {
	return i.fitness
}

func (i *HIndividual) SetFitness(newFitness float64) {
	i.fitness = newFitness
}

type HMutater struct {
}

func (m HMutater) Mutate(indiv evoli.Individual) (evoli.Individual, error) {
	x := rand.Float64()*20 - 10
	y := rand.Float64()*20 - 10
	vx := rand.Float64()*20 - 10
	vy := rand.Float64()*20 - 10
	return &FIndividual{
		x: []float64{x, y},
		v: []float64{vx, vy},
	}, nil
}

type HCrosser struct {
}

func (h HCrosser) Cross(indiv1, indiv2 evoli.Individual) (evoli.Individual, error) {
	fIndiv1, _ := indiv1.(*FIndividual)
	fIndiv2, _ := indiv2.(*FIndividual)
	return &FIndividual{
		x: []float64{(fIndiv1.x[0] + fIndiv2.x[0]) / 2, (fIndiv1.x[1] + fIndiv2.x[1]) / 2},
		v: []float64{(fIndiv1.v[0] + fIndiv2.v[0]) / 2, (fIndiv1.v[1] + fIndiv2.v[1]) / 2},
	}, nil
}

type HEvaluater struct {
}

func (e HEvaluater) Evaluate(indiv evoli.Individual) (Fitness float64, err error) {
	fIndiv, ok := indiv.(*FIndividual)
	if !ok {
		return 0, fmt.Errorf("invalid individual type")
	}
	return f(fIndiv.x[0], fIndiv.x[1]), nil
}

func main() {
	pop := evoli.NewPopulation(50)
	for i := 0; i < pop.Cap(); i++ {
		x := rand.Float64()*20 - 10
		y := rand.Float64()*20 - 10
		vx := rand.Float64()*20 - 10
		vy := rand.Float64()*20 - 10
		pop.Add(&FIndividual{
			x: []float64{x, y},
			v: []float64{vx, vy},
		})
	}
	crosser := HCrosser{}
	mutater := HMutater{}
	evaluator := HEvaluater{}
	mutationProbability := .02
	selecter := evoli.NewTruncationSelecter()
	survivorSize := 30

	ga := evoli.NewGenetic(pop, selecter, survivorSize, crosser, mutater, mutationProbability, evaluator)

	for i := 0; i < 100; i++ {
		err := ga.Next()
		if err != nil {
			panic(err)
		}
	}

	// evaluate the latest population
	for _, v := range ga.Population().Slice() {
		f, err := evaluator.Evaluate(v)
		if err != nil {
			panic(err)
		}
		v.SetFitness(f)
	}

	fmt.Printf("Max Value: %.2f\n", ga.Alpha().Fitness())
}
```

```bash
Max Value: 1.00
```