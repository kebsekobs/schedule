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

	test.gnrt.Rooms = []*generation.Room{
		{
			ID:       "1",
			Capacity: 10,
		},
		{
			ID:       "2",
			Capacity: 15,
		},
		{
			ID:       "3",
			Capacity: 20,
		},
	}

	test.gnrt.CommonClasses = []*generation.CommonClass{
		{
			ID: 1,
			Teacher: &generation.Teacher{
				ID:   2,
				Name: "b",
			},
			Groups: []*generation.Group{
				{
					ID:       "1",
					Quantity: 5,
				},
				{
					ID:       "2",
					Quantity: 7,
				},
				{
					ID:       "3",
					Quantity: 8,
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
			Groups: []*generation.Group{
				{
					ID:       "2",
					Quantity: 7,
				},
				{
					ID:       "3",
					Quantity: 8,
				},
			},
			Name:  "test2",
			Hours: 5,
		},
	}

	test.gnrt.Groups = []*generation.Group{
		{
			ID:       "1",
			Quantity: 5,
		},
		{
			ID:       "2",
			Quantity: 7,
		},
		{
			ID:       "3",
			Quantity: 8,
		},
	}
	test.gnrt.Classes = []*generation.Class{
		{
			ID: 1,
			Teacher: &generation.Teacher{
				ID:   1,
				Name: "a",
			},
			Group: &generation.Group{
				ID:       "1",
				Quantity: 5,
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
			Group: &generation.Group{
				ID:       "2",
				Quantity: 7,
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
