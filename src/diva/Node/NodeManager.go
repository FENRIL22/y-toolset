package node

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
	DataList []Node
	//ctx interface{}
}

func (s *NodeManager) init() {
}

func NewNodeManager() *NodeManager {
	s := new(NodeManager)

	s.init()

	return s
}
