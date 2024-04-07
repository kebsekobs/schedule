package generation

import (
	"log"
)

type TimeTable struct {
	// Список пар для каждой группы
	// Ключ это id группы
	GroupSlots map[string]*Slot

	Classes       []*Class       // Список пар(без потоковых)
	Groups        []*Group       // Список групп
	CommonClasses []*CommonClass // Список потоковых пар
}

func (t *TimeTable) Init(hours, days int) {

	// выделяем память по кол-ву групп
	t.GroupSlots = make(map[string]*Slot, len(t.Groups))

	for _, class := range t.CommonClasses {
		for i, group := range class.Groups {
			t.fillSlots(group, class.makeClass(i), hours, days)
		}

	}

	for _, class := range t.Classes {
		t.fillSlots(class.Group, class, hours, days)
	}
}

func (t *TimeTable) fillSlots(group *Group, class *Class, hours, days int) {
	hourCount := 1

	if slots, ok := t.GroupSlots[group.ID]; !ok {
		t.GroupSlots[group.ID] = &Slot{
			Classes: make([]*Class, hours*days),
		}
	} else if len(slots.Classes) == 0 {
		t.GroupSlots[group.ID].Classes = make([]*Class, days*hours)
	}

	// suppose java has to be taught for 5 hours then we make 5
	// slots for java, we keep track through hourcount
	for {
		if hours*days <= t.GroupSlots[group.ID].I {
			log.Printf("У группы %v не осталось свободных слотов для пар", group.ID)
			break
		}
		t.GroupSlots[group.ID].Classes[t.GroupSlots[group.ID].I] = class
		t.GroupSlots[group.ID].I++
		if hourCount < class.Hours {
			hourCount++
		} else {
			break
		}
	}
}
