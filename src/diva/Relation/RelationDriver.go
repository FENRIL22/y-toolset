package relation

import (
	//"fmt"
	//"time"
	//"net/http"
	"errors"
	"context"
	//"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	//"strconv"
)

/* == Struct ======================================================== */
type RelationDriver struct{
	context context.Context
}

func NewRelationDriver(ctx context.Context) *RelationDriver {
	s := new(RelationDriver)

	s.init(ctx)

	return s
}

func (s *RelationDriver) init(ctx context.Context) {
	s.context = ctx
}

func (s *RelationDriver) deinit() {
}

func (s *RelationDriver) Put(re Relation) error {
	_ = datastore.NewKey(s.context, "Relation", "", 0, nil)
	return nil
}

func (s *RelationDriver) PutMulti(re []Relation) error {
	//key := datastore.NewKey(s.context, "Relation", "", 0, nil)
	_ = datastore.NewKey(s.context, "Relation", "", 0, nil)

	return nil
}

func (s *RelationDriver) Get(id int64) (Relation, error) {
	ctx := s.context
	var rl []Relation

	q := datastore.NewQuery("Relation").Filter("id =", id)
	q.GetAll(ctx, &rl)

	if len(rl) < 1 {
		return nil, errors.New("Data Not Found")
	} else if len(rl) > 1 {
		return nil, errors.New("Unknown Error")
	}

	return rl[0], nil
}

func (s *RelationDriver) GetMulti(ids []int64) ([]Relation, error) {
	return []Relation{}, nil
}

func (s *RelationDriver) GetAll() ([]Relation, error) {



	return []Relation{}, nil
}

func (s *RelationDriver) Remove(){}
func (s *RelationDriver) RemoveMulti(){}

