package generation_test

import (
	"fmt"
	"testing"

	"github.com/kebsekobs/schedule/backend/internal/generation"
)

func TestTimeTable(t *testing.T) {

	test := struct {
		timeTable generation.TimeTable
	}{}

	test.timeTable.Classes = []*generation.Class{
		{
			ID: 1,
			Teacher: &generation.Teacher{
				ID:   1,
				Name: "a",
			},
			Room: &generation.Room{
				ID: "1",
			},
			Group: &generation.Group{
				ID: "1",
			},
			Name:  "qwer",
			Hours: 5,
		},
		{
			ID: 3,
			Teacher: &generation.Teacher{
				ID:   3,
				Name: "c",
			},
			Room: &generation.Room{
				ID: "3",
			},
			Group: &generation.Group{
				ID: "2",
			},
			Name:  "zxcv",
			Hours: 2,
		},
	}

	test.timeTable.CommonClasses = []*generation.CommonClass{
		{
			ID: 2,
			Teacher: &generation.Teacher{
				ID:   2,
				Name: "b",
			},
			Room: &generation.Room{
				ID: "2",
			},
			Groups: []*generation.Group{
				{
					ID: "1",
				},
				{
					ID: "2",
				},
				{
					ID: "3",
				},
			},
			Name:  "asdf",
			Hours: 3,
		},
		{
			ID: 4,
			Teacher: &generation.Teacher{
				ID:   2,
				Name: "a",
			},
			Room: &generation.Room{
				ID: "3",
			},
			Groups: []*generation.Group{
				{
					ID: "2",
				},
				{
					ID: "3",
				},
			},
			Name:  "asdf",
			Hours: 2,
		},
	}

	test.timeTable.Groups = []*generation.Group{
		{
			ID: "1",
		},
		{
			ID: "2",
		},
		{
			ID: "3",
		},
	}

	test.timeTable.Init(5, 2)
	t.Run("letsgo", func(t *testing.T) {
		for group, value := range test.timeTable.GroupSlots {
			fmt.Println(group)

			for _, i := range value.Classes {
				fmt.Println(i)
			}
			fmt.Println()
			fmt.Println()

			fmt.Println()

		}
	})
}
