package tasker

import (
	"fmt"
	//"html"
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
	//ret := s.r.FormValue("post_value")
	_ = s.r.FormValue("post_value")
	////a := html.EscapeString(ret)

	//s.Return(fmt.Sprintln(ret))
	s.Return(fmt.Sprintf("%#v\n", s.r.Body))
}
