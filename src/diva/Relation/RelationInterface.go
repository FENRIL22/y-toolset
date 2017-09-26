package relation

import (
	//"fmt"
)

/* == Interface ===================================================== */
type Relation interface{
	Id() int64					//GetID
	Type() string				//GetType
	Get()
	Set()
	//Save() bool					//Save to Datastore
	//Get() (interface{}, string)	//Get Primary Section and Description
	String() string
}

/* == Struct ======================================================== */
type RelationBase struct {
	id int64
	dataType string
	//src int64
	//dst int64
	// some property
}

func (s *RelationBase) Id() int64 {
	return s.id
}

func (s *RelationBase) Type() string {
	return s.dataType
}

//func (s *RelationBase) String() string {
//	var retStr string
//
//	retStr = fmt.Sprintln("Do Test.")
//	retStr += fmt.Sprintln("Cake is a lie.")
//
//	return retStr
//}


/*
//virtual method
func NewAnyStruct(need data) (*AnyStruct)
func (s *NodeAny) init() (void)
func (s *NodeAny) String() (string)
*/

