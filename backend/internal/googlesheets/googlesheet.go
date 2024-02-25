package googlesheets

import (
	"fmt"

	"google.golang.org/api/sheets/v4"
)

type Servicehandle struct {
	Googlesheet *sheets.Service
}

// Метод для получения числа заполненных строк в таблице
// Получает: spreadsheetId string - айди таблицы, sheet string - имя листа
// Возвращает число строк
func (service Servicehandle) GetLenght(spreadsheetId, sheet string) (int, error) {
	resp, err := service.Googlesheet.Spreadsheets.Values.Get(spreadsheetId, sheet).Do()
	if err != nil {
		return 0, err
	}
	return len(resp.Values), nil
}

// Метод для получения всех записей из таблицы, начиная с указанной позиции
// Получает: spreadsheetId string - айди таблицы, readRange string - диапазон вида "ИмяЛиста!A1:W",
// состоящий из имени листа и позиции, начиная с которой мы вытаскиваем данные
// Возвращает значения из таблицы в виде мапы map[int]map[int]string
func (service Servicehandle) GetData(
	spreadsheetId string,
	readRange string,
) (map[int]map[int]string, error) {
	resp, err := service.Googlesheet.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()

	tableData := map[int]map[int]string{}
	if err != nil {
		return tableData, fmt.Errorf("unable to retrieve data from spreadsheet: %v", err)
	}

	for i, row := range resp.Values {
		strMap := map[int]string{}
		for j, val := range row {
			if val != nil {
				strMap[j] = val.(string)
			} else {
				strMap[j] = ""
			}
		}
		tableData[i] = strMap
	}
	return tableData, nil

}

// Метод для получения всех записей из таблицы, начиная с указанной позиции
// Получает: spreadsheetId string - айди таблицы, readRange string - диапазон вида "ИмяЛиста!A1:W", состоящий из имени листа
// и позиции, начиная с которой мы вытаскиваем данные
// Возвращает значения из таблицы в виде среза [][]string
func (service Servicehandle) GetDataSlices(
	spreadsheetId string,
	readRange string,
) ([][]string, error) {
	resp, err := service.Googlesheet.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()

	tableData := [][]string{}
	if err != nil {
		return tableData, fmt.Errorf("unable to retrieve data from spreadsheet: %v", err)
	}

	for _, row := range resp.Values {
		rowSlice := []string{}
		for _, val := range row {
			if val != nil {
				rowSlice = append(rowSlice, val.(string))
			} else {
				rowSlice = append(rowSlice, "")
			}
		}
		tableData = append(tableData, rowSlice)
	}
	return tableData, nil

}

// Метод для очистки ячеек таблицы
// Получает: spreadsheetId string - айди таблицы, sheet string - имя листа, position - позиция,
// начиная с которой мы удаляяем данные
func (service Servicehandle) ClearDataGoogleSheetStrings(spreadsheetId, sheet, position string) error {
	_, err := service.Googlesheet.Spreadsheets.Values.Clear(spreadsheetId, sheet+"!"+position, &sheets.ClearValuesRequest{}).Do()
	if err != nil {
		return fmt.Errorf("unable to Clear datasheet. %v", err)
	}
	return nil
}

// Метод для вставки данных с предворительной очисткой таблицы
// Получает: spreadsheetId string - айди таблицы, sheet string - имя листа,
// records [][]interface - двумерный массив с данными
func (service Servicehandle) ClearAndInsertIntoGoogleSheet(spreadsheetId, sheet string, records [][]interface{}) error {
	var vr sheets.ValueRange
	vr.Values = append(vr.Values, records...)
	_, err := service.Googlesheet.Spreadsheets.Values.Clear(spreadsheetId, sheet, &sheets.ClearValuesRequest{}).Do()
	if err != nil {
		return fmt.Errorf("unable to Clear datasheet. %v", err)
	}
	_, err = service.Googlesheet.Spreadsheets.Values.Update(spreadsheetId, sheet, &vr).ValueInputOption("RAW").Do()
	if err != nil {
		return fmt.Errorf("unable to Update datasheet. %v", err)
	}
	return nil
}

// Метод для вставки данных в таблицу, начиная с указанной позиции
// Получает: spreadsheetId string - айди таблицы, sheet string - имя листа, position - позиция,
// начиная с которой мы вставляем данные, recordsString [][]string - двумерный массив с данными
func (service Servicehandle) InsertIntoGoogleSheetStrings(spreadsheetId, sheet, position string, recordsString [][]string) error {
	records := formatRecordsStringToInterface(recordsString)
	return service.InsertIntoGoogleSheet(spreadsheetId, sheet, position, records)
}

// Метод для вставки данных в таблицу, начиная с указанной позиции
// Получает: spreadsheetId string - айди таблицы, sheet string - имя листа, position - позиция,
// начиная с которой мы вставляем данные, records [][]interface - двумерный массив с данными
func (service Servicehandle) InsertIntoGoogleSheet(spreadsheetId, sheet, position string, records [][]interface{}) error {
	var vr sheets.ValueRange
	vr.Values = append(vr.Values, records...)
	_, err := service.Googlesheet.Spreadsheets.Values.Clear(spreadsheetId, sheet+"!"+position, &sheets.ClearValuesRequest{}).Do()
	if err != nil {
		return fmt.Errorf("unable to Clear datasheet. %v", err)
	}
	_, err = service.Googlesheet.Spreadsheets.Values.Update(spreadsheetId, sheet+"!"+position, &vr).ValueInputOption("RAW").Do()
	if err != nil {
		return fmt.Errorf("unable to Update datasheet. %v", err)
	}
	return nil
}

// Метод для добавления данных в таблицу к уже имеющимя там записям
// Получает: spreadsheetId string - айди таблицы, sheet string - имя листа, position - позиция,
// начиная с которой мы вставляем данные, records [][]interface - двумерный массив с данными
func (service Servicehandle) AppendToGoogleSheet(spreadsheetId, sheet, position string, records [][]interface{}) error {
	var vr sheets.ValueRange
	vr.Values = append(vr.Values, records...)
	_, err := service.Googlesheet.Spreadsheets.Values.Append(spreadsheetId, sheet+"!"+position, &vr).ValueInputOption("RAW").Do()
	if err != nil {
		return fmt.Errorf("unable to Update datasheet. %v", err)
	}
	return nil
}

// Вспомогательная функция, преобразующая двумерный срез строк в двумерный срез интерфейсов
func formatRecordsStringToInterface(recordsString [][]string) [][]interface{} {
	var formatedTable [][]interface{}
	for _, row := range recordsString {
		var formatedRow []interface{}
		for _, cell := range row {
			formatedRow = append(formatedRow, cell)
		}
		formatedTable = append(formatedTable, formatedRow)
	}
	return formatedTable
}
