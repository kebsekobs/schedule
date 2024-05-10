package generation

type Room struct {
	ID       string
	Capacity int // Вместимость аудитории
}

type SlotRoom struct {
	Room    *Room
	ClassID int
}
