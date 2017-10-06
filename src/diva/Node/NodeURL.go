package node

import (
	"fmt"
)


/* == Struct ======================================================== */
type NodeURL struct {
	NodeBase
	URL string
	TEST string
}

func NewNodeURL() *NodeURL{
	s := new(NodeURL)

	s.init()

	return s
}

func (s *NodeURL) init() {
	s.DataType = "NodeURL"
	s.Title = "korehaURL"
	s.TEST = "DO test."
}

func (s *NodeURL) SetURL(u string){
	s.URL = u
}

func (s *NodeURL) String() string {
	return fmt.Sprintln(s.NodeBase, s.URL, s.TEST, s.ID)
}
