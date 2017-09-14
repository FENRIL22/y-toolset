package diva

import (
	"fmt"
	"time"
	"net/http"
	"errors"
	//"context"
	"google.golang.org/appengine"
	//"appengine"
	//"appengine/datastore"
	"google.golang.org/appengine/datastore"
	//"strconv"
)

/* == Interface ===================================================== */
type Memo interface {
	Get()
	Put()
	Search()
	GetRelation()
}

/* == Struct ======================================================== */
type MemoManager struct{
	//ctx interface{}
}

func (m *MemoManager) init(r *http.Request) error{
	return errors.New("(*>△<)<ナーンナーンっっ")
}

func NewMemoManager(r *http.Request) *MemoManager {
	m := new(MemoManager)

	_ = m.init(r)

	return m
}

//func MemoManagerHandler(w http.ResponseWriter, r *http.Request){
//	fmt.Fprint(w, "hello diva!")
//}



/* == Testing ======================================================= */


type Employee struct{
	Name string
	Role string
	HireDate time.Time
}


func Add(w http.ResponseWriter, r *http.Request){
	ctx := appengine.NewContext(r)
	e1 := Employee{
		Name: "Joe Citizen",
		Role: "Manager",
		HireDate: time.Now(),
		//Account: user.Current(c).String(),
	}
	//put
	//key := datastore.NewIncompleteKey(ctx, "employee", nil)
	key := datastore.NewKey(ctx, "employee", "", 0, nil)

	key, _ = datastore.Put(ctx, key, &e1)

	fmt.Fprintf(w, "hogehoge\n")
	//fmt.Fprintf(w, strconv.FormatInt(key.IntID(), 10))
	fmt.Fprintf(w, "foo : %v\n",  key.IntID())
	//fmt.Fprintf(w, key.Kind())
	//fmt.Fprintf(w, key.StringID())
}

func Get(w http.ResponseWriter, r *http.Request){
	ctx := appengine.NewContext(r)

	//q := datastore.NewQuery("employee").Order("-Name")
	q := datastore.NewQuery("employee")
	//q := datastore.NewQuery("employee").KeysOnly()

	var ed Employee
	var ep []Employee
	keys, _ := q.GetAll(ctx, &ep)


	//get

	fmt.Fprintln(w, "hoge", "figa")
	fmt.Fprintln(w, keys)
	_ = datastore.Get(ctx, keys[0], &ed)
	fmt.Fprintln(w, ep)
	fmt.Fprintln(w, ed)

	iter := q.Run(ctx)
	for{
		var d Employee
		_, err := iter.Next(&d)

		if err == datastore.Done{
			break
		}

		fmt.Fprintln(w, "Data: ", d)
	}


}

func Update(r *http.Request) Employee{
	ctx := appengine.NewContext(r)
	//update
	e2 := Employee{
		Name: "Puma Nyancat",
		Role: "Manager",
		HireDate: time.Now(),
		//Account: user.Current(c).String(),
	}

	//key := datastore.NewIncompleteKey(ctx, "employee", nil)
	key := datastore.NewKey(ctx, "employee", "", 4923613069180928, nil)

	_, _ = datastore.Put(ctx, key, &e2)
	return e2
}

func Delete(w http.ResponseWriter, r *http.Request){
	ctx := appengine.NewContext(r)
	var ee []Employee
	////delete
	//key := datastore.NewIncompleteKey(ctx, "employee", nil)
	//_ = datastore.Delete(ctx, key)
	q := datastore.NewQuery("employee").Filter("Name =", "Joe Citizen").KeysOnly()
	keys, _ := q.GetAll(ctx, &ee)

	for _, k := range keys{
		_ = datastore.Delete(ctx, k)
	}
}
/*
*/
