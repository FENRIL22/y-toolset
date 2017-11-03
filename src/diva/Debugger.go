package diva

import (
	"fmt"
	//"time"
	"net/http"
	//"errors"
	"context"
	"google.golang.org/appengine"
	//"google.golang.org/appengine/datastore"
	"lib/diva/Node"
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

func (s *Debugger) Add() {
	//ts := node.NewNodeURL()
	//ts.SetURL("TestbyTest.com")
	ts := node.NewNodeText()
	ts.SetDescription("HogeHogeHOge")
	ds := node.NewNodeStorage(s.dctx.ctx)
	ds.Save(ts)
}

func (s * Debugger) Call() {
	ds := node.NewNodeStorage(s.dctx.ctx)
	data := ds.Find("NodeURL").Title("korehaURL").Description("tete").Get()

	fmt.Fprintln(s.dctx.w, data)

}
//Delete
//ts.SetID(4960996464525312)
//err := ds.Delete(ts)
//fmt.Fprintln(s.dctx.w, hoge, err)

////EDIT
//for _,v := range(hoge) {
////edit by pointer
//	//fmt.Fprintln(s.dctx.w, v.SetTitle("Fugaa"))
//	v.SetTitle("gaa")
//}
//
//ds.SaveMulti(hoge)
