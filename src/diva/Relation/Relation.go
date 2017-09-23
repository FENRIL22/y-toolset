package diva

import (
	"fmt"
)

/* == Interface ===================================================== */
type Relation interface{
	Id() int					//GetID
	Type() string				//GetType
	Get()
	Set()
	//Save() bool					//Save to Datastore
	//Get() (interface{}, string)	//Get Primary Section and Description
	//String() string
	//Display()					//func for debug
	//TODO:RelationManager
	//Sync()						//Sync With Datastore
	//Pull()						//Pull from Datastore
}

/* == Struct ======================================================== */
type RelationVector struct {
	src int64
	dst int64
	// some property
}

func (s *RelationVector) Type() string {
	return "Relation"
}

func (s *RelationVector) Display() string {
	return "hoge"
}

func (s *RelationVector) String() string {
	var retStr string

	retStr = fmt.Sprintln("Do Test.")
	retStr += fmt.Sprintln("Cake is a lie.")

	return retStr
}


/*
//virtual method
func (n *BaseNode) Type() string {}
func (n *BaseNode) Display() {}
*/
