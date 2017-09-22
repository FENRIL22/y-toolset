package diva

import (
)


/* == Interface ===================================================== */
type Node interface{
	Id() int					//GetID
	Type() string				//GetType
	Save() bool					//Save to Datastore
	Get() (interface{}, string)	//Get Primary Section and Description
	String() string
	//Display()					//func for debug
}

/* == Struct ======================================================== */
type BaseNode struct {
	id int
	dataType string;
}

func (s *BaseNode) Id() int {
	return s.id
}

/*
//virtual method
func (s *BaseNode) Type() (string)
func (s *BaseNode) String() (string)
func (s *BaseNode) Display() (AnyData)
func NewAnyStruct(need data) (*AnyStruct)
*/
