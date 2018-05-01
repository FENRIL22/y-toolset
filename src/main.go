package hello

import (
	//"golang.org/x/net/context"
	//"google.golang.org/appengine"
	//"google.golang.org/appengine/datastore"
	//"time"
	//for test
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"lib/diva"
	"lib/sluck"
	"lib/tasker"
	"net/http"
)

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/hoge", hogeHandler)
	r.HandleFunc("/urldisas/{id}", urlDisassemblyHandler)
	r.HandleFunc("/exp/get", getHandler)
	r.HandleFunc("/exp/update", updateHandler)
	r.HandleFunc("/exp/delete", deleteHandler)
	r.HandleFunc("/exp/jsonping", jsonPingHandler)
	r.HandleFunc("/exp/add", debugAddHandler)
	r.HandleFunc("/exp/call", debugCallHandler)
	r.HandleFunc("/api/sluck/endpoint/{response}", sluckHandler)
	r.HandleFunc("/tasker/test", taskerTestHandler)
	http.Handle("/", r)
}

func hogeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hogefuga")
}

func taskerTestHandler(w http.ResponseWriter, r *http.Request) {
	t := tasker.NewTasker(w, r)
	t.Test()
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
	fmt.Fprintln(w, id)
}

func debugAddHandler(w http.ResponseWriter, r *http.Request) {
	nd, _ := diva.NewDebugger(w, r)
	nd.Add()
}

func debugCallHandler(w http.ResponseWriter, r *http.Request) {
	nd, _ := diva.NewDebugger(w, r)
	nd.Call()
}

type Ping struct {
	status int    `json:"status"`
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
	if responsetype == "eventsub" {
		sl.EventAPI()
	} else if responsetype == "interactivereq" {
		//sl.InteractiveAPI()
		sluck.ResponseToRequest(w, r)
	} else {
		fmt.Fprintln(w, "hoge")
	}
}
