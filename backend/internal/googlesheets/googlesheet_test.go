package googlesheets

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGoogleSheetGet(t *testing.T) {

	test := struct {
		name             string
		googleSheetId    string
		googleSheetRange string
		want             int
		want2            map[int]map[int]string
		conn             *Servicehandle
	}{}

	var err error

	test.name = "base getData test"
	test.googleSheetId = "1sPfcYLOYLHwD6jquD5hbW7BGhKbdsjNd"
	test.googleSheetRange = "Sheet"
	test.want = 0
	test.want2 = map[int]map[int]string{0: {0: "sample_A1", 1: "sample_B1"}, 1: {0: "sample_A2", 1: "sample_B2"}}
	test.conn, err = InitGoogleSheetConnection()
	if err != nil {
		fmt.Printf("Ошибка создания подключения, завершаем тест: \n%s", err)
	}

	t.Run(test.name, func(t *testing.T) {

		if got, err := test.conn.GetData(test.googleSheetId, test.googleSheetRange); err != nil || !reflect.DeepEqual(got, test.want2) {
			if err != nil {
				t.Errorf(err.Error())
			} else {
				t.Errorf("candy() = %v, want %v", got, test.want2)
			}
		}
	})
}
func TestGoogleSheetInsert(t *testing.T) {

	test := struct {
		name                string
		googleSheetId       string
		googleSheetSheet    string
		googleSheetPosition string
		data                [][]string
		conn                *Servicehandle
	}{}

	var err error

	test.name = "base getData test"
	test.googleSheetId = "1sPfcYLOYLHwD6jquD5hbW7BGhKbdsjNd"
	test.googleSheetSheet = "Лист1"
	test.googleSheetPosition = "C1:D"

	test.data = [][]string{{"A1", "A2"}, {"B1", "B2"}}

	test.conn, err = InitGoogleSheetConnection()
	if err != nil {
		t.Errorf("Ошибка создания подключения, завершаем тест: \n%s", err)
		return
	}

	t.Run(test.name, func(t *testing.T) {
		if err := test.conn.InsertIntoGoogleSheetStrings(test.googleSheetId, test.googleSheetSheet, test.googleSheetPosition, test.data); err != nil {
			if err != nil {
				t.Errorf(err.Error())
			}
		}
	})
}
func TestGoogleSheetClear(t *testing.T) {

	test := struct {
		name                string
		googleSheetId       string
		googleSheetSheet    string
		googleSheetPosition string
		data                [][]string
		conn                *Servicehandle
	}{}

	var err error

	test.name = "base getData test"
	test.googleSheetId = "1sPfcYLOYLHwD6jquD5hbW7BGhKbdsjNd"
	test.googleSheetSheet = "Лист1"
	test.googleSheetPosition = "C1:C"

	test.conn, err = InitGoogleSheetConnection()
	if err != nil {
		t.Errorf("Ошибка создания подключения, завершаем тест: \n%s", err)
		return
	}

	t.Run(test.name, func(t *testing.T) {
		if err := test.conn.ClearDataGoogleSheetStrings(test.googleSheetId, test.googleSheetSheet, test.googleSheetPosition); err != nil {
			if err != nil {
				t.Errorf(err.Error())
			}
		}
	})
}

func TestGoogleSheetSimpleGet(t *testing.T) {

	test := struct {
		name             string
		googleSheetId    string
		googleSheetSheet string
		conn             *Servicehandle
	}{}

	var err error

	test.name = "base getData test"
	test.googleSheetId = "1sPfcYLOYLHwD6jquD5hbW7BGhKbdsjNd"

	test.conn, err = InitGoogleSheetConnection()
	if err != nil {
		t.Errorf("Ошибка создания подключения, завершаем тест: \n%s", err)
		return
	}

	t.Run(test.name, func(t *testing.T) {
		if len, err := test.conn.GetLenght(test.googleSheetId, "Data!A1:A"); err != nil {
			if err != nil {
				t.Errorf(err.Error())
			}
		} else {
			fmt.Println(len)
		}
	})
}
