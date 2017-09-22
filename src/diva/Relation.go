package diva

import (
	"fmt"
)


/* == Struct ======================================================== */
type Relation struct {
	BaseNode

	// some property
}

func (s *Relation) Type() string {
	return "Relation"
}

func (s *Relation) Display() string {
	return "hoge"
}

func (s *Relation) String() string {
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
