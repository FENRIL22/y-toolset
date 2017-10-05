package node

import (
)


/* == Struct ======================================================== */
type NodeText struct {
	*NodeBase
}

func NewNodeText() *NodeText{
	s := new(NodeText)
	s.NodeBase = new(NodeBase)

	s.init()

	return s
}

func (s *NodeText) init() {
	s.Title = "Text node desu"
	s.dataType = "NodeText"
}

func (s *NodeText) String() string {
	return "implementing Now... : NodeURL"
}
