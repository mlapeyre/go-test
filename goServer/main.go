package main

import (
	"net/http"
	"github.com/mlapeyre/go-test/goFileServer"
	"log"
)


func main() {
	var storageHandler * goFileServer.StorageHandlerDefinition = goFileServer.New("/browse/", "/tmp/martial")
	http.HandleFunc(storageHandler.UrlBase, storageHandler.CreateHandler())
	error := http.ListenAndServe(":8080", nil)
	if(error!=nil){
		log.Fatal(error)
	}
}
