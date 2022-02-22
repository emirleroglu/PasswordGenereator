package excel

import (
	"fmt"

	"strconv"

	"github.com/xuri/excelize/v2"
)

// initialize excel
func WriteExcelInit() {
	file := excelize.NewFile()
	index := file.NewSheet("Password")
	file.SetColWidth("Password", "B", "D", 30)
	file.SetCellValue("Password", "B1", "Email")
	file.SetCellValue("Password", "C1", "Password")
	file.SetCellValue("Password", "A1", "Id")
	file.SetCellValue("Password", "D1", "Domain")
	file.SetCellValue("Password", "E2", "Counter")
	file.SetCellValue("Password", "E3", 2)
	file.SetActiveSheet(index)
	if err := file.SaveAs("secret.xlsx"); err != nil {
		fmt.Println(err)
	}
}

// function adding records to excel
func ExcelAddRecord(email string, password string, domain string) (done bool) {
	file, err := excelize.OpenFile("secret.xlsx")
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer func() {
		if err := file.Close; err != nil {
			fmt.Println(err)
		}

	}()
	// find counter E3
	cell, err := file.GetCellValue("Password", "E3")
	if err != nil {
		fmt.Println(err)
		return false
	}

	file.SetCellValue("Password", "B"+cell, email)
	file.SetCellValue("Password", "C"+cell, password)
	file.SetCellValue("Password", "D"+cell, domain)
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
	return true

}

// Reading excel and find mail and password
func ReadExcel(domain string) (email string, pass string) {
	file, err := excelize.OpenFile("secret.xlsx")
	if err != nil {
		fmt.Println(err)
		return "err", "err"
	}
	defer func() {
		if err := file.Close; err != nil {
			fmt.Println(err)
		}

	}()
	result, err := file.SearchSheet("Password", domain)
	if err != nil {
		fmt.Println(err.Error)
		return "err", "err"
	}
	x := result[0][1]
	id := string(x)

	mail, err := file.GetCellValue("Password", "B"+id)
	if err != nil {
		fmt.Println(err)
		return "err", "err"
	}

	password, err := file.GetCellValue("Password", "C"+id)
	if err != nil {
		fmt.Println(err)
		return "err", "err"
	}

	return mail, password

}
