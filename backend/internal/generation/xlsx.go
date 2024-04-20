package generation

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func SaveXLSX(chromosome Chromosome) {
	// Пример двумерного массива данных

	// Создаем новый Excel файл
	xlsx := excelize.NewFile()

	// Создаем новый лист
	xlsx.NewSheet("Sheet1")
	var c int
	// Записываем данные из двумерного массива на лист
	for groupID, gene := range chromosome.Genes {
		cell := excelize.ToAlphaString(c+1) + "1"
		xlsx.SetCellValue("Sheet1", cell, groupID)
		for r, slot := range gene.Slots {
			cell := excelize.ToAlphaString(c+1) + fmt.Sprint(slot.Value+2)

			var cellValue string

			if chromosome.TimeTable.GroupSlots[groupID].Classes[r] != nil {
				cellValue = chromosome.TimeTable.GroupSlots[groupID].Classes[r].Name + "\n" +
					chromosome.TimeTable.GroupSlots[groupID].Classes[r].Teacher.Name + "\n" +
					chromosome.TimeTable.GroupSlots[groupID].Classes[r].Room.ID
			}

			xlsx.SetCellValue("Sheet1", cell, cellValue)
		}
		c++
	}

	// Сохраняем файл
	err := xlsx.SaveAs("output.xlsx")
	if err != nil {
		fmt.Println("Error saving file:", err)
		return
	}

	fmt.Println("Excel файл успешно создан")
}
