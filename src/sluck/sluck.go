package sluck

import (
	"io"
	"io/ioutil"
	"os"
	"fmt"
	//"log"
	"strings"
	"net/http"
	"net/url"
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

const (
	// action is used for slack attament action.
	actionSelect = "select"
	actionStart  = "start"
	actionCancel = "cancel"
)

func (s *Sluck) EventAPI(){
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

			params := slack.PostMessageParameters{
				Attachments: []slack.Attachment{
					attachment,
				},
			}

			//普通に投稿するとき
			_, _, _ = api.PostMessage(pp.String("channel"), "こんにちは", slack.PostMessageParameters{})

			//Interactive API
			if _, _, err := api.PostMessage(pp.String("channel"), "", params); err != nil {
				log.Errorf(ctx, "failed to post message: ", err)
			}

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


var attachment = slack.Attachment{
     Text:       "Which beer do you want? :beer:",
     Color:      "#f9a41b",
     CallbackID: "beer",
     Actions: []slack.AttachmentAction{
         {
             Name: actionSelect,
             Type: "select",
             Options: []slack.AttachmentActionOption{
                 {
                     Text:  "Asahi Super Dry",
                     Value: "Asahi Super Dry",
                 },
                 {
                     Text:  "Kirin Lager Beer",
                     Value: "Kirin Lager Beer",
                 },
                 {
                     Text:  "Sapporo Black Label",
                     Value: "Sapporo Black Label",
                 },
                 {
                     Text:  "Suntory Malt's",
                     Value: "Suntory Malts",
                 },
                 {
                     Text:  "Yona Yona Ale",
                     Value: "Yona Yona Ale",
                 },
             },
         },

         {
             Name:  actionCancel,
             Text:  "Cancel",
             Type:  "button",
             Style: "danger",
         },
     },
 }


func ResponseToRequest(w http.ResponseWriter, r *http.Request){
	ctx := appengine.NewContext(r)

	if r.Method != http.MethodPost {
         log.Errorf(ctx, "[ERROR] Invalid method: %s", r.Method)
         w.WriteHeader(http.StatusMethodNotAllowed)
         return
     }

     buf, err := ioutil.ReadAll(r.Body)
     if err != nil {
         log.Errorf(ctx, "[ERROR] Failed to read request body: %s", err)
         w.WriteHeader(http.StatusInternalServerError)
         return
     }

     jsonStr, err := url.QueryUnescape(string(buf)[8:])
     if err != nil {
         log.Errorf(ctx, "[ERROR] Failed to unespace request body: %s", err)
         w.WriteHeader(http.StatusInternalServerError)
         return
     }

     var message slack.AttachmentActionCallback
     if err := json.Unmarshal([]byte(jsonStr), &message); err != nil {
         log.Errorf(ctx, "[ERROR] Failed to decode json message from slack: %s", jsonStr)
         w.WriteHeader(http.StatusInternalServerError)
         return
     }

     // Only accept message from slack with valid token
     //if message.Token != h.verificationToken {
     //    log.Errorf(ctx, "[ERROR] Invalid token: %s", message.Token)
     //    w.WriteHeader(http.StatusUnauthorized)
     //    return
     //}

	 action := message.Actions[0]
     switch action.Name {
     case actionSelect:
         value := action.SelectedOptions[0].Value

         // Overwrite original drop down message.
         originalMessage := message.OriginalMessage
         originalMessage.Attachments[0].Text = fmt.Sprintf("OK to order %s ?", strings.Title(value))
         originalMessage.Attachments[0].Actions = []slack.AttachmentAction{
             {
                 Name:  actionStart,
                 Text:  "Yes",
                 Type:  "button",
                 Value: "start",
                 Style: "primary",
             },
             {
                 Name:  actionCancel,
                 Text:  "No",
                 Type:  "button",
                 Style: "danger",
             },
         }

         w.Header().Add("Content-type", "application/json")
         w.WriteHeader(http.StatusOK)
         json.NewEncoder(w).Encode(&originalMessage)
         return
     case actionStart:
         title := ":ok: your order was submitted! yay!"
         responseMessage(w, message.OriginalMessage, title, "")
         return
     case actionCancel:
         title := fmt.Sprintf(":x: @%s canceled the request", message.User.Name)
         responseMessage(w, message.OriginalMessage, title, "")
         return
     default:
         log.Errorf(ctx, "[ERROR] ]Invalid action was submitted: %s", action.Name)
         w.WriteHeader(http.StatusInternalServerError)
         return
     }
}

func responseMessage(w http.ResponseWriter, original slack.Message, title, value string) {
     original.Attachments[0].Actions = []slack.AttachmentAction{} // empty buttons
     original.Attachments[0].Fields = []slack.AttachmentField{
         {
             Title: title,
             Value: value,
             Short: false,
         },
     }

     w.Header().Add("Content-Type", "application/json")
     w.WriteHeader(http.StatusOK)
     json.NewEncoder(w).Encode(&original)
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

//type SlackListener struct {
//	client    *slack.Client
//	botID     string
//	channelID string
//}

//func (s *SlackListener) InteractiveAPI(w http.ResponseWriter, r *http.Request){
//	ctx := appengine.NewContext(r)
//	// Start listening slack events
//	slack.SetHTTPClient(urlfetch.Client(ctx))
//	// 環境変数からアクセストークンを取得. SLACK_TOKEN は appengine/secret.yaml に定義されている.
//	token := os.Getenv("SLACK_TOKEN")
//
//	log.Infof(ctx, "%v\n",token)

