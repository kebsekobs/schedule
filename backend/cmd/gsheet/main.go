package main

import (
	"fmt"

	"github.com/kebsekobs/schedule/backend/internal/googlesheets"
)

func main() {

	test := struct {
		name             string
		googleSheetId    string
		googleSheetRange string
		want             int
		want2            map[int]map[int]string
	}{}

	test.name = "base getData test"
	test.googleSheetId = "1ugqfiEdWiyP7Dt8d7zqWCGrSs2xQ4pMeQgiYXJuTMfs"
	test.googleSheetRange = "Sheet"

	conn, err := googlesheets.InitGoogleSheetConnection()
	if err != nil {
		panic(err)
	}
	data, err := conn.GetData(test.googleSheetId, test.googleSheetRange)
	if err != nil {
		panic(err)
	}
	fmt.Println(data[5])
}
