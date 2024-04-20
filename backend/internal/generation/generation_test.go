package generation_test

import (
	"fmt"
	"testing"

	"github.com/kebsekobs/schedule/backend/internal/generation"
)

func TestGeneration(t *testing.T) {

	test := struct {
		gnrt generation.Generation
	}{}

	test.gnrt.Days = 10
	test.gnrt.Hours = 5

	test.gnrt.CommonClasses = []*generation.CommonClass{
		{
			ID: 1,
			Teacher: &generation.Teacher{
				ID:   2,
				Name: "b",
			},
			Room: &generation.Room{
				ID: "1",
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
			Name:  "test1",
			Hours: 2,
		},
		{
			ID: 4,
			Teacher: &generation.Teacher{
				ID:   1,
				Name: "a",
			},
			Room: &generation.Room{
				ID: "1",
			},
			Groups: []*generation.Group{
				{
					ID: "2",
				},
				{
					ID: "3",
				},
			},
			Name:  "test2",
			Hours: 5,
		},
	}

	test.gnrt.Groups = []*generation.Group{
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
	test.gnrt.Classes = []*generation.Class{
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
			Hours: 10,
		},
		{
			ID: 2,
			Teacher: &generation.Teacher{
				ID:   1,
				Name: "a",
			},
			Room: &generation.Room{
				ID: "1",
			},
			Group: &generation.Group{
				ID: "2",
			},
			Name:  "zxcv",
			Hours: 10,
		},
	}

	test.gnrt.StartGeneration()
	t.Run("letsgo", func(t *testing.T) {
		// for group, value := range test.gnrt.Genes {
		// 	fmt.Println(group)
		// 	fmt.Println()

		// 	for key, i := range value.Slots {
		// 		fmt.Println(key)
		// 		fmt.Println(i)
		// 	}
		// 	fmt.Println()
		// 	fmt.Println()

		// 	fmt.Println()

		// }
		fmt.Println(test.gnrt.FinalSon)
	})
}
