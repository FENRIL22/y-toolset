package node

import (
)


/* == Interface ===================================================== */
type Node interface{
	GetID() int64					//GetID
	GetType() string				//GetType
	SetID(int64)
	SetTitle(string)
	SetDescription(string)
	//SetId()
	//Save() bool					//Save to Datastore
	//Get() (interface{}, string)	//Get Primary Section and Description
	//String() string
}

/* == Struct ======================================================== */
type NodeBase struct {
	ID int64
	DataType string
	Title string
	Description string
}

func (s *NodeBase) GetID() int64 {
	return s.ID
}

func (s *NodeBase) GetType() string {
	return s.DataType
}

func (s *NodeBase) SetTitle(str string) {
	s.Title = str
}

func (s *NodeBase) SetDescription(str string) {
	s.Description = str
}

func (s *NodeBase) SetID(inid int64) {
	s.ID = inid
}

/*
//virtual method
func NewAnyStruct(need data) (*AnyStruct)
func (s *NodeAny) init() (void)
func (s *NodeAny) String() (string)
*/
