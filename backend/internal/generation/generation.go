package generation

import (
	"fmt"
	"log"
	"math/rand"
	"slices"
	"strconv"
)

var (
	populationSize = 1000
	maxGenerations = 500
	crossoverRate  = 0.5
	mutationRate   = 0.1
)

type Generation struct {
	Groups                 []*Group
	CommonClasses          []*CommonClass
	Classes                []*Class
	Rooms                  []*Room
	firstGeneration        []Chromosome
	newGeneration          []Chromosome
	TimeTable              TimeTable
	firstGenerationFitness float64
	newGenerationFitness   float64
	FinalSon               *Chromosome
	Hours, Days            int
}

func (g *Generation) StartGeneration() {
	g.TimeTable = TimeTable{
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
	g.firstGeneration = make([]Chromosome, 0, populationSize)

	for i := 0; i < populationSize; i++ {
		chromosome := Chromosome{
			CommonClasses: g.CommonClasses,
			Groups:        g.Groups,
			TimeTable:     g.TimeTable,
			Hours:         g.Hours,
			Days:          g.Days,
			Rooms:         g.Rooms,
		}
		for j := 0; ; j++ {
			chromosome.Init()
			if chromosome.Fitness != -1 {
				break
			}
			log.Println("try №", j)
		}

		log.Println("chromosome №", i, "fitness", chromosome.Fitness)

		g.firstGeneration = append(g.firstGeneration, chromosome)

		g.firstGenerationFitness += chromosome.Fitness
	}

	slices.SortFunc(
		g.firstGeneration,
		func(a, b Chromosome) int {
			return a.Compare(b)
		},
	)
}

func (g *Generation) createNewGeneration() {
	var (
		father, mother     Chromosome
		numberOfGeneration int
	)

	son := Chromosome{
		CommonClasses: g.CommonClasses,
		Groups:        g.Groups,
		TimeTable:     g.TimeTable,
		Hours:         g.Hours,
		Days:          g.Days,
		Rooms:         g.Rooms,
	}

	for numberOfGeneration < maxGenerations {
		log.Println("number of generation: ", numberOfGeneration)
		g.newGeneration = make([]Chromosome, 0, populationSize)
		count := 0
		var percent float64
		log.Println("teacher point = ", g.firstGeneration[count].TeacherPoint)
		log.Println("room point = ", g.firstGeneration[count].RoomPoint)
		SaveXLSX(g.firstGeneration[count], strconv.Itoa(numberOfGeneration))

		// Элитизм
		for ; count < populationSize/10; count++ {
			g.newGeneration = append(
				g.newGeneration,
				g.firstGeneration[count],
			)
			g.newGenerationFitness += g.firstGeneration[count].Fitness
			log.Println("top 5 fitnesses = ", g.firstGeneration[count].Fitness)
		}

		for ; count < populationSize; count++ {
			father = g.selectParentRoulette(20.0)
			// log.Println("father fitnesses = ", father.Fitness)
			mother = g.selectParentRoulette(10.0)
			// log.Println("mother fitnesses = ", mother.Fitness)

			// crossover
			if rand.Float64() < crossoverRate {
				// log.Println("crossover")
				temp := son.crossover(father, mother)
				son = temp
			} else {
				// log.Println("no crossover")

				son = father
			}
			// log.Println("fitnesses b4 = ", son.Fitness)

			// mutation
			// linkSon := &son
			son.customMutation()

			// son = *linkSon
			// log.Println("fitnesses = ", son.Fitness)
			if son.Fitness == 1.0 {
				break
			}

			g.newGeneration = append(g.newGeneration, son)
			g.newGenerationFitness += son.GetFitness()
			showPercent(&populationSize, &count, &percent, 0.1, 0)
		}

		//if chromosome with fitness 1 found
		if count < populationSize {

			g.FinalSon = &son
			break
		}

		//if chromosome with required fitness not found in this generation
		g.firstGeneration = g.newGeneration
		slices.SortFunc(
			g.firstGeneration,
			func(a, b Chromosome) int {
				return a.Compare(b)
			},
		)
		numberOfGeneration++
	}
	SaveXLSX(*g.FinalSon, "final")
}

// func (g *Generation) createNewGeneration() {
// 	var (
// 		// father, mother     Chromosome
// 		numberOfGeneration int
// 	)

// 	for numberOfGeneration < maxGenerations {
// 		log.Println("number of generation: ", numberOfGeneration)
// 		g.newGeneration = make([]Chromosome, 0, populationSize)
// 		count := 0
// 		var percent float64
// 		log.Println("teacher point = ", g.firstGeneration[count].TeacherPoint)
// 		log.Println("room point = ", g.firstGeneration[count].RoomPoint)

// 		// Элитизм
// 		for ; count < populationSize/10; count++ {
// 			g.newGeneration = append(
// 				g.newGeneration,
// 				g.firstGeneration[count],
// 			)
// 			g.newGenerationFitness += g.firstGeneration[count].Fitness
// 			log.Println("top 10 fitnesses = ", g.firstGeneration[count].Fitness)
// 		}

// 		foundPerfectSon := false

// 		var wg sync.WaitGroup
// 		numThreads := 10
// 		elementsPerThread := (populationSize - count) / numThreads

// 		for i := 0; i < numThreads; i++ {
// 			wg.Add(1)
// 			go func(startIndex, endIndex, goIndex int) {
// 				defer wg.Done()
// 				for j := startIndex; j < endIndex; j++ {
// 					father := g.selectParentRoulette()
// 					mother := g.selectParentRoulette()
// 					son := Chromosome{
// 						CommonClasses: g.CommonClasses,
// 						Groups:        g.Groups,
// 						TimeTable:     g.TimeTable,
// 						Hours:         g.Hours,
// 						Days:          g.Days,
// 						Rooms:         g.Rooms,
// 					}

// 					if rand.Float64() < crossoverRate {
// 						temp := son.crossover(father, mother)
// 						son = temp
// 					} else {
// 						son = father
// 					}

// 					// mutation
// 					son.customMutation()

// 					if son.Fitness == 1.0 {
// 						foundPerfectSon = true
// 						g.FinalSon = &son
// 						SaveXLSX(*g.FinalSon, "final")
// 						break
// 					}

// 					g.newGeneration = append(g.newGeneration, son)
// 					g.newGenerationFitness += son.GetFitness()
// 					showPercent(&endIndex, &startIndex, &percent, 0.05, goIndex)
// 				}
// 			}(0, elementsPerThread, i)
// 		}

// 		wg.Wait()

// 		if foundPerfectSon {
// 			return
// 		}

// 		g.firstGeneration = g.newGeneration
// 		slices.SortFunc(g.firstGeneration, func(a, b Chromosome) int {
// 			return a.Compare(b)
// 		})

// 		numberOfGeneration++
// 	}
// }

func (g *Generation) selectParentRoulette(percent float64) Chromosome {
	fgf := g.firstGenerationFitness / percent
	var (
		currentSum float64
		count      int
	)
	rndFloat := rand.Float64() * fgf

	for currentSum < rndFloat {
		currentSum += g.firstGeneration[count].Fitness
		count++
	}
	return g.firstGeneration[count-1]
}

func showPercent(countall, cur *int, percent *float64, step_percent float64, i int) float64 {
	if float64(*cur)/float64(*countall) > *percent+step_percent {
		*percent = 0
		for *percent < float64(*cur)/float64(*countall) {
			*percent += step_percent
		}
		*percent -= step_percent
		log.Println("goroutine №", i, "percent =  ", fmt.Sprintf("%.2f", *percent*100), "%, current = ", *cur)
	}
	return *percent

}
