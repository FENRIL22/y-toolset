package hello

import (
	//"golang.org/x/net/context"
	//"google.golang.org/appengine"
	//"google.golang.org/appengine/datastore"
	//"time"
	//for test
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"lib/diva"
	"lib/sluck"
)


func init(){
	r := mux.NewRouter()
	r.HandleFunc("/hoge", hogeHandler)
	r.HandleFunc("/urldisas/{id}", urlDisassemblyHandler)
	r.HandleFunc("/add", addHandler)
	r.HandleFunc("/get", getHandler)
	r.HandleFunc("/update", updateHandler)
	r.HandleFunc("/delete", deleteHandler)
	r.HandleFunc("/jsonping", jsonPingHandler)
	r.HandleFunc("/api/sluck/endpoint/{response}", sluckHandler)
	http.Handle("/", r)
}


func hogeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hogefuga")
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	diva.Get(w, r)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	diva.Add(w, r)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	diva.Update(r)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	diva.Delete(w, r)
}

func urlDisassemblyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintln(w, id);
}

type Ping struct {
    status int `json:"status"`
    Result string `json:"result"`
}

func jsonPingHandler(w http.ResponseWriter, r *http.Request) {
     ping := Ping{http.StatusOK, "ok"}

     res, err := json.Marshal(ping)

     if err != nil {
         http.Error(w, err.Error(), http.StatusInternalServerError)
         return
     }

     w.Header().Set("Content-Type", "application/json")
     w.Write(res)
 }

 func sluckHandler(w http.ResponseWriter, r *http.Request) {
	 vars := mux.Vars(r)
	 responsetype := vars["response"]

	 sl := sluck.NewSluck(w, r)

	 //fmt.Fprintln(w, sl.Hoge())
	 if( responsetype == "eventsub" ){
		 sl.EventAPI()
	 }else if( responsetype == "interactivereq"){
		//sl.InteractiveAPI()
		sluck.ResponseToRequest(w, r)
	 }else{
		 fmt.Fprintln(w, "hoge")
	 }
 }
