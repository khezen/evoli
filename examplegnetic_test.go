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

func ExampleNewGenetic() {

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
			panic(err.Error())
		}
	}

	// evaluate the latest population
	for _, v := range ga.Population().Slice() {
		f, err := evaluator.Evaluate(v)
		if err != nil {
			panic(err.Error())
		}
		v.SetFitness(f)
	}

	fmt.Printf("Max Value: %.2f\n", ga.Alpha().Fitness())

	// Output: Max Value: 1.00
}
