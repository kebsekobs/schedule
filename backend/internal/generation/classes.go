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

type ByHours []*CommonClass

func (a ByHours) Len() int           { return len(a) }
func (a ByHours) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByHours) Less(i, j int) bool { return a[i].Hours > a[j].Hours }

type ByGroupQuantitySum []*CommonClass

func (ccs ByGroupQuantitySum) Len() int {
	return len(ccs)
}

func (ccs ByGroupQuantitySum) Less(i, j int) bool {
	sumI, sumJ := 0, 0
	for _, group := range ccs[i].Groups {
		sumI += group.Quantity
	}
	for _, group := range ccs[j].Groups {
		sumJ += group.Quantity
	}
	return sumI > sumJ
}

func (ccs ByGroupQuantitySum) Swap(i, j int) {
	ccs[i], ccs[j] = ccs[j], ccs[i]
}

type ClassWithHours interface {
	GetHours() int
}

func FilterAnyClassesByHours(classes []ClassWithHours) []ClassWithHours {
	filteredClasses := []ClassWithHours{}
	for _, c := range classes {
		if c.GetHours() != 0 {
			filteredClasses = append(filteredClasses, c)
		}
	}
	return filteredClasses
}

func (c Class) GetHours() int {
	return c.Hours
}

func (c CommonClass) GetHours() int {
	return c.Hours
}

func FilterCommonClassesByHours(commonClasses []*CommonClass) []*CommonClass {
	filteredCommonClasses := []*CommonClass{}
	for _, c := range commonClasses {
		if c.Hours != 0 {
			filteredCommonClasses = append(filteredCommonClasses, c)
		}
	}
	return filteredCommonClasses
}

func FilterClassesByHours(classes []Class) []*Class {
	filteredClasses := []*Class{}
	for _, c := range classes {
		if c.Hours != 0 {
			filteredClasses = append(filteredClasses, &c)
		}
	}
	return filteredClasses
}
