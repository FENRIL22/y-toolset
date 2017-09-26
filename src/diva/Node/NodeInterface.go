package node

import (
)


/* == Interface ===================================================== */
type Node interface{
	Id() int64					//GetID
	Type() string				//GetType
	Save() bool					//Save to Datastore
	Get() (interface{}, string)	//Get Primary Section and Description
	String() string
}

/* == Struct ======================================================== */
type NodeBase struct {
	id int64
	dataType string;
}

func (s *NodeBase) Id() int64 {
	return s.id
}

func (s *NodeBase) Type() string {
	return s.dataType;
}

/*
//virtual method
func NewAnyStruct(need data) (*AnyStruct)
func (s *NodeAny) init() (void)
func (s *NodeAny) String() (string)
*/
