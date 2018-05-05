package dsdriver

import (
	"context"
	"google.golang.org/appengine/datastore"
)

type DS struct {
	c context.Context
}

func NewDSDriver(c context.Context) *DS {
	s := new(DS)
	s.c = c

	return s
}

//TODO:
/*
構造体を渡せば書き込みが出来，Getする時点で構造体の配列として返せる形にする．
入れる時も出す時もとりあえず配列を前提とし，後で自動判別のものに切り替える
*/
//where(table), key
