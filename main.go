package main

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

func main() {
	errWkbk, err := excelize.OpenFile("test-autofilter-issue.xlsx")
	if err != nil {
		log.Fatalf("err opening excel wkbk : '%v'", err)
	}
	sheetNm := "test"

	err = errWkbk.AutoFilter(sheetNm, fmt.Sprintf("A1:%v%v", "E", 5), nil)
	if err != nil {
		log.Printf("err filtering the %s sheet: %v", sheetNm, err)
	}
}
