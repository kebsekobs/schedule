package generation

import (
	"fmt"
	"math/rand"
)

var (
	crossoverRate = 1.0
	mutationRate  = 0.1
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
}

func (c *Chromosome) Init(hours, days int) {
	c.Genes = make(map[string]*Gene, len(c.Groups))

	for _, group := range c.Groups {
		gene := Gene{}
		gene.Flags = make([]bool, hours*days)
		gene.Slots = make([]*GeneSlot, 0, hours*days)
		c.Genes[group.ID] = &gene
	}

	// заполняем слоты под потоковые пары
	for _, class := range c.CommonClasses {
		hourCount := 1
		slots := rand.Perm(hours * days)
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
		gene.Fill(hours, days)
	}

	c.Fitness = c.GetFitness(hours, days)
}

func (c *Chromosome) GetFitness(hours, days int) float64 {
	/*
	   Посчитать фитнес:

	   посчитать кол-во повторяющихся преподов и аудиторий в одно время

	   можно присваивать аудиторию в этом методе (подумай над этим !!)
	   замечание: мне это не нравится
	*/

	var teacherPoint float64
	var roomPoint float64

	for i := 0; i < hours*days; i++ {

		teachers := make(map[int]int)
		rooms := make(map[string]int)
		fmt.Println(c.TimeTable.GroupSlots)
		for _, group := range c.Groups {
			if _, ok := c.TimeTable.GroupSlots[group.ID]; !ok {
				continue
			}
			slot := c.TimeTable.GroupSlots[group.ID].Classes[i]

			if slot != nil {
				if classID, ok := teachers[slot.Teacher.ID]; ok && classID != slot.ID {
					teacherPoint++
				} else {
					teachers[slot.Teacher.ID] = slot.ID
				}
				if classID, ok := rooms[slot.Room.ID]; ok && classID != slot.ID {
					roomPoint++
				} else {
					rooms[slot.Room.ID] = slot.ID
				}
			}
		}

	}
	fmt.Println(teacherPoint)
	fmt.Println(roomPoint)

	fitness := 1 - ((teacherPoint + roomPoint) / float64((len(c.Groups)-1.0)*2*hours*days))
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
