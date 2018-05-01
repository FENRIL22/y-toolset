package tasker

import (
	"fmt"
	//"html"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Tasker struct {
	w http.ResponseWriter
	r *http.Request
}

func NewTasker(w http.ResponseWriter, r *http.Request) *Tasker {
	s := new(Tasker)
	s.w = w
	s.r = r

	return s
}

func (s *Tasker) Return(str string) {
	s.w.WriteHeader(200)
	s.w.Header().Set("Content-Type", "text/html; charset=utf8")
	//s.w.Write([]byte("こんにちは"))
	//fmt.Fprintln(s.w, str)
	s.w.Write([]byte(str))
}

func (s *Tasker) Test() {
	var dat map[string]interface{}
	bt, _ := ioutil.ReadAll(s.r.Body)
	err := json.Unmarshal(bt, &dat)
	if err != nil {
		s.Return(err.Error())
	}
	//json.NewDecoder(s.r.Body).Decode(dat)
	s.Return(fmt.Sprintln(dat))
}
