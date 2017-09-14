# y-toolset
## はじめに
GAEに関してのメモも同時にする．
エキサイト翻訳並の文章校正しかしてません．
気が向いたら整理するかも．

## 基本方針
とりあえずAncestorを使わずQueryでやる．
(強い整合性を必要としない用途なため)

## 情報源(ほぼ参考を総括する)
* [App Engine + Golangでチートシートを作るだけのアプリを作った](http://twinbird-htn.hatenablog.com/entry/2017/01/06/180815)
	* Appengineについての情報がバランス良く纏まっている
* [GAE/Go Datastoreの基本的なクエリ](http://knightso.hateblo.jp/entry/2014/05/26/103458)
	* 上の中にあり
	* Queryの条件付け等について記載
* [The datastore package](https://cloud.google.com/appengine/docs/standard/go/datastore/reference#NewKey)
	* APIの型などの情報

* [スクレイピング結果をJSONで返すAPIをGAE/Goで作る。](http://qiita.com/CST_negi/items/ce81fde4dc910e063e61)
	* JSON応答の一つの形
* [GolangでレスポンスにJSONで返す](http://konboi.hatenablog.com/entry/2014/09/23/172756)
	* API作る時に使う
* [GAE/Go で Events API を使った Slack ボットをつくる](http://qiita.com/nirasan/items/3060260a63202638f0f4)
	* SlackをUIとして使う予定




## ファイル構造の参考
* [Google App EngineでGoを動かすときに知っておくべきこと](http://motemen.hatenablog.com/entry/2016/11/gae-go-building#アプリケーションを構成するパッケージの置き場所とインポート)
* [cloud Strage](http://otiai10.hatenablog.com/entry/2016/04/03/015038)
* [datastore api](https://cloud.google.com/appengine/docs/standard/go/datastore/reference#Key)
* [datastore query - データ取得](http://knightso.hateblo.jp/entry/2014/05/26/103458)
* [datastore client package](https://cloud.google.com/datastore/docs/reference/libraries?hl=ja)

## URLからの処理とか
* [Go+GAE+Cloud Datastoreで簡単なREST APIを構築](http://qiita.com/silverfox/items/81769425e51f24e676d2)

## hoge 
* [gcloud deploy](https://cloud.google.com/appengine/docs/standard/go/tools/uploadinganapp)




## Initialize
### インストール(GCP)
まずgolangをbrewなどで入れて[Download](https://cloud.google.com/sdk/)して解凍して出てきたファイルを置いときたい場所に置いてインストールして`gcloud init`してGOPATHを`/(path-to-->file)/google-cloud-sdk/platform/google_appengine/gopath`とかに設定する．このGOPATHはGoを入れてから出現するのでその時に見ると良い．GOPATHはパッケージのところに作ったりすることもできるので，よしなにやる．


### GCPの準備
アカウントはあるものとする．
プロジェクトの作成
GAEを有効化する(作成しただけかも)


### デプロイファイル作成
ここを参考に．とりあえずmakeファイルは作った．

[PaaS最前線！たったの15分でできるGAE/GO入門！](http://www.apps-gcp.com/gae-go-gettingstart-01/#i-4)


### 確認ほか
GOPATHを確認して`dev_appserver.py app.yaml`としてGoが入ってないと言われたらY押して入れて立ち上げを確認．


### 色々参考
[Go製サーバーをGoogleApp Engineに設置してみる&Google Cloud SQLの設定](http://onga-tec.hatenadiary.jp/entry/2016/08/13/032441)
[Google App EngineでGoのウェブアプリケーションをまず動かしてみる](http://qiita.com/taizo/items/845fcfc58cfd0b30020a)
[Google App Engineを無料で運用する方法（2016年版）](http://koni.hateblo.jp/entry/2016/01/06/130613)


## システム構成
### 入れたい機能
* [翻訳](https://soarcode.jp/posts/264)

### 将来的
api専用のprojectを作ってそこにアクセスとか．

* [アクセス制限](http://www.apps-gcp.com/gae-go-rest-api-1/)
* [アクセス制限の例](https://github.com/AustenConrad/go-lang-example-wiki-app/blob/master/app.yaml)
* [app.yaml Reference](https://cloud.google.com/appengine/docs/standard/go/config/appref)
* [Google App Engine - app.yamlについて](http://dackdive.hateblo.jp/entry/2015/03/14/225139)


### SPA(Single Page Application)化とか
やっぱり早い方が良い

* [Mithril.js](http://mithril-ja.js.org)
	* [最速MVCフレームワークMithril.jsの速度の秘密](http://qiita.com/shibukawa/items/890d24874655439932ec)
	* [Electron + Mithrilで、ふつうのデスクトップアプリを作る](http://qiita.com/shibukawa/items/e1836a8c98413448f746)


### Backend
* Google App Engine
* Google Datastore
	* [Go+GAE+Cloud Datastoreで簡単なREST APIを構築](http://qiita.com/silverfox/items/81769425e51f24e676d2)
* Google (user api? 認証用)



## 参考
[GAE](https://github.com/keitarog/foodle)
