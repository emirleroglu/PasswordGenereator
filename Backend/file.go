package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/excel"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type record struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Domain   string `json:"domain"`
}

type domain struct {
	Domain string `json:"domain"`
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/init", createExcelFile).Methods("GET")
	router.HandleFunc("/addRecord", addRecord).Methods("POST")
	router.HandleFunc("/read", getRecord).Methods("POST")

	//	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	//	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	//	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	//corsObj := handlers.AllowedOrigins([]string{"*"})
	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8080"}, //you service is available and allowed for this base url
		AllowedMethods: []string{
			http.MethodGet, //http methods for your app
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		}})
	// start server listen
	// with error handling
	//log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
	//log.Fatal(http.ListenAndServe(":8080", handlers.CORS(corsObj)(router)))
	http.ListenAndServe(":8080", corsOpts.Handler(router))

}

func addRecord(w http.ResponseWriter, r *http.Request) {
	var myRecord record
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Wrong !")
	}

	json.Unmarshal(reqBody, &myRecord)

	excel.ExcelAddRecord(myRecord.Email, myRecord.Password, myRecord.Domain)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(myRecord)
}

func getRecord(w http.ResponseWriter, r *http.Request) {
	var myDomain domain
	var myRecord record
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Wrong !")
	}
	json.Unmarshal(reqBody, &myDomain)
	email, password := excel.ReadExcel(myDomain.Domain)
	myRecord.Email = email
	myRecord.Password = password
	myRecord.Domain = myDomain.Domain
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(myRecord)

}

func createExcelFile(w http.ResponseWriter, r *http.Request) {
	excel.WriteExcelInit()
}
