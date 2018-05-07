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

func Save() {
	//interfaceを引数に取り，errorを返す
	//基本的にuuidは生成時にセットし，keyとして用いる
	//存在しないかをチェックし(もしかしたらkey生成時に検知してくれるか？)，上書きオプションがあれば上書きとする
	//https://hori-ryota.com/blog/golang-highly-versatile-method-with-arbitary-options/
}

func Load() {
	//基本的にreflectで変数の生成まで行う形とする
	//無理そうなら型情報をmapの形で保存させ，それを元に生成させる形とする
}

//TODO:
/*
構造体を渡せば書き込みが出来，Getする時点で構造体の配列として返せる形にする．
入れる時も出す時もとりあえず配列を前提とし，後で自動判別のものに切り替える
*/
//where(table), key
