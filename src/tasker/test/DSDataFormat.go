package dsdriver

import (
	"errors"
)

type DSFormat struct {
	internal map[string]interface{}
}

func NewDSFormat() *DSFromat {
	s := new(DS)

	return s
}

func (s *DSFormat) Add(k string, v interface{}) {
	s.internal[k] = v
}

func (s *DSFormat) Value(k string) (interface{}, error) {
	v, ok := s.internal[k]

	if !ok {
		return nil, errors.New("NotFound")
	}
	return v, nil
}
