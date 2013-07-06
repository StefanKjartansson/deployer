package main

import (
    "net/http"
    "flag"
    "log"
    "github.com/StefanKjartansson/deployer"
)

func main() {

    var http_listen string
    flag.StringVar(&http_listen, "http", "127.0.0.1:3999", "host:port to listen on")
    flag.Parse()
    deployer.ConfigServer()
    log.Fatal(http.ListenAndServe(http_listen, nil))
}
