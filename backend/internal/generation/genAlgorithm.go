package generation

type GenAlgorithm struct {
	AverageFitness        int
	PastAverageFitness    int
	Running               bool
	Chromosomes           []Chromosome
	Data                  map[string]any
	StayInRoomAssignments map[string]any
	TournamentSize        float64
	ElitePercent          float64
	MutationRate          float64
	LowVariety            int
	HighestFitness        int
	LowestFitness         int
}
