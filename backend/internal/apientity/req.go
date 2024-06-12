package api

type Group struct {
	ID         int    `json:"id"`
	Magistracy bool   `json:"magistracy"`
	GroupID    int    `json:"groupId"`
	Capacity   string `json:"capacity"`
	Name       string `json:"name"`
}

type Teacher struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Initials string `json:"initials"`
}

type Classroom struct {
	ID          string `json:"id"`
	ClassroomID string `json:"classroomId"`
	Capacity    string `json:"capacity"`
}

type Discipline struct {
	ID              int      `json:"id"`
	DisciplinesId   int      `json:"disciplinesId"`
	Name            string   `json:"name"`
	Teachers        string   `json:"teachers"`
	Hours           string   `json:"hours"`
	RelatedGroupsId []string `json:"relatedGroupsId"`
}
