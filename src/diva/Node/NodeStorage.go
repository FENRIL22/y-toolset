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

func (s *NodeStorage) FindByName(dataName string) {
	//TODO:name ->FindByFuzzyName?
}

func (s *NodeStorage) New(dataType string) interface{}{
	//TODO:name ->newNode/Resister?

	var ndi Node

	//switch dataType {
	//case "URL":
	//	ndi = NewNodeURL()
	//case "Text":
	//	ndi := NewNodeText()
	////case "audio"//
	//default: break
	//}

	ndi = NewNodeURL()

	var i6 int64 = 12321
	ndi.SetId(i6)

	key := datastore.NewKey(s.context, "Node", "", 0, nil)
	keyID, _ := datastore.Put(s.context, key, *ndi)
	//q := datastore.NewQuery("Node").Filter("id =", 12321)
	//key, _ = q.GetAll(s.context, ndi)

	//return key
	return keyID.IntID()
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
}
*/
