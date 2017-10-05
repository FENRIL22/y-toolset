//TODO:futurework
// Driver to Datastore/SQL
// 
/* 
INPUT:
Driver(Datastore/SQL)


*/
package node
//
//import (
//	//"fmt"
//	//"time"
//	//"net/http"
//	"errors"
//	"context"
//	//"google.golang.org/appengine"
//	"google.golang.org/appengine/datastore"
//	//"strconv"
//)
//
///* == Struct ======================================================== */
//type NodeDriver struct{
//	context context.Context
//}
//
//func NewNodeDriver(ctx context.Context) *NodeDriver {
//	s := new(NodeDriver)
//
//	s.init(ctx)
//
//	return s
//}
//
//func (s *NodeDriver) init(ctx context.Context) {
//	s.context = ctx
//}
//
//func (s *NodeDriver) Put(nd Node) error {
//	_ = datastore.NewKey(s.context, "Node", "", 0, nil)
//	return nil
//}
//
//func (s *NodeDriver) PutMulti(nd []Node) error {
//	//key := datastore.NewKey(s.context, "Relation", "", 0, nil)
//	_ = datastore.NewKey(s.context, "Node", "", 0, nil)
//
//	return nil
//}
//
//func (s *NodeDriver) Get(id int64) (Node, error) {
//	ctx := s.context
//	var nd []Node
//
//	q := datastore.NewQuery("Node").Filter("id =", id)
//	q.GetAll(ctx, &nd)
//
//	if len(nd) < 1 {
//		return nil, errors.New("Data Not Found")
//	} else if len(nd) > 1 {
//		return nil, errors.New("Unknown Error")
//	}
//
//	return nd[0], nil
//}
//
//func (s *NodeDriver) GetMulti(ids []int64) ([]Node, error) {
//	return []Node{}, nil
//}
//
//func (s *NodeDriver) Remove(){}
//func (s *NodeDriver) RemoveMulti(){}
//
