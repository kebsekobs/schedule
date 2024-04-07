package generation

import (
	"math/rand"
	"slices"
)

var (
	populationSize = 1000
	maxGenerations = 100
)

type Generation struct {
	Groups                 []*Group
	CommonClasses          []*CommonClass
	Classes                []*Class
	Rooms                  []*Room
	firstGeneration        []*Chromosome
	newGeneration          []*Chromosome
	timeTable              *TimeTable
	firstGenerationFitness float64
	newGenerationFitness   float64
}

func (g *Generation) StartGeneration(hours, days int) {
	g.timeTable = &TimeTable{
		Groups:        g.Groups,
		Classes:       g.Classes,
		CommonClasses: g.CommonClasses,
	}
	g.timeTable.Init(hours, days)

	g.initPopulation(hours, days)

	g.createNewGeneration(hours, days)

}

func (g *Generation) initPopulation(hours, days int) {
	g.firstGeneration = make([]*Chromosome, 0, populationSize)

	for i := 0; i < populationSize; i++ {
		chromosome := Chromosome{
			CommonClasses: g.CommonClasses,
			Groups:        g.Groups,
			TimeTable:     g.timeTable,
		}

		chromosome.Init(hours, days)

		g.firstGeneration = append(g.firstGeneration, &chromosome)

		g.firstGenerationFitness += chromosome.Fitness
	}

	slices.SortFunc(
		g.firstGeneration,
		func(a, b *Chromosome) int {
			return a.Compare(*b)
		},
	)
}

func (g *Generation) createNewGeneration(hours, days int) {
	var (
		father, mother, son Chromosome
		numberOfGeneration  int
	)

	for numberOfGeneration < maxGenerations {
		g.newGeneration = make([]*Chromosome, 0, populationSize)
		count := 0

		// Элитизм
		for ; count < populationSize/10; count++ {
			g.newGeneration = append(
				g.newGeneration,
				g.firstGeneration[count],
			)
			g.newGenerationFitness += g.firstGeneration[count].Fitness
		}

		for count < populationSize {
			father = *g.selectParentRoulette()
			mother = *g.selectParentRoulette()

		}
	}
}

func (g *Generation) selectParentRoulette() *Chromosome {
	g.firstGenerationFitness /= 10.0
	var (
		currentSum float64
		count      int
	)
	rndFloat := rand.Float64() * g.firstGenerationFitness

	for currentSum < rndFloat {
		currentSum += g.firstGeneration[count].Fitness
		count++
	}
	return g.firstGeneration[count]
}
