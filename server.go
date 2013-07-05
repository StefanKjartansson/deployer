package deployer

import (
	"code.google.com/p/go.net/websocket"
	"code.google.com/p/gorilla/mux"
	"fmt"
	"net/http"
)

func configServer() {

	idRe := "{id:[0-9a-f]{8}(-?[0-9a-f]{4}){3}-?[0-9a-f]{12}}"
	router := new(mux.Router)

	router.HandleFunc("/",
		IndexHandler)

	s := router.PathPrefix("/projects").Subrouter()

	s.HandleFunc("/", ProjectListHandler).Methods("GET")
	s.HandleFunc(fmt.Sprintf("/%s/", idRe), ProjectDetailHandler).Methods("GET")
	s.HandleFunc(fmt.Sprintf("/%s/deploy", idRe), DeployHandler).Methods("POST")

	http.Handle("/", router)
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	http.Handle("/ws", websocket.Handler(wsHandler))
}
