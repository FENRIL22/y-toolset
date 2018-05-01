package tasker

/* this is dust */

import (
	//"google.golang.org/appengine"
	//"net/http"
	"io/ioutil"
	"net/http/httptest"
	//"strings"
	"bytes"
	//"fmt"
	"google.golang.org/appengine/aetest"
	"testing"
)

func TestDB(t *testing.T) {
	//ctx, done, err := aetest.NewContext()
	//if err != nil {
	//	t.Fatal(err)
	//}
	//defer done()
	//ctxが確保できる
	//ほかログイン等にも使えるかも
}

func TestResp(t *testing.T) {

	opt := aetest.Options{StronglyConsistentDatastore: true}
	instance, err := aetest.NewInstance(&opt)
	if err != nil {
		t.Fatalf("Failed to create aetest instance: %v", err)
	}
	defer instance.Close()

	//req, _ := instance.NewRequest("POST", "/tasker/test", strings.NewReader(`{"post_value":"hoge","foo":"foo"}`))
	jsonStr := `{"post_value":"hoge","foo":"foo"}`
	req, _ := instance.NewRequest("POST", "/", bytes.NewBuffer([]byte(jsonStr)))
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Type", "application/json")

	t.Log(req)
	// コンテキストの取得
	//ctx := appengine.NewContext(req)

	res := httptest.NewRecorder()

	tasker := NewTasker(res, req)
	tasker.Test()

	b, _ := ioutil.ReadAll(res.Body)

	t.Log(string(b))
}
