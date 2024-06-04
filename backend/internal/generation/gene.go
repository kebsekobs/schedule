package generation

import (
	"math/rand"
)

// Gene представляет собой последовательность выбора слотов для одной группы
type Gene struct {
	GroupID     string
	Slots       []*GeneSlot
	Flags       []bool // true == занято
	Hours, Days int
	Point       int
}

type ByPoint []*Gene

func (a ByPoint) Len() int           { return len(a) }
func (a ByPoint) Less(i, j int) bool { return a[i].Point > a[j].Point }
func (a ByPoint) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type GeneSlot struct {
	Value int
	Flag  bool // true == потоковая пара
}

func (g *Gene) Fill() {
	/*
		Потоковые пары раскидываем перед не потоковыми

		Так как у разных хромосом потоковые пары
		будут раскиданы случайно(скорее всего будут не совпадать)

		Нужно добавить какой-нибудь флаг для потоковых пар

		При этом флаге потоковые пары будут браться
		от одного из родителей

		Это позволит при кроссовере гарантировать
		что потоковые пары будут в одно время

		При мутации этот флаг будет показывать какие слоты
		двигать не надо
	*/

	// генерим случайную последовательность пар для каждой группы
	slots := rand.Perm(g.Hours * g.Days)
	// проходимся по всем слотам
	for _, value := range slots {
		// проверяем занят ли слот
		if !g.Flags[value] {
			// зажигаем флаг
			g.Flags[value] = true
			// добавляем слот в список
			g.Slots = append(g.Slots, &GeneSlot{
				Value: value,
				Flag:  false,
			})

		}
	}

}

func (g *Gene) Clear() {
	var clearSlots []*GeneSlot
	for _, slot := range g.Slots {
		if !slot.Flag {
			g.Flags[slot.Value] = false
		} else {
			clearSlots = append(clearSlots, slot)
		}
	}
	g.Slots = clearSlots
}
