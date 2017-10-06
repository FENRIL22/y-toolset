package node

import (
	//"fmt"
	//"time"
	//"net/http"
	//"errors"
	"context"
	//"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	//"strconv"
)

/* == Struct ======================================================== */
type NodeStorage struct{
	context context.Context
}

func NewNodeStorage(ctx context.Context) *NodeStorage {
	s := new(NodeStorage)

	s.init(ctx)

	return s
}

func (s *NodeStorage) init(ctx context.Context) {
	s.context = ctx
}

func (s *NodeStorage) FindByType(dataType string) {

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

func (s *NodeStorage) Delete(){
	//TODO:name ->deleteNode?
}

func (s *NodeStorage) Flush(){
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
