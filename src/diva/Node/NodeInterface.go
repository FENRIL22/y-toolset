package node

import (
)


/* == Interface ===================================================== */
type Node interface{
	Id() int64					//GetID
	Type() string				//GetType
	SetId(int64)
	SetTitle(string)
	SetDescription(string)
	//SetId()
	//Save() bool					//Save to Datastore
	//Get() (interface{}, string)	//Get Primary Section and Description
	//String() string
}

/* == Struct ======================================================== */
type NodeBase struct {
	id int64
	dataType string
	Title string
	Description string
}

func (s *NodeBase) Id() int64 {
	return s.id
}

func (s *NodeBase) Type() string {
	return s.dataType
}

func (s *NodeBase) SetTitle(str string) {
	s.Title = str
}

func (s *NodeBase) SetDescription(str string) {
	s.Description = str
}

func (s *NodeBase) SetId(inid int64) {
	s.id = inid
}

/*
//virtual method
func NewAnyStruct(need data) (*AnyStruct)
func (s *NodeAny) init() (void)
func (s *NodeAny) String() (string)
*/
