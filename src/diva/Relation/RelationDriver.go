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

func (s *RelationDriver) Put(re Relation) error {
	_ = datastore.NewKey(s.context, "Relation", "", 0, nil)
	//key := datastore.NewKey(s.context, "Relation", "", 0, nil)
	return nil
}

func (s *RelationDriver) PutMulti(re []Relation) error {
	//key := datastore.NewKey(s.context, "Relation", "", 0, nil)

	return errors.New("Undefined function")
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
	//return []Relation{}, nil
	return nil, errors.New("UndefineFunc")
}

func (s *RelationDriver) GetAll() ([]Relation, error) {
	ctx := s.context
	var rl []Relation

	q := datastore.NewQuery("Relation")
	q.GetAll(ctx, &rl)

	return rl, nil
}

//TODO:correspond to int64 + Relation?
func (s *RelationDriver) Remove(id int64) error{
	ctx := s.context
	var rl []Relation

	q := datastore.NewQuery("Relation").Filter("id =", id).KeysOnly()
	keys, e := q.GetAll(ctx, &rl)

	if e != nil {
		return e
	}

	for _, k := range keys{
		e = datastore.Delete(ctx, k)
	}

	return e
}

func (s *RelationDriver) RemoveMulti() error{
	return errors.New("UndefineFunc")
}

