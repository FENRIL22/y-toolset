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

/* == Struct ======================================================== */
type NodeStorage struct{
	context context.Context
	dataTypeList []string
}

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

func (s *NodeStorage) FindByType(dataType string) {

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

func (s *NodeStorage) Register(ndi Node) {
	//TODO:name ->newNode/New?
	key := datastore.NewKey(s.context, ndi.GetType(), "", 0, nil)
	k, _ := datastore.Put(s.context, key, ndi)
	ndi.SetID(k.IntID())
	key = datastore.NewKey(s.context, ndi.GetType(), "", k.IntID(), nil)
	_, _ = datastore.Put(s.context, key, ndi)
}

func (s *NodeStorage) GetByID (id int64) Node{
	key := datastore.NewKey(s.context, "NodeURL", "", id, nil)
	var n Node
	switch key.Kind(){
	case "NodeURL":
		n = NewNodeURL()
	case "NodeText":
		n = NewNodeText()
	}
	datastore.Get(s.context, key, n)

	return n
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

func (s *NodeStorage) Flush(){
}

func (s *NodeStorage) AllocNode(name string) Node {
	var nd Node

	if name == "NodeURL" {
		nd = NewNodeURL()
	} else if name == "NodeText" {
		nd = NewNodeText()
	}

	if nd == nil {
		return nil
	}

	return nd
}

func (s *NodeStorage) AllocNodes(name string, length int) []Node {
	var nds []Node

	for i:=0;i<length;i++ {
		nd := s.AllocNode(name)
		nds = append(nds, nd)
	}

	return nds
}

func (s *NodeStorage) GetByKey(key *datastore.Key) {
}

func (s *NodeStorage) GetByKeys(keys *[]datastore.Key) {
}
/*
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
*/
