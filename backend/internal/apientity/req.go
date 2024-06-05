package api

type Group struct {
	ID         string `json:"id"`
	Magistracy bool   `json:"magistracy"`
	GroupID    string `json:"groupId"`
	Name       string `json:"name"`
	Capacity   string `json:"capacity"`
}

type Teacher struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Initials string `json:"initials"`
}

type Classroom struct {
	ID          string `json:"id"`
	ClassroomID string `json:"classroomId"`
	Capacity    string `json:"capacity"`
}

type Discipline struct {
	ID              string   `json:"id"`
	DisciplinesId   string   `json:"disciplinesId"`
	Name            string   `json:"name"`
	Teachers        string   `json:"teachers"`
	Hours           int      `json:"hours"`
	RelatedGroupsId []string `json:"relatedGroupsId"`
}
