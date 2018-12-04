package html

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/gorilla/securecookie"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32),
)

func getUserName(request *http.Request) (userName string) {
	cookie, err := request.Cookie("session")
	if err == nil {
		cookieValue := make(map[string]string)
		err = cookieHandler.Decode("session", cookie.Value, &cookieValue)
		if err == nil {
			userName = cookieValue["username"]
		}
	}
	return userName
}

func setSession(userName string, response http.ResponseWriter) {
	value := map[string]string{
		"username": userName,
	}
	encode, err := cookieHandler.Encode("session", value)
	if err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encode,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}

	http.SetCookie(response, cookie)
}

func login(response http.ResponseWriter, request *http.Request) {
	username := request.FormValue("username")
	password := request.FormValue("password")
	target := "/"
	if username != "" && password != "" {
		setSession(username, response)
		target = "/home"
	}
	http.Redirect(response, request, target, http.StatusFound)
}

func logout(response http.ResponseWriter, request *http.Request) {
	clearSession(response)
	http.Redirect(response, request, "/", http.StatusFound)
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("static/login-form.html")
	parsedTemplate.Execute(w, nil)
}

func homePage(respone http.ResponseWriter, request *http.Request) {
	userName := getUserName(request)
	if userName != "" {
		data := map[string]interface{}{"userName": userName}

		parsedTemplate, _ := template.ParseFiles("static/home.html")
		parsedTemplate.Execute(respone, data)
	} else {
		http.Redirect(respone, request, "/", http.StatusFound)
	}
}

func HttpForm() {
	var router = mux.NewRouter()
	router.HandleFunc("/", loginPage)
	router.HandleFunc("/home", homePage)
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/logout", logout).Methods("POST")
	http.Handle("/", router)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("error starting http server: ", err)
		return
	}
}
