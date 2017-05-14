# y-toolset
## Initialize
### インストール(GCP)
後に追記

### GCPの準備
アカウントはあるものとする．
プロジェクトの作成
GAEを有効化する(作成しただけかも)

$ gcloud auth login
gcloud components install app-engine-go
? $ gcloud components update gae-go

### デプロイファイル作成
ここを参考に．

[PaaS最前線！たったの15分でできるGAE/GO入門！](http://www.apps-gcp.com/gae-go-gettingstart-01/#i-4)

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

### Backend
* Google App Engine
* Google Datastore
* Google (user api? 認証用)



## 参考
[GAE](https://github.com/keitarog/foodle)
