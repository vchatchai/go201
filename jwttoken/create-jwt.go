package jwttoken

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"

	"github.com/dgrijalva/jwt-go"
)

const (
	CONN_HOST             = "localhost"
	CONN_PORT             = "8080"
	CLAIM_ISSUER          = "ee56054"
	CLAIM_EXPIRY_IN_HOURS = 24
)

type Employee struct {
	Id        int    `json:"id"`
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
}

type Employees []Employee

var employeees Employees

func init() {
	employeees = []Employee{
		Employee{Id: 1, Firstname: "Foo", Lastname: "Bar"},
		Employee{Id: 2, Firstname: "Bas", Lastname: "Qux"},
	}
}

var signature = []byte("secret")

func getToken(w http.ResponseWriter, r *http.Request) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * CLAIM_EXPIRY_IN_HOURS).Unix(),
		Issuer:    CLAIM_ISSUER,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(signature)
	w.Write([]byte(tokenString))
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API is up and runnings")
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(employeees)
}

func CreateJWt() {
	mux := mux.NewRouter()
	mux.HandleFunc("/status", getStatus).Methods("GET")
	mux.HandleFunc("/get-token", getToken).Methods("GET")
	mux.HandleFunc("/getEmployees", getEmployees).Methods("GET")

	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, mux)
	handlers.LoggingHandler(os.Stdout, mux)
	if err != nil {
		panic(err)
	}
}
