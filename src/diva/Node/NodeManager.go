package node

import (
	//"fmt"
	//"time"
	//"net/http"
	//"errors"
	"context"
	//"google.golang.org/appengine"
	//"google.golang.org/appengine/datastore"
	//"strconv"
)

/* == Struct ======================================================== */
type NodeManager struct{
	//Delete it?
	*nodeMpAry []NodeManipulator
	context context.Context
}

func NewNodeManager(ctx context.Context) *NodeManager {
	s := new(NodeManager)

	s.init(ctx)

	return s
}

func (s *NodeManager) init(ctx context.Context) {
	s.context = ctx
	s.nodeMpAry = new([]NodeManipulator, 0, 5)
}

func (s *NodeManager) Search(){}

func (s *NodeManager) Put() {
}

func (s *NodeManager) Create(){
}

func (s *NodeManager) Delete(){
}
