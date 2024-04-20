package generation

import (
	"fmt"
	"log"
	"math/rand"
)

var (
	maxMutations = 500000
)

// Chromosome представляет собой одно расписание
type Chromosome struct {
	Genes         map[string]*Gene
	Fitness       float64
	CrossoverRate float64
	MutationRate  float64
	Groups        []*Group
	CommonClasses []*CommonClass
	TimeTable     *TimeTable
	Hours, Days   int
}

func (c *Chromosome) Init() {
	c.Genes = make(map[string]*Gene, len(c.Groups))

	for _, group := range c.Groups {
		gene := Gene{
			Hours: c.Hours,
			Days:  c.Days,
		}
		gene.Flags = make([]bool, c.Hours*c.Days)
		gene.Slots = make([]*GeneSlot, 0, c.Hours*c.Days)
		c.Genes[group.ID] = &gene
	}

	// заполняем слоты под потоковые пары
	for _, class := range c.CommonClasses {
		hourCount := 1
		slots := rand.Perm(c.Hours * c.Days)
		for _, value := range slots {
			for _, group := range class.Groups {
				if c.Genes[group.ID].Flags[value] {
					continue
				}
			}
			for _, group := range class.Groups {
				c.Genes[group.ID].Flags[value] = true
				c.Genes[group.ID].Slots =
					append(c.Genes[group.ID].Slots, &GeneSlot{
						Value: value,
						Flag:  true,
					})
			}
			if hourCount < class.Hours {
				hourCount++
			} else {
				break
			}
		}
	}

	// заполняем все остальные слоты
	for _, gene := range c.Genes {
		gene.Fill()
		log.Println("final: ", len(gene.Slots))

	}

	c.Fitness = c.GetFitness()
}

func (c *Chromosome) GetFitness() float64 {
	/*
	   Посчитать фитнес:

	   посчитать кол-во повторяющихся преподов и аудиторий в одно время

	   можно присваивать аудиторию в этом методе (подумай над этим !!)
	   замечание: мне это не нравится

	   выдавать минимально подходящую аудиторию(если такой нет выдаем штраф, если есть удаляем аудиторию из списка)
	*/

	var teacherPoint float64
	var roomPoint float64

	type TimeSlots struct {
		Teachers map[int]int
		Rooms    map[string]int
	}
	timeSlots := make(map[int]TimeSlots)

	// Записываем данные из двумерного массива на лист
	for groupID, gene := range c.Genes {
		for i, slot := range gene.Slots {
			if i >= 50 {
				log.Println(i)
				continue
			}
			if _, ok := c.TimeTable.GroupSlots[groupID]; !ok {
				continue
			}
			class := c.TimeTable.GroupSlots[groupID].Classes[i]

			if _, ok := timeSlots[slot.Value]; !ok {
				timeSlots[slot.Value] = TimeSlots{
					Rooms:    make(map[string]int),
					Teachers: make(map[int]int),
				}
			}

			if class != nil {
				if classID, ok := timeSlots[slot.Value].Teachers[class.Teacher.ID]; ok && classID != class.ID {
					teacherPoint++
				} else {
					timeSlots[slot.Value].Teachers[class.Teacher.ID] = class.ID
				}
				if classID, ok := timeSlots[slot.Value].Rooms[class.Room.ID]; ok && classID != class.ID {
					roomPoint++
				} else {
					timeSlots[slot.Value].Rooms[class.Room.ID] = class.ID
				}
			}
		}
	}

	fmt.Println("teacherPoint: ", teacherPoint)
	fmt.Println("roomPoint: ", roomPoint)

	fitness := 1 - ((teacherPoint + roomPoint) / float64((len(c.Groups)-1.0)*2.0*c.Hours*c.Days))
	return fitness
}

func (c *Chromosome) Compare(chromosome Chromosome) int {
	if c.Fitness == chromosome.Fitness {
		return 0
	} else if c.Fitness > chromosome.Fitness {
		return -1
	}
	return 1
}

// custom mutation
func (c *Chromosome) customMutation() {
	var newFitness float64
	oldFitness := c.GetFitness()
	group := c.Groups[rand.Intn(len(c.Groups))]

	for i := 0; newFitness < oldFitness; i++ {
		gene := *c.Genes[group.ID]

		gene.Clear()
		gene.Fill()

		c.Genes[group.ID] = &gene
		newFitness = c.GetFitness()

		if i >= maxMutations {
			break
		}
	}
}

// Two point crossover
func (c *Chromosome) crossover(father, mother Chromosome) Chromosome {
	group := c.Groups[rand.Intn(len(c.Groups))]

	fatherGene := *father.Genes[group.ID]
	motherGene := *mother.Genes[group.ID]

	oldFatherGene := *father.Genes[group.ID]
	oldMotherGene := *mother.Genes[group.ID]

	fatherGene.Clear()
	motherGene.Clear()

	for _, slot := range oldFatherGene.Slots {
		if !slot.Flag && !motherGene.Flags[slot.Value] {
			motherGene.Slots = append(motherGene.Slots, slot)
			motherGene.Flags[slot.Value] = true
		}
	}

	for _, slot := range oldMotherGene.Slots {
		if !slot.Flag && !fatherGene.Flags[slot.Value] {
			fatherGene.Slots = append(fatherGene.Slots, slot)
			fatherGene.Flags[slot.Value] = true
		}
	}

	fatherGene.Fill()
	motherGene.Fill()

	father.Genes[group.ID] = &fatherGene
	mother.Genes[group.ID] = &motherGene

	if father.GetFitness() > mother.GetFitness() {
		return father
	}
	return mother

}
