package deployer

import (
	"code.google.com/p/gorilla/mux"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"time"
)

const (
	json_header = "application/json; charset=utf-8"
)

type Context struct {
}

func IndexHandler(w http.ResponseWriter, req *http.Request) {

	t, err := template.ParseFiles("./templates/index.html")

	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, Context{})

}

func ProjectListHandler(w http.ResponseWriter, req *http.Request) {

	log.Println("List handler")
	hasWritten := false

	w.Header().Set("Content-Type", json_header)
	w.Write([]byte("["))

	for _, p := range Projects {
		if hasWritten {
			w.Write([]byte(","))
		}
		b, err := json.Marshal(p)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(b)
		hasWritten = true
	}
	w.Write([]byte("]"))
}

func ProjectDetailHandler(w http.ResponseWriter, req *http.Request) {

	log.Println("Detail handler")
	vars := mux.Vars(req)
	w.Header().Set("Content-Type", json_header)

	b, err := json.Marshal(Projects[vars["id"]])
	if err != nil {
		log.Fatal(err)
	}
	w.Write(b)
}

func DeployHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Deploy handler")
	vars := mux.Vars(req)
	w.Header().Set("Content-Type", json_header)

	j := Job{
		ID:        Uuid(),
		Status:    "not started",
		ProjectID: vars["id"],
		Started:   time.Now(),
	}

	b, err := json.Marshal(j)
	if err != nil {
		log.Fatal(err)
	}
	go j.Start()
	log.Println("Called start")
	w.Write(b)
}
