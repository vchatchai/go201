package html

import (
	"net/http"
)

func Static() {

	fileSever := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileSever))
	http.HandleFunc("/", renderTemplate)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		panic(err)
	}
}
