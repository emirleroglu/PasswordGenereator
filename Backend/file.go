package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"main/excel"
	"net/http"

	"github.com/gorilla/mux"
)

type record struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Domain   string `json:"domain"`
}

func main() {
	excel.WriteExcelInit()
	//var x = excel.ExcelAddRecord("emirlerogluhalil@gmail.com", "12345", "github.com")
	//fmt.Println(x)
	//email, pass := excel.ReadExcel("github.com")
	//fmt.Println(email, pass)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/addRecord", addRecord).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func addRecord(w http.ResponseWriter, r *http.Request) {
	var myRecord record
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &myRecord)

	excel.ExcelAddRecord(myRecord.Email, myRecord.Password, myRecord.Domain)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(myRecord)
}

func getRecord(w http.ResponseWriter, r *http.Request) {

}
