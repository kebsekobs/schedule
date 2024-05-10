package generation

import (
	"cmp"
	"fmt"
	"log"
	"math/rand"
	"slices"
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
	Rooms         []*Room
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
	newTimeSlot:
		for _, value := range slots {
			for _, group := range class.Groups {
				if c.Genes[group.ID].Flags[value] {
					continue newTimeSlot
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
		if len(gene.Slots) > 50 {
			for _, el := range gene.Slots {
				fmt.Println(el.Value)
			}
		}
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
		Rooms    []SlotRoom
	}
	timeSlots := make(map[int]TimeSlots)

	slices.SortFunc(c.Rooms, func(i, j *Room) int {
		return cmp.Compare(i.Capacity, j.Capacity)
	})

	for i := 0; i < (c.Hours * c.Days); i++ {
		rooms := make([]SlotRoom, 0, len(c.Rooms))
		for _, room := range c.Rooms {
			rooms = append(rooms, SlotRoom{
				Room: room,
			})
		}
		timeSlots[i] = TimeSlots{
			Teachers: make(map[int]int),
			Rooms:    rooms,
		}
	}

	for groupID, gene := range c.Genes {
		for i, slot := range gene.Slots {
			if i >= 50 {
				continue
			}
			if _, ok := c.TimeTable.GroupSlots[groupID]; !ok {
				continue
			}
			class := c.TimeTable.GroupSlots[groupID].Classes[i]

			if _, ok := timeSlots[slot.Value]; !ok {
				timeSlots[slot.Value] = TimeSlots{
					Teachers: make(map[int]int),
				}
			}

			if class != nil {
				if classID, ok := timeSlots[slot.Value].Teachers[class.Teacher.ID]; ok && classID != class.ID {
					teacherPoint++
				} else {
					timeSlots[slot.Value].Teachers[class.Teacher.ID] = class.ID
				}

				if class.Room == nil {
					for j, room := range timeSlots[slot.Value].Rooms {
						if room.Room.Capacity < class.Group.Quantity {
							continue
						}

						if room.ClassID == class.ID {
							class.Room = room.Room
							break
						}
						if room.ClassID == 0 && room.Room.Capacity >= class.Group.Quantity {
							class.Room = room.Room
							room.ClassID = class.ID
							timeSlot := timeSlots[slot.Value]
							timeSlot.Rooms[j] = room
							timeSlots[slot.Value] = timeSlot

							break
						}
					}
				}
				if class.Room == nil {
					roomPoint++
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
