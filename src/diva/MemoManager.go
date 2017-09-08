package diva

import (
	//"fmt"
	"time"
	"net/http"
	//"context"
	"google.golang.org/appengine"
	//"appengine"
	//"appengine/datastore"
	"google.golang.org/appengine/datastore"
)

/* == Interface ===================================================== */
type Memo interface {
	Get()
	Put()
}

/* == Struct ======================================================== */
type MemoManager struct{
	//ctx interface{}
}

//func (m *MemoManager) init(r *http.Request) err{
//	ctx := appengine.Newcontext(r)
//
//	return nil
//}

func NewMemoManager(r *http.Request) *MemoManager {
	m := new(MemoManager)

	//m.init(r)

	return m
}

//func MemoManagerHandler(w http.ResponseWriter, r *http.Request){
//	fmt.Fprint(w, "hello diva!")
//}




type Employee struct{
	Name string
	Role string
	HireDate time.Time
}


func Add(r *http.Request){
	ctx := appengine.NewContext(r)
	e1 := Employee{
		Name: "Joe Citizen",
		Role: "Manager",
		HireDate: time.Now(),
		//Account: user.Current(c).String(),
	}
	//put
	key := datastore.NewIncompleteKey(ctx, "employee", nil)
	key, _ = datastore.Put(ctx, key, &e1)
}

func Get(r *http.Request) Employee{
	ctx := appengine.NewContext(r)
	//get
	key := datastore.NewIncompleteKey(ctx, "employee", nil)
	var e2 Employee
	_ = datastore.Get(ctx, key , &e2)
	return e2
}

func Update(r *http.Request) Employee{
	ctx := appengine.NewContext(r)
	//update
	e2 := Employee{
		Name: "Joe Nyancat",
		Role: "Manager",
		HireDate: time.Now(),
		//Account: user.Current(c).String(),
	}

	//key := datastore.NewIncompleteKey(ctx, "employee", nil)
	key := datastore.NewKey(ctx, "employee", "", 5311740673785856, nil)

	_, _ = datastore.Put(ctx, key, &e2)
	return e2
}

func Delete(r *http.Request){
	ctx := appengine.NewContext(r)
	////delete
	key := datastore.NewIncompleteKey(ctx, "employee", nil)
	_ = datastore.Delete(ctx, key)
}
/*
*/
