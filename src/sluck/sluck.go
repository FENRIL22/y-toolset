package sluck

import (
	"io"
	"os"
	//"log"
	"strings"
	"net/http"
	"encoding/json"
	"github.com/nlopes/slack"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)


type Sluck struct{
	w http.ResponseWriter
	r *http.Request
	slackToken string
}

func (s *Sluck) Hoge() string{
	return "Hoge"
}

func (s *Sluck) Recog(){
	ctx := appengine.NewContext(s.r)
	//p, err := DecodeJSON(s.r.Body)
	p, _ := DecodeJSON(s.r.Body)

	//if err != nil {
	//    http.Error(s.w, err.Error(), http.StatusBadRequest)
	//    return
	//}

	switch p.Type() {
	case "url_verification":
		// イベント通知先URLの認証用アクセスへのレスポンス
		s.w.Header().Set("Content-Type", "text/plain")
		s.w.WriteHeader(http.StatusOK)
		s.w.Write([]byte(p.String("challenge")))
		return
		// イベント通知共通の種別
	case "event_callback":
		// イベント通知の詳細を取得する
		//data, ok := p["event"].(map[string]interface{})
		data, _ := p["event"].(map[string]interface{})
		//if !ok {
		//	http.Error(w, "failed to cast", http.StatusBadRequest)
		//	return
		//}
		pp := Payload(data)
		// ボット宛にメッセージを送信されたらレスポンスを返す
		log.Infof(ctx, "type: %v\n",pp.String("type"))
		log.Infof(ctx, "text: %v\n",pp.String("text"))
		if pp.String("type") == "message" && strings.Index(pp.String("text"), "<@U73LRSS75>") != -1 {
			// GAE から http リクエストを送信する場合は urlfetch ライブラリを利用する必要がある
			slack.SetHTTPClient(urlfetch.Client(ctx))
			// 環境変数からアクセストークンを取得. SLACK_TOKEN は appengine/secret.yaml に定義されている.
			token := os.Getenv("SLACK_TOKEN")

			log.Infof(ctx, "%v\n",token)

			api := slack.New(token)
			_, _, _ = api.PostMessage(pp.String("channel"), "こんにちは", slack.PostMessageParameters{})
			//_, _, err = api.PostMessage(pp.String("channel"), "こんにちは", slack.PostMessageParameters{})
			//if err != nil {
			//	log.Debugf(ctx, "failed to post message: %+v", err)
			//	http.Error(w, err.Error(), http.StatusBadRequest)
			//	return
			//}
			//log.Debugf(ctx, "Success")
		}
		s.w.WriteHeader(http.StatusOK)
		return
	}

}

func NewSluck(w http.ResponseWriter, r* http.Request) *Sluck{
	sl := new(Sluck)
	sl.w = w
	sl.r = r

	return sl
}

type Payload map[string]interface{}

func DecodeJSON(r io.Reader) (Payload, error) {
    data := make(map[string]interface{})
    if err := json.NewDecoder(r).Decode(&data); err != nil {
        return nil, err
    }
    return data, nil
}

func (p Payload) String(key string) string {
    if v, ok := p[key]; !ok {
        return ""
    } else if vv, ok := v.(string); !ok {
        return ""
    } else {
        return vv
    }
}

func (p Payload) Type() string {
    return p.String("type")
}

