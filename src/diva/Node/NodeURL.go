package node

import (
)


/* == Struct ======================================================== */
type NodeURL struct {
	*NodeBase
}

func NewNodeURL() *NodeURL{
	s := new(NodeURL)
	s.NodeBase = new(NodeBase)

	s.init()

	return s
}

func (s *NodeURL) init() {
	s.dataType = "NodeURL"
	s.Title = "korehaURL"
}

func (s *NodeURL) String() string {
	return "implementing Now... : NodeURL"
}
