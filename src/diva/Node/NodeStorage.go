package node

import (
	//"fmt"
	//"time"
	//"net/http"
	"errors"
	"context"
	//"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	//"strconv"
)

/* == Struct -finder- ============================================== */
type NodeStorage struct{
	context context.Context
	dataTypeList []string
}

type NodeFinder struct{
	Node []Node
}

func (s *NodeFinder) Get() []Node {
	return s.Node
}

func (s *NodeFinder) Title(title string) *NodeFinder {
	var nl []Node

	for _,v := range(s.Node) {
		if v.GetTitle() == title {
			nl = append(nl, v)
		}
	}

	s.Node = nl

	return s
}

//Impliment
//func (s *NodeFinder) FuzzyTitle(title string) *NodeFinder {
//}

func (s *NodeFinder) Description(desc string) *NodeFinder {
	var nl []Node

	for _,v := range(s.Node) {
		if v.GetDescription() == desc {
			nl = append(nl, v)
		}
	}

	s.Node = nl

	return s
}

//Impliment
//func (s *NodeFinder) FuzzyDescription(title string) *NodeFinder {
//}

/* == Struct -Main- ================================================= */
func NewNodeStorage(ctx context.Context) *NodeStorage {
	s := new(NodeStorage)

	s.init(ctx)

	return s
}

func (s *NodeStorage) init(ctx context.Context) {
	s.context = ctx
	s.dataTypeList = []string{
		"NodeURL",
		"NodeText",
	}
}

/* create NodeFinder */
func (s *NodeStorage) Find(dataType string) *NodeFinder {
	nf := new(NodeFinder)

	switch dataType {
	case "NodeURL":
		nf.Node = s.getNodeURL()
	case "NodeText":
		nf.Node = s.getNodeText()
	}

	return nf
}

func (s *NodeStorage) FindByType(dataType string) []Node {
	var ndl []Node

	switch dataType {
	case "NodeURL":
		ndl = s.getNodeURL()
	case "NodeText":
		ndl = s.getNodeText()
	}

	return ndl
}

func (s *NodeStorage) FindByTitle(title string) interface{}{
	var ndl []Node

	//q := datastore.NewQuery("NodeURL").Filter("Title =", title)

	//keys, _ := q.GetAll(s.context, &ndl)
	//for _, v := range(s.dataTypeList) {
	//	q := datastore.NewQuery(v).Filter("Title =", title)

	//	keys, _ := q.GetAll(s.context, &ndl)

	//	if len(keys) > 0 {
	//		return ndl
	//	}
	//}


	return ndl

	//key := datastore.NewQuery()

}


func (s *NodeStorage) FindByName(dataName string, searchType string) {
	//TODO:name ->FindByFuzzyName?
	if searchType == "fuzzy" {
	} else if searchType == "strict" {
	}
}

func (s *NodeStorage) SaveMulti(node []Node) {
	for _, v := range(node) {
		s.Save(v)
	}
}

func (s *NodeStorage) Save(ndi Node) {
	//TODO:name ->newNode/New?
	if ndi.GetID() == 0 {
		key := datastore.NewIncompleteKey(s.context, ndi.GetType(), nil)
		k, _ := datastore.Put(s.context, key, ndi)

		id := k.IntID()
		datastore.Delete(s.context, key)
		datastore.Delete(s.context, k)

		ndi.SetID(id)
		key = datastore.NewKey(s.context, ndi.GetType(), "", id, nil)
		_, _ = datastore.Put(s.context, key, ndi)
	} else {
		key := datastore.NewKey(s.context, ndi.GetType(), "", ndi.GetID(), nil)
		_, _ = datastore.Put(s.context, key, ndi)
	}
}

func (s *NodeStorage) Delete(nd Node) error {
	//TODO:name ->deleteNode?
	var ee Node
	q := datastore.NewQuery(nd.GetType()).Filter("ID =", nd.GetID()).KeysOnly()
	//q := datastore.NewQuery(nd.GetType()).Filter("Title =", "korehaURL").KeysOnly()
	keys, err := q.GetAll(s.context, &ee)

	if err != nil {
		return errors.New("Datastore Error")
	}

	if len(keys) < 1 {
		return errors.New("NotFound")
	}

	err = datastore.DeleteMulti(s.context, keys)

	if err != nil {
		return errors.New("Datastore Error:delete")
	}

	return nil
}

// -- low area --
func (s *NodeStorage) UpCast(nodes []interface{}) []Node {
	var node Node
	var nodel []Node

	for _, v:= range(nodes) {
		node = v.(Node)
		nodel = append(nodel, node)
	}

	return nodel
}

func (s *NodeStorage) getNodeURL() []Node {
	q := datastore.NewQuery("NodeURL")

	iter := q.Run(s.context)

	var nodel []Node
	n := new(NodeURL)
	var nt Node
	for{
		n = new(NodeURL)
		_, err := iter.Next(n)
		nt = n

		if err == datastore.Done{
			break
		}

		nodel = append(nodel, nt)
	}


	return nodel
}

func (s *NodeStorage) getNodeText() []Node {
	q := datastore.NewQuery("NodeText")

	iter := q.Run(s.context)

	var nodel []Node
	n := new(NodeText)
	var nt Node
	for{
		n = new(NodeText)
		_, err := iter.Next(n)
		nt = n

		if err == datastore.Done{
			break
		}

		nodel = append(nodel, nt)
	}


	return nodel
}

//func (s *NodeStorage) AllocNode(name string) Node {
//	var nd Node
//
//	if name == "NodeURL" {
//		nd = NewNodeURL()
//	} else if name == "NodeText" {
//		nd = NewNodeText()
//	}
//
//	if nd == nil {
//		return nil
//	}
//
//	return nd
//}
//

//func (s *NodeStorage) GetByID (id int64) Node{
//	key := datastore.NewKey(s.context, "NodeURL", "", id, nil)
//
//	n := s.AllocNode("NodeURL")
//
//	datastore.Get(s.context, key, n)
//
//	return n
//}

/*
func (s *NodeStorage) Flush(){
}

func (s *NodeStorage) GetByKey(key *datastore.Key) {
}

func (s *NodeStorage) GetByKeys(keys *[]datastore.Key) {
}

func (s *NodeManager) UpCast(nodegl interface{}) Node {
	//TODO:Need?
}

func (s *NodeManager) DownCast(nd Node) interface{} {
	//TODO:Need?
	//var i6 int64 = 12321
	//nd.SetId(i6)

	//nd := ndi.(*NodeURL)

	//switch v := ndi.(type) {
	//case NodeURL:
	//	var nd *NodeURL
	//	nd = ndi.(*NodeURL)
	//case NodeText:
	//	var nd *NodeText
	//	nd = ndi.(*NodeText)
	//default:
	//	nd := NewNodeText()
	//}
}
func (s *NodeStorage) 0ldgetNodeURL() []Node {
	q := datastore.NewQuery("NodeURL")

	var ndl []NodeURL
	_, _ = q.GetAll(s.context, &ndl)

	var node Node
	var nodel []Node

	for _, v:= range(ndl) {
		node = v
		nodel = append(nodel, node)
	}

	return nodel
}
*/
