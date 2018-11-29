package html

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/gorilla/schema"

	"github.com/asaskevich/govalidator"
)

const (
	USERNAME_ERROR_MESSAGE = "Please enter a valid Username"
	PASSWORD_ERROR_MESSAGE = "Please enter a valid Passwords"
	GENERIC_ERROR_MESSAGE  = "Validation Error"
)

func validateUser(w http.ResponseWriter, r *http.Request, user *User) (bool, string) {
	valid, validationError := govalidator.ValidateStruct(user)
	if !valid {
		usernameError := govalidator.ErrorByField(validationError, "Username")
		passwordError := govalidator.ErrorByField(validationError, "Password")
		if usernameError != "" {
			log.Printf("username validation error: %s ", usernameError)
			return valid, USERNAME_ERROR_MESSAGE
		}
		if passwordError != "" {
			log.Printf("password validation error : %s", passwordError)
			return valid, PASSWORD_ERROR_MESSAGE
		}
	}

	return valid, GENERIC_ERROR_MESSAGE
}

func loginGet(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("login-form.html")
	template.Execute(w, nil)
}

func loginPost(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	r.ParseForm()
	decoder := schema.NewDecoder()
	decodeError := decoder.Decode(user, r.Form)
	if decodeError != nil {
		log.Printf("error mapping parsed form data to struct: %s", decodeError)
		return
	}

	fmt.Fprintf(w, "Hello %s ", user.Username)

}

func Validator() {
	router := mux.NewRouter()
	router.HandleFunc("/login", loginGet).Methods("GET")
	router.HandleFunc("/login", loginPost).Methods("POST")
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
	if err != nil {
		panic(err)
	}
}
