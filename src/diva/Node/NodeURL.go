package node

import (
)


/* == Struct ======================================================== */
type NodeURL struct {
	*NodeBase
}

func NewNodeURL(id int64) *NodeURL{
	s := new(NodeURL)
	s.NodeBase = new(NodeBase)

	s.init(id)

	return s
}

func (s *NodeURL) init(id int64) {
	s.id = id
	s.dataType = "NodeURL"
}

func (s *NodeURL) String() string {
	return "implementing Now... : NodeURL"
}
