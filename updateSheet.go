package main

import (
	"context"
	"fmt"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
	"io/ioutil"
)

// get the {id,name} of the sheets  - make an set
// Read - update preVars
// Compare preVars vs curVars -> send notifications
// Look at the flags -> send notifications

func main() {

	partsArray := [...]string{"Motherboard-AMD", "Motherboard-Intel", "CPU-AMD", "CPU-Intel", "APU-AMD", "APU-Intel","Memory","SSD","PSU","Cases"}
	//partsProviderArray := [...]{"Pc-Canada","newegg","shoprbc","vuugo","memoryExpress","staples","amazon.ca","canadaComputers","mikesComputerParts","Dell"}
	KeyFileName := "secret_key.json"
	SheetId := "115grjUrcIqiK9UNfBrGjEPhtVhCoISvji3ZmvVIN0ZM"
	productDesColNo := 2
	columnDesColNo := 1

	//go through all the sheets
	for sheetIndex,_:= range partsArray {
		sheet := connectToGoogleSheet(KeyFileName,SheetId,sheetIndex)

		//length of rows in the sheet
		rowNum:= len(sheet.Rows)

		//for each row - excludes the first row - description
		for rowIndex := columnDesColNo; rowIndex < rowNum; rowIndex++ {

			//selected row among all rows
			row := sheet.Rows[rowIndex]

			//ignore first 2 columns start i=2, till row ends, iterate colums at ++2
			for columnIndex := productDesColNo; columnIndex < len(row); columnIndex+=2 {

				xpathLink := row[columnIndex]

				//if not empty take xpath and forward
				if xpathLink.Value != ""{
					fmt.Println(xpathLink.Value)
				}
			}
		}

	}

}


// this is to connect to the spreedsheets
func connectToGoogleSheet(KeyFileName string, SheetId string, sheetIndex int) *spreadsheet.Sheet{

	data, err := ioutil.ReadFile(KeyFileName)
	checkError(err)

	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	checkError(err)

	client := conf.Client(context.TODO())
	service := spreadsheet.NewServiceWithClient(client)
	spreadsheet, err := service.FetchSpreadsheet(SheetId)
	checkError(err)

	//strconv.FormatUint(uint64(u), 10)
	index := uint(float64(sheetIndex))
	sheet, err := spreadsheet.SheetByIndex(index)
	checkError(err)

	return sheet
}

//checks error is something went wrong
func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}