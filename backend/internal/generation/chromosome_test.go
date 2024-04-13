package generation_test

import (
	"fmt"
	"testing"

	"github.com/kebsekobs/schedule/backend/internal/generation"
)

func TestInit(t *testing.T) {

	test := struct {
		chromosome generation.Chromosome
	}{}

	test.chromosome.CommonClasses = []*generation.CommonClass{
		// {
		// 	ID: 1,
		// 	Teacher: &generation.Teacher{
		// 		ID:   2,
		// 		Name: "b",
		// 	},
		// 	Room: &generation.Room{
		// 		ID: "1",
		// 	},
		// 	Groups: []*generation.Group{
		// 		{
		// 			ID: "1",
		// 		},
		// 		{
		// 			ID: "2",
		// 		},
		// 		{
		// 			ID: "3",
		// 		},
		// 	},
		// 	Name:  "asdf",
		// 	Hours: 3,
		// },
		// {
		// 	ID: 4,
		// 	Teacher: &generation.Teacher{
		// 		ID:   1,
		// 		Name: "a",
		// 	},
		// 	Room: &generation.Room{
		// 		ID: "1",
		// 	},
		// 	Groups: []*generation.Group{
		// 		{
		// 			ID: "2",
		// 		},
		// 		{
		// 			ID: "3",
		// 		},
		// 	},
		// 	Name:  "asdf",
		// 	Hours: 2,
		// },
	}

	test.chromosome.Groups = []*generation.Group{
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

	test.chromosome.TimeTable = &generation.TimeTable{
		Classes: []*generation.Class{
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
		},

		CommonClasses: test.chromosome.CommonClasses,

		Groups: test.chromosome.Groups,
	}

	test.chromosome.TimeTable.Init(5, 2)

	test.chromosome.Init(5, 2)
	t.Run("letsgo", func(t *testing.T) {
		for group, value := range test.chromosome.Genes {
			fmt.Println(group)
			fmt.Println()

			for key, i := range value.Slots {
				fmt.Println(key)
				fmt.Println(i)
			}
			fmt.Println()
			fmt.Println()

			fmt.Println()

		}
		fmt.Println(test.chromosome.Fitness)
	})
}
