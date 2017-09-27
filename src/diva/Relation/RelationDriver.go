package relation

import (
	//"fmt"
	//"time"
	//"net/http"
	//"errors"
	"context"
	//"google.golang.org/appengine"
	//"google.golang.org/appengine/datastore"
	//"strconv"
)

/* == Struct ======================================================== */
type RelationDriver struct{
	context context.Context
	databus chan []Relation
}

func NewRelationDriver(ctx context.Context) *RelationDriver {
	s := new(RelationDriver)

	s.init(ctx)

	return s
}

func (s *RelationDriver) init(ctx context.Context) {
	s.databus = make(chan []Relation)
	s.context = ctx
}

func (s *RelationDriver) deinit() {
	//deinit
	//sync datastore
	//delete datastore connection
}

func (s *RelationDriver) Put(re []Relation) error {
	return nil
}

func (s *RelationDriver) Update(re []Relation) error {
}

func (s *RelationDriver) Get(list []int64) ([]Relation, error) {
	return []Relation{}, nil
}
