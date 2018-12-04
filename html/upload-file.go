package html

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func upload(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		log.Printf("err getting a file for the provided form key:  %s", err)
		return
	}

	defer file.Close()

	out, pathError := ioutil.TempFile(os.TempDir(), "upload")
	if pathError != nil {
		panic(err)
	}

	defer out.Close()
	_, copyErr := io.Copy(out, file)

	if copyErr != nil {
		panic(copyErr)
	}

	fmt.Fprintf(w, "File uploaded successfully: %s  path %s", header.Filename, out.Name())

}

func index(w http.ResponseWriter, r *http.Request) {
	uploadfile, err := template.ParseFiles("static/upload-file.html")
	if err != nil {
		panic(err)
	}

	uploadfile.Execute(w, nil)
}

func UploadFile() {
	router := mux.NewRouter()
	router.HandleFunc("/", index)
	router.HandleFunc("/upload", upload)
	http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
}
