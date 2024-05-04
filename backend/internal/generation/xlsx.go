package generation

import (
	"fmt"
	"log"
	"strconv"

	"github.com/tealeg/xlsx"
)

func SaveXLSX(chromosome Chromosome) {
	// Создаем новый Excel файл
	xlsx := xlsx.NewFile()

	// Создаем новый лист
	sheet, err := xlsx.AddSheet("Sheet1")
	if err != nil {
		fmt.Println("Error creating sheet:", err)
		return
	}

	var c int

	// Записываем данные из двумерного массива на лист
	for groupID, gene := range chromosome.Genes {
		sheet.Cell(0, c+1).SetValue(groupID)

		for r, slot := range gene.Slots {

			if chromosome.TimeTable.GroupSlots[groupID].Classes[r] != nil {
				cellValue := "Предмет: " + chromosome.TimeTable.GroupSlots[groupID].Classes[r].Name + "\n" +
					"Преподаватель: " + chromosome.TimeTable.GroupSlots[groupID].Classes[r].Teacher.Name + "\n" +
					"Ауд.: " + chromosome.TimeTable.GroupSlots[groupID].Classes[r].Room.ID + "\n" +
					"Вместимость ауд.: " + strconv.Itoa(chromosome.TimeTable.GroupSlots[groupID].Classes[r].Room.Capacity) + "\n" +
					"Кол-во студентов: " + strconv.Itoa(chromosome.TimeTable.GroupSlots[groupID].Classes[r].Group.Quantity)

				sheet.Cell(slot.Value+1, c+1).SetValue(cellValue)
			}
		}
		c++
	}

	// Сохраняем файл
	err = xlsx.Save("data/output.xlsx")
	if err != nil {
		log.Println("Error saving file:", err)
		return
	}

	fmt.Println("Excel файл успешно создан")
}

func ParseRooms(file *xlsx.File) ([]*Room, error) {
	file, err := xlsx.OpenFile("data/rooms.xlsx")
	if err != nil {
		return nil, err
	}

	var data []*Room

	for _, sheet := range file.Sheets {
		for _, row := range sheet.Rows {
			id := row.Cells[0].String()
			capacity, err := row.Cells[1].Int()
			if err != nil {
				continue
			}
			data = append(data, &Room{
				ID:       id,
				Capacity: capacity,
			})
		}
	}

	return data, nil
}
