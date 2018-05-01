package tasker

/* this is dust */

import (
	//"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
	//"net/http"
	"io/ioutil"
	"net/http/httptest"
	//"strings"
	"bytes"
	"fmt"
	"testing"
)

func TestResp(t *testing.T) {
	//or?
	//ctx, done, err := aetest.NewContext()
	//if err != nil {
	//	t.Fatal(err)
	//}
	//defer done()

	opt := aetest.Options{StronglyConsistentDatastore: true}
	instance, err := aetest.NewInstance(&opt)
	if err != nil {
		t.Fatalf("Failed to create aetest instance: %v", err)
	}
	defer instance.Close()

	t.Log("Input")
	fmt.Scan(&sttt)
	//req, _ := instance.NewRequest("POST", "/tasker/test", strings.NewReader(`{"post_value":"hoge","foo":"foo"}`))
	jsonStr := `"post_value":"hoge","foo":"foo"`
	req, _ := instance.NewRequest("POST", "/tasker", bytes.NewBuffer([]byte(jsonStr)))
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// コンテキストの取得
	//ctx := appengine.NewContext(req)

	var sttt string

	res := httptest.NewRecorder()

	tasker := NewTasker(res, req)
	tasker.Test()

	b, _ := ioutil.ReadAll(res.Body)

	t.Log(string(b))
}
