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

	test.chromosome.Days = 10
	test.chromosome.Hours = 5

	test.chromosome.CommonClasses = []*generation.CommonClass{
		{
			ID: 1,
			Teacher: &generation.Teacher{
				ID:   2,
				Name: "bbb",
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
			Name:  "math",
			Hours: 3,
		},
		{
			ID: 4,
			Teacher: &generation.Teacher{
				ID:   1,
				Name: "aaa",
			},
			Room: &generation.Room{
				ID: "2",
			},
			Groups: []*generation.Group{
				{
					ID: "2",
				},
				{
					ID: "3",
				},
			},
			Name:  "eng",
			Hours: 2,
		},
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
		Hours: 5,
		Days:  10,
		Classes: []*generation.Class{
			{
				ID: 3,
				Teacher: &generation.Teacher{
					ID:   1,
					Name: "aaa",
				},
				Room: &generation.Room{
					ID: "1",
				},
				Group: &generation.Group{
					ID: "1",
				},
				Name:  "rus",
				Hours: 10,
			},
			{
				ID: 2,
				Teacher: &generation.Teacher{
					ID:   1,
					Name: "aaa",
				},
				Room: &generation.Room{
					ID: "5",
				},
				Group: &generation.Group{
					ID: "2",
				},
				Name:  "lit",
				Hours: 10,
			},
		},

		CommonClasses: test.chromosome.CommonClasses,

		Groups: test.chromosome.Groups,
	}

	test.chromosome.TimeTable.Init()

	test.chromosome.Init()
	t.Run("letsgo", func(t *testing.T) {
		for group, value := range test.chromosome.Genes {
			fmt.Println(group)
			fmt.Println()

			for key, i := range value.Slots {
				fmt.Println(key)
				fmt.Println(i)
				fmt.Println()
				fmt.Println()

				fmt.Println()

			}
			fmt.Println()
			fmt.Println()

			fmt.Println()

		}
		generation.SaveXLSX(test.chromosome)
		fmt.Println(test.chromosome.Fitness)
	})
}
