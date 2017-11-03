package node

import (
	"fmt"
)


/* == Struct ======================================================== */
type NodeText struct {
	//*NodeBase
	NodeBase
}

func NewNodeText() *NodeText{
	s := new(NodeText)
	//s.NodeBase = new(NodeBase)

	s.init()

	return s
}


func (s *NodeText) init() {
	s.Title = "Text node desu"
	s.DataType = "NodeText"
}

func (s *NodeText) SetText(str string) {
}

func (s *NodeText) String() string {
	return fmt.Sprintln(s.NodeBase)
}
