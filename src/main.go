package hello

import (
	"fmt"
	//"golang.org/x/net/context"
	//"golang.com/appengine"
	//"appengine"
	//"golang.com/appengine/datastore"
	"net/http"
	"lib/diva"
)


func init(){
	http.HandleFunc("/", hogehandler)
	http.HandleFunc("/add", addhandler)
	http.HandleFunc("/update", updatehandler)
	http.HandleFunc("/delete", deletehandler)
}


func hogehandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hogefuga")
}

func addhandler(w http.ResponseWriter, r *http.Request) {
	diva.Add(r)
}

func updatehandler(w http.ResponseWriter, r *http.Request) {
	diva.Update(r)
}

func deletehandler(w http.ResponseWriter, r *http.Request) {
	diva.Delete(r)
}
