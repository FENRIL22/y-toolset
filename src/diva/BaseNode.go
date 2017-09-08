package diva

import (
)


/* == Interface ===================================================== */
type Node interface{
	Id() int					//GetID
	Type() string				//GetType
	Save() bool					//Save to Datastore
	Get() (interface{}, string)	//Get Primary Section and Description
	Display()					//func for debug
}

/* == Struct ======================================================== */
type BaseNode struct {
	id int
}

func (n *BaseNode) Id() int {
	return n.id
}

/*
//virtual method
func (n *BaseNode) Type() string {}
func (n *BaseNode) Display() {}
*/
