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
	group := Group{
		ID: c.Groups[i].ID,
	}
	for _, g := range c.Groups {
		group.Quantity += g.Quantity
	}
	return &Class{
		ID:      c.ID,
		Teacher: c.Teacher,
		Room:    c.Room,
		Group:   &group,
		Name:    c.Name,
		Hours:   c.Hours,
	}
}
