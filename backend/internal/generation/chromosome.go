package generation

import (
	"cmp"
	"math/rand"
	"slices"
	"sort"
)

var (
	maxMutations = 100000
)

// Chromosome представляет собой одно расписание
type Chromosome struct {
	Genes         map[string]*Gene
	Fitness       float64
	CrossoverRate float64
	MutationRate  float64
	Groups        []*Group
	CommonClasses []*CommonClass
	TimeTable     TimeTable
	Rooms         []*Room
	Hours, Days   int
	Flag          string
	TeacherPoint  int
	RoomPoint     int
	// mu    sync.Mutex
	// mutex sync.RWMutex
}

func (c *Chromosome) Init() {
	c.Genes = make(map[string]*Gene, len(c.Groups))

	for _, group := range c.Groups {
		gene := &Gene{
			GroupID: group.ID,
			Hours:   c.Hours,
			Days:    c.Days,
		}
		gene.Flags = make([]bool, c.Hours*c.Days)
		gene.Slots = make([]*GeneSlot, 0, c.Hours*c.Days)
		c.Genes[group.ID] = gene
	}

	// заполняем слоты под потоковые пары
	for _, class := range c.CommonClasses {
		hourCount := 1
		slots := rand.Perm(c.Hours * c.Days)
	newTimeSlot:
		for i, value := range slots {
			for _, group := range class.Groups {
				if c.Genes[group.ID].Flags[value] {
					if i+1 == len(slots) {
						// log.Println(class.ID, c.Fitness, "ГРУППЫ !!!")
						c.Fitness = -1 // не получилось
					}
					continue newTimeSlot
				}
			}
			tmp := c.deepCopyGenes(c.Genes)

			for _, group := range class.Groups {
				c.Genes[group.ID].Flags[value] = true
				c.Genes[group.ID].Slots =
					append(c.Genes[group.ID].Slots, &GeneSlot{
						Value: value,
						Flag:  true,
					})
				fitness := c.GetFitness()
				if fitness != 1 {
					if i+1 == len(slots) {
						// log.Println(class.ID, "румы", c.RoomPoint, "преподы", c.TeacherPoint, c.Fitness)

						c.Fitness = -1 // не получилось
					}
					c.Genes = c.deepCopyGenes(tmp)
					continue newTimeSlot
				}
			}
			if hourCount < class.Hours {
				hourCount++
			} else {
				break
			}
		}
	}

	// if c.Fitness != -1 {
	// 	SaveXLSX(*c, "stream")
	// 	log.Println(c.Fitness)
	// }

	// заполняем все остальные слоты
	for _, gene := range c.Genes {
		// log.Println(key, len(gene.Slots))
		gene.Fill()
		// log.Println(key, len(gene.Slots))
		// log.Println(key, len(c.Genes[key].Slots))

		// c.Genes[key] = gene
	}

	if c.Fitness != -1 {
		c.Fitness = c.GetFitness()
	}

	// log.Println()
	// log.Println()

	// log.Println()

	// for key, gene := range c.Genes {
	// 	log.Println(key, len(gene.Slots))
	// }
}

func (c *Chromosome) GetFitness() float64 {
	// c.mutex.RLock()
	// defer c.mutex.RUnlock()
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
		gene.Point = 0
		for i, slot := range gene.Slots {
			if _, ok := c.TimeTable.GroupSlots[groupID]; !ok {
				continue
			}
			class := c.TimeTable.GroupSlots[groupID].Classes[i]

			if _, ok := timeSlots[slot.Value]; !ok {
				timeSlots[slot.Value] = TimeSlots{
					Teachers: make(map[int]int),
				}
			}

			if class.ID != 0 {
				if classID, ok := timeSlots[slot.Value].Teachers[class.Teacher.ID]; ok && classID != class.ID {
					teacherPoint++
					if !slot.Flag {
						gene.Point++
					}
				} else {
					timeSlots[slot.Value].Teachers[class.Teacher.ID] = class.ID
				}
				class.Room = nil
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
				if class.Room == nil {
					if !slot.Flag {
						roomPoint++
						gene.Point++
					}
					class.Room = &Room{
						ID:       "no audience",
						Capacity: 0,
					}
				}
				c.TimeTable.GroupSlots[groupID].Classes[i] = class
			}
		}
	}

	// fmt.Println("teacherPoint: ", teacherPoint)
	// fmt.Println("roomPoint: ", roomPoint)
	c.TeacherPoint = int(teacherPoint)
	c.RoomPoint = int(roomPoint)
	c.Fitness = 1 - ((teacherPoint + roomPoint) / float64((len(c.Groups)-1.0)*2.0*c.Hours*c.Days))

	// c.Fitness = 1 - ((teacherPoint) / float64((len(c.Groups)-1.0)*1.0*c.Hours*c.Days))
	return c.Fitness
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
// func (c *Chromosome) customMutation() Chromosome {
// 	var newFitness float64
// 	oldFitness := c.GetFitness()
// 	fmt.Println("old: ", oldFitness)
// 	group := c.Groups[rand.Intn(len(c.Groups))]

// 	for i := 0; newFitness <= oldFitness; i++ {
// 		gene := *c.Genes[group.ID]

// 		gene.Clear()
// 		gene.Fill()

// 		c.Genes[group.ID] = &gene
// 		newFitness = c.GetFitness()
// 		fmt.Println("new: ", newFitness)

// 		if i >= maxMutations || newFitness == 1.0 {
// 			break
// 		}
// 	}
// 	fmt.Println("last: ", c.GetFitness())

// 	return *c
// }

func (c *Chromosome) customMutation() {
	var newFitness float64
	groups := make([]string, 0, 30)
	b4mutation := c.deepCopy()

	oldFitness := c.GetFitness()
	// fmt.Println("old: ", oldFitness)
	if len(c.Groups) < 30 {
		groups = []string{c.Groups[rand.Intn(len(c.Groups))].ID}
	} else {
		// Преобразуем map в срез для сортировки по значению Point
		geneSlice := make(ByPoint, 0, len(c.Genes))
		for _, gene := range c.Genes {
			geneSlice = append(geneSlice, gene)
		}

		// Сортируем срез по убыванию значения Point
		sort.Sort(geneSlice)

		// Выбираем 30 ключей с наибольшим Point
		for i, gene := range geneSlice {
			if i > 30 || gene.Point == 0 {
				break
			}
			groups = append(groups, gene.GroupID)
		}

		// // Создаем карту для отслеживания уже выбранных индексов
		// selectedIndexes := make(map[int]bool)

		// // Выбор 10 случайных неповторяющихся элементов
		// for i := 0; i < 10; {
		// 	index := rand.Intn(len(c.Groups))
		// 	if !selectedIndexes[index] {
		// 		selectedIndexes[index] = true
		// 		groups = append(groups, c.Groups[index].ID)
		// 		i++
		// 	}
		// }
	}

	for i := 0; newFitness <= oldFitness; i++ {
		for _, groupID := range groups {
			geneCopy := *c.Genes[groupID] // Создаем копию гена

			geneCopy.Clear()
			geneCopy.Fill()

			c.Genes[groupID] = &geneCopy // Сохраняем копию гена обратно в хромосому
		}

		newFitness = c.GetFitness()

		// fmt.Println("new: ", newFitness)

		if i >= maxMutations || newFitness == 1.0 {
			break
		}
	}
	// fmt.Println("last: ", c.GetFitness())

	if oldFitness > newFitness {
		c = b4mutation
		c.GetFitness()
	}
}

// Two point crossover
func (c *Chromosome) crossover(father, mother Chromosome) Chromosome {
	groups := make([]string, 0, 30)

	if len(c.Groups) < 30 {
		groups = []string{c.Groups[rand.Intn(len(c.Groups))].ID}
	} else {
		// Создаем карту для отслеживания уже выбранных индексов
		selectedIndexes := make(map[int]bool)

		// Выбор 30 случайных неповторяющихся элементов
		for i := 0; i < 30; {
			index := rand.Intn(len(c.Groups))
			if !selectedIndexes[index] {
				selectedIndexes[index] = true
				groups = append(groups, c.Groups[index].ID)
				i++
			}
		}
	}

	for _, groupID := range groups {

		fatherGene := *father.Genes[groupID]
		motherGene := *mother.Genes[groupID]

		oldFatherGene := *father.Genes[groupID]
		oldMotherGene := *mother.Genes[groupID]

		fatherGene.Clear()
		motherGene.Clear()

		// раскидываем непотоковые пары
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

		// заполняем оставшиеся
		fatherGene.Fill()
		motherGene.Fill()

		father.Genes[groupID] = &fatherGene
		mother.Genes[groupID] = &motherGene
		// c.mutex.RUnlock()
	}
	if father.GetFitness() > mother.GetFitness() {
		return father
	}
	return mother

}

func (c *Chromosome) deepCopy() *Chromosome {
	// Создаем новую переменную для хранения копии хромосомы
	newChromosome := &Chromosome{
		Genes:         make(map[string]*Gene),
		Fitness:       c.Fitness,
		CrossoverRate: c.CrossoverRate,
		MutationRate:  c.MutationRate,
		Groups:        make([]*Group, len(c.Groups)),
		CommonClasses: make([]*CommonClass, len(c.CommonClasses)),
		TimeTable:     c.TimeTable,
		Rooms:         make([]*Room, len(c.Rooms)),
		Hours:         c.Hours,
		Days:          c.Days,
		Flag:          c.Flag,
		TeacherPoint:  c.TeacherPoint,
		RoomPoint:     c.RoomPoint,
	}

	// Копируем информацию о генах
	for key, gene := range c.Genes {
		newGene := &Gene{
			GroupID: gene.GroupID,
			Slots:   make([]*GeneSlot, len(gene.Slots)),
			Flags:   make([]bool, len(gene.Flags)),
			Hours:   gene.Hours,
			Days:    gene.Days,
			Point:   gene.Point,
		}

		// Копируем слоты гена
		for i, slot := range gene.Slots {
			newGene.Slots[i] = &GeneSlot{
				Value: slot.Value,
				Flag:  slot.Flag,
			}
		}
		// Копируем флаги гена
		copy(newGene.Flags, gene.Flags)

		newChromosome.Genes[key] = newGene
	}

	// Копируем информацию о группах
	copy(newChromosome.Groups, c.Groups)

	// Копируем информацию о общих классах
	copy(newChromosome.CommonClasses, c.CommonClasses)

	// Копируем информацию о комнатах
	copy(newChromosome.Rooms, c.Rooms)

	return newChromosome
}

func (c *Chromosome) deepCopyGenes(genes map[string]*Gene) map[string]*Gene {
	tmp := make(map[string]*Gene, len(genes))

	for groupID, gene := range genes {
		localGene := *gene
		localGene.Flags = make([]bool, c.Hours*c.Days)
		localGene.Slots = make([]*GeneSlot, 0, c.Hours*c.Days)
		for _, slot := range gene.Slots {
			localSlot := *slot

			localGene.Slots = append(localGene.Slots, &localSlot)
		}
		copy(localGene.Flags, gene.Flags)
		tmp[groupID] = &localGene
	}
	return tmp
}
