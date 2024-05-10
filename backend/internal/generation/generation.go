package generation

import (
	"log"
	"math/rand"
	"slices"
)

var (
	populationSize = 1000
	maxGenerations = 100
	crossoverRate  = 0.5
	mutationRate   = 0.1
)

type Generation struct {
	Groups                 []*Group
	CommonClasses          []*CommonClass
	Classes                []*Class
	Rooms                  []*Room
	firstGeneration        []*Chromosome
	newGeneration          []*Chromosome
	TimeTable              *TimeTable
	firstGenerationFitness float64
	newGenerationFitness   float64
	FinalSon               *Chromosome
	Hours, Days            int
}

func (g *Generation) StartGeneration() {
	g.TimeTable = &TimeTable{
		Groups:        g.Groups,
		Classes:       g.Classes,
		CommonClasses: g.CommonClasses,
		Hours:         g.Hours,
		Days:          g.Days,
	}

	g.TimeTable.Init()

	g.initPopulation()

	g.createNewGeneration()

}

func (g *Generation) initPopulation() {
	g.firstGeneration = make([]*Chromosome, 0, populationSize)

	for i := 0; i < populationSize; i++ {
		chromosome := Chromosome{
			CommonClasses: g.CommonClasses,
			Groups:        g.Groups,
			TimeTable:     g.TimeTable,
			Hours:         g.Hours,
			Days:          g.Days,
			Rooms:         g.Rooms,
		}

		chromosome.Init()

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

func (g *Generation) createNewGeneration() {
	var (
		father, mother     Chromosome
		numberOfGeneration int
	)

	son := &Chromosome{
		CommonClasses: g.CommonClasses,
		Groups:        g.Groups,
		TimeTable:     g.TimeTable,
		Hours:         g.Hours,
		Days:          g.Days,
		Rooms:         g.Rooms,
	}

	for numberOfGeneration < maxGenerations {
		log.Println("number of generation: ", numberOfGeneration)
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

		for ; count < populationSize; count++ {
			father = *g.selectParentRoulette()
			mother = *g.selectParentRoulette()

			// crossover
			if rand.Float64() < crossoverRate {
				log.Println("crossover")
				temp := son.crossover(father, mother)
				son = &temp
			} else {
				log.Println("no crossover")

				son = &father
			}

			// mutation
			son.customMutation()
			log.Println("fitnesses = ", son.Fitness)
			if son.Fitness == 1.0 {
				break
			}

			g.newGeneration = append(g.newGeneration, son)
			g.newGenerationFitness += son.GetFitness()
		}

		//if chromosome with fitness 1 found
		if count < populationSize {

			g.FinalSon = son
			break
		}

		//if chromosome with required fitness not found in this generation
		g.firstGeneration = g.newGeneration
		slices.SortFunc(
			g.firstGeneration,
			func(a, b *Chromosome) int {
				return a.Compare(*b)
			},
		)
		numberOfGeneration++
	}
	SaveXLSX(*g.FinalSon)
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
