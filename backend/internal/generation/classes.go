package generation

type Class struct {
	ID      int
	Teacher *Teacher
	Room    *Room
	Group   *Group
	Name    string
	Hours   int
}

type CommonClass struct {
	ID      int
	Teacher *Teacher
	Room    *Room
	Groups  []*Group
	Name    string
	Hours   int
}

func (c *CommonClass) makeClass(i int) *Class {
	return &Class{
		ID:      c.ID,
		Teacher: c.Teacher,
		Room:    c.Room,
		Group:   c.Groups[i],
		Name:    c.Name,
		Hours:   c.Hours,
	}
}
