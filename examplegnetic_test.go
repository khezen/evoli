package evoli_test

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
	return &HIndividual{
		x: []float64{x, y},
		v: []float64{vx, vy},
	}, nil
}

type HCrosser struct {
}

func (h HCrosser) Cross(parent1, parent2 evoli.Individual) (child1, child2 evoli.Individual, err error) {
	fIndiv1, _ := parent1.(*HIndividual)
	fIndiv2, _ := parent2.(*HIndividual)
	w := 0.1 + 0.8*rand.Float64()
	return &HIndividual{
			x: []float64{w*fIndiv1.x[0] + (1-w)*fIndiv2.x[0], w*fIndiv1.x[1] + (1-w)*fIndiv2.x[1]},
			v: []float64{w*fIndiv1.v[0] + (1-w)*fIndiv2.v[0], w*fIndiv1.v[1] + (1-w)*fIndiv2.v[1]},
		}, &HIndividual{
			x: []float64{(1-w)*fIndiv1.x[0] + w*fIndiv2.x[0], (1-w)*fIndiv1.x[1] + w*fIndiv2.x[1]},
			v: []float64{(1-w)*fIndiv1.v[0] + w*fIndiv2.v[0], (1-w)*fIndiv1.v[1] + w*fIndiv2.v[1]},
		}, nil
}

type HEvaluater struct {
}

func (e HEvaluater) Evaluate(indiv evoli.Individual) (Fitness float64, err error) {
	fIndiv, ok := indiv.(*HIndividual)
	if !ok {
		return 0, fmt.Errorf("invalid individual type")
	}
	return h(fIndiv.x[0], fIndiv.x[1]), nil
}

func ExampleNewGenetic() {

	pop := evoli.NewPopulation(50)
	for i := 0; i < pop.Cap(); i++ {
		x := rand.Float64()*20 - 10
		y := rand.Float64()*20 - 10
		vx := rand.Float64()*20 - 10
		vy := rand.Float64()*20 - 10
		pop.Add(&HIndividual{
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

	// Output: Max Value: 1.00
}
