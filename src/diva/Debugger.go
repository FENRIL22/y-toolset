package diva

import (
	"fmt"
	//"time"
	"net/http"
	//"errors"
	"context"
	"google.golang.org/appengine"
	//"google.golang.org/appengine/datastore"
	////"strconv"
)

type DivaContext struct{
	w http.ResponseWriter
	r *http.Request
	ctx context.Context
	data chan int
	Ret chan interface{}
}

func (s *DivaContext) init(w http.ResponseWriter, r *http.Request){
	s.w = w
	s.r = r
	s.ctx = appengine.NewContext(r)
	s.data = make(chan int)
	s.Ret = make(chan interface{}, 10)
}

func NewDivaContext(w http.ResponseWriter, r *http.Request) *DivaContext{
	s := new(DivaContext)
	s.init(w, r)

	return s
}

/* == Struct ======================================================== */
type Debugger struct{
	dctx *DivaContext
	//ctx interface{}
}

func (s *Debugger) init(dc *DivaContext) {
	s.dctx = dc
}

func NewDebugger(w http.ResponseWriter, r *http.Request) (*Debugger, *DivaContext) {
	s := new(Debugger)
	dc := NewDivaContext(w, r)

	s.init(dc)

	return s, dc
}

func (s * Debugger) Call() {
	fmt.Fprintln(s.dctx.w, "teste")
	s.dctx.Ret <- "hogefuga"

}
