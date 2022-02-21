package excel

import (
	"fmt"

	"strconv"

	"github.com/xuri/excelize/v2"
)

func WriteExcelInit() {
	file := excelize.NewFile()
	file.NewSheet("Password")
	file.SetCellValue("Password", "B1", "Email")
	file.SetCellValue("Password", "C1", "Password")
	file.SetCellValue("Password", "A1", "Id")
	file.SetCellValue("Password", "E2", "Counter")
	file.SetCellValue("Password", "E3", 2)

	if err := file.SaveAs("secret.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func ExcelAddRecord(email string, password string) {
	file, err := excelize.OpenFile("secret.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := file.Close; err != nil {
			fmt.Println(err)
		}

	}()

	cell, err := file.GetCellValue("Password", "E3")
	if err != nil {
		fmt.Println(err)
		return
	}

	file.SetCellValue("Password", "B"+cell, email)
	file.SetCellValue("Password", "C"+cell, password)
	i, err := strconv.Atoi(cell)
	if err != nil {
		// handle error
		fmt.Println(err)

	}
	var id int = i
	id--
	i++
	file.SetCellValue("Password", "A"+cell, id)

	file.SetCellValue("Password", "E3", i)

	if err := file.SaveAs("secret.xlsx"); err != nil {
		fmt.Println(err)
	}

}

func ReadExcel() {
	file, err := excelize.OpenFile("secret.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := file.Close; err != nil {
			fmt.Println(err)
		}

	}()

	cell, err := file.GetCellValue("Password", "B1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
}
