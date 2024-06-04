package generation

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/tealeg/xlsx"
	"golang.org/x/text/unicode/norm"
)

func SaveXLSX(chromosome Chromosome, name string) {
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
		if _, ok := chromosome.TimeTable.GroupSlots[groupID]; !ok {
			continue
		}

		for r, slot := range gene.Slots {

			if chromosome.TimeTable.GroupSlots[groupID].Classes[r].ID != 0 {
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
	err = xlsx.Save("data/output_" + name + ".xlsx")
	if err != nil {
		log.Println("Error saving file:", err)
		return
	}

	fmt.Println("Excel файл успешно создан")
}

func ParseRooms() ([]*Room, error) {
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

func ParseClasses() (map[int]*Class, map[string]*Group, map[string]*Teacher, error) {
	xlFile, err := xlsx.OpenFile("data/classes.xlsx")
	if err != nil {
		return nil, nil, nil, err
	}

	groupsMap := make(map[string]*Group)
	teacherMap := make(map[string]*Teacher)
	classMap := make(map[int]*Class)

	for _, sheet := range xlFile.Sheets {
		for rowIndex, row := range sheet.Rows {
			if rowIndex < 5 {
				// Пропуск обработки строк до 6-й строки
				continue
			}

			// Пропустим пустые строки
			if len(row.Cells) == 0 {
				continue
			}

			// Получение значений из ячеек
			classID, err := row.Cells[columnToNumber("A")].Int() // ID занятия
			if err != nil {
				// log.Println(err)
				continue
			}
			name := row.Cells[columnToNumber("F")].String()       // Название предмета
			groupID := row.Cells[columnToNumber("I")].String()    // ID группы
			quantity, err := row.Cells[columnToNumber("J")].Int() // Кол-во учеников
			if err != nil {
				// log.Println(err)
				continue
			}
			classType := row.Cells[columnToNumber("L")].String()    // Тип занятия
			teacherName := row.Cells[columnToNumber("AB")].String() // Имя преподавателя
			streamID, err := row.Cells[columnToNumber("AE")].Int()  // ID потока
			if err != nil {
				streamID = 0
			}
			educationType := row.Cells[columnToNumber("AU")].String() // Форма обучения

			hours, err := row.Cells[columnToNumber("BB")].Int() // Часы в неделю
			if err != nil {
				// log.Println(err)
				continue
			}

			if name == "" {
				continue
			}

			if strings.ToLower(educationType) != "очная" {
				continue
			}

			if strings.ToLower(classType) != "лек" && strings.ToLower(classType) != "пр" {
				continue
			}

			// Применение остальных условий и фильтров по ячейкам для обработки данных
			if teacherName != "" && len(teacherName) > 1 {
				teacherName = trimFirstChar(teacherName)
			} else {
				continue
			}

			group := Group{
				ID:       groupID,
				Quantity: quantity,
			}

			teacher := Teacher{
				Name: teacherName,
			}

			if _, ok := groupsMap[groupID]; !ok {
				groupsMap[groupID] = &group
			}

			if _, ok := teacherMap[teacherName]; !ok {
				teacherMap[teacherName] = &teacher
			}

			if _, ok := classMap[classID]; !ok {
				classMap[classID] = &Class{
					ID:      streamID,
					Group:   &group,
					Hours:   hours,
					Teacher: &teacher,
					Name:    name,
				}
			}
		}
	}
	return classMap, groupsMap, teacherMap, nil
}

func columnToNumber(column string) int {
	columnNumber := 0
	power := 1
	for i := len(column) - 1; i >= 0; i-- {
		columnNumber += int(column[i]-'A'+1) * power
		power *= 26
	}
	return columnNumber - 1
}

func trimFirstChar(s string) string {
	_, size := utf8.DecodeRuneInString(s)
	normalized := norm.NFC.String(s)
	return normalized[size:]
}
