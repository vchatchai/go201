package httpsession

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/securecookie"
)

var cookiehandler *securecookie.SecureCookie

func init() {
	cookiehandler = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))
}

func createCookie(w http.ResponseWriter, r *http.Request) {
	value := map[string]string{
		"username": "Foo",
	}

	base64Encoded, err := cookiehandler.Encode("key", value)
	log.Println(value, base64Encoded)
	if err == nil {
		cookie := &http.Cookie{
			Name:  "first-cookie",
			Value: base64Encoded,
			Path:  "/",
		}

		http.SetCookie(w, cookie)
		w.Write([]byte(fmt.Sprintf("Cookie created.")))
	} else {
		w.Write([]byte(fmt.Sprintf("Cookie can't created.")))

	}
}

func readCookie(w http.ResponseWriter, r *http.Request) {
	log.Printf("Reading Cookie..")
	cookie, err := r.Cookie("first-cookie")
	if err == nil {
		value := make(map[string]string)
		err = cookiehandler.Decode("key", cookie.Value, &value)
		fmt.Println(cookie, value)
		if err == nil {
			w.Write([]byte(fmt.Sprintf("Hello %v \n", value["username"])))
		}
	} else {
		log.Printf("Cookie not found..")
		w.Write([]byte(fmt.Sprintf("Hello")))
	}
}

func Cookie() {
	http.HandleFunc("/create", createCookie)
	http.HandleFunc("/read", readCookie)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("error starting http server: ", err)
	}
}
