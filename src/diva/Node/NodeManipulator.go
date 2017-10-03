package node
//TODO:Delete It?

import (
	//"fmt"
	//"time"
	//"net/http"
	//"errors"
	//"context"
	//"google.golang.org/appengine"
	//"google.golang.org/appengine/datastore"
	//"strconv"
)

/* NodeManager
 * DataStoreとMemoManagerを仲介する *
 * ドライバとして利用               */
/* == Struct ======================================================== */
type NodeManager struct{
	*nodeDriver NodeDriver
}

func NewNodeManager(nd *NodeDriver) *NodeManager {
	s := new(NodeManager)

	s.init(nd NodeDriver)

	return s
}

func (s *NodeManager) init(nd *NodeDriver) {
	s.nodeDriver = nd
}

func (s *NodeManager) Get(nd Node) {
}

func (s *NodeManager) Put() {
}

func (s *NodeManager) Create(){
}

func (s *NodeManager) Delete(){
}
