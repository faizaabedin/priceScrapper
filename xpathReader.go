package main

import (
	"fmt"
)


//make a request with xpath quest for each xapth for each provider
//for each provider - starting from row 2
//make request with with each xpath - starting from column 3
//what u get from the request made then u save in variables - which is the column names
//then depending on the values of the variables u get update the sheets


func main() {

	//partsArray := [...]string{"Motherboard-AMD", "Motherboard-Intel", "CPU-AMD", "CPU-Intel", "APU-AMD", "APU-Intel","Memory","SSD","PSU","Cases"}
	//partsProviderArray := [...]{"Pc-Canada","newegg","shoprbc","vuugo","memoryExpress","staples","amazon.ca","canadaComputers","mikesComputerParts","Dell"}
	KeyFileName := "secret_key.json"
	SheetId := "1EDgr3SdhHXWGCBDi9EZGwEeVkvjBzQPox2c2to4bOPM"
	sheetIndex := 3
	productDesColNo := 2
	columnDesColNo := 1

		//connect to the sheets
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




//// this is to connect to the spreedsheets
//func connectToGoogleSheet(KeyFileName string, SheetId string, sheetIndex int) *spreadsheet.Sheet{
//
//	data, err := ioutil.ReadFile(KeyFileName)
//	checkError(err)
//
//	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
//	checkError(err)
//
//	client := conf.Client(context.TODO())
//	service := spreadsheet.NewServiceWithClient(client)
//	spreadsheet, err := service.FetchSpreadsheet(SheetId)
//	checkError(err)
//
//	//strconv.FormatUint(uint64(u), 10)
//	index := uint(float64(sheetIndex))
//	sheet, err := spreadsheet.SheetByIndex(index)
//	checkError(err)
//
//	return sheet
//}
//
////checks error is something went wrong
//func checkError(err error) {
//	if err != nil {
//		panic(err.Error())
//	}
//}
