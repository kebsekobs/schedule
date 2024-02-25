package entity

type Class struct {
	Teacher *Teacher
	Room    *Room
	Groups  []*Group
	Name    string
	Hours   int
}
