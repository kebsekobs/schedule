package entity

type Chromosome struct {
	Gene          []*Gene
	Fitness       float64
	Point         int
	CrossoverRate float64
	MutationRate  float64
	Hours         int
	Days int
}
