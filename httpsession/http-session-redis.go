package httpsession

import (
	"fmt"
	"log"
	"net/http"

	redisStore "gopkg.in/boj/redistore.v1"
)

const (
	SESSSION_NAME = "session-name"
)

var sessionStore *redisStore.RediStore
var err error

func init() {
	sessionStore, err = redisStore.NewRediStore(10, "tcp", ":6379", "", []byte("secret-key"))

	if err != nil {
		log.Fatal("error getting redis store:", err)
	}

}

func redisHome(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, SESSSION_NAME)
	var authenticated interface{} = session.Values["authenticated"]
	if authenticated != nil {
		isAuthenticated := session.Values["authenticated"].(bool)
		if !isAuthenticated {
			http.Error(w, "You are unauthorized to view the page", http.StatusForbidden)
			return
		}
		fmt.Fprintln(w, "Home Page")
	} else {
		http.Error(w, "You are unauthrized to view the page", http.StatusForbidden)
	}
}

func redisLogin(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, SESSSION_NAME)
	session.Values["authenticated"] = true
	if err = session.Save(r, w); err != nil {
		log.Fatalf("Error saving session: %v", err)
	}
	fmt.Fprintln(w, "You have successfully logged in.")

}

func redisLogout(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, SESSSION_NAME)
	session.Values["authenticated"] = false
	if err := session.Save(r, w); err != nil {
		log.Fatalf("Error saving session: %v", err)
	}

	fmt.Fprintln(w, "You have successfully logged out.")
}

func Redis() {

	http.HandleFunc("/home", redisHome)
	http.HandleFunc("/login", redisLogin)
	http.HandleFunc("/logout", redisLogout)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	defer sessionStore.Close()
	if err != nil {
		log.Fatal("error starting http server: ", err)
	}

}
