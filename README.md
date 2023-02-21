# Go to Gym
## タスク
- [ ] 調査
    - [ ] https://github.com/fsouza/fake-gcs-server
- [ ] Ginのミドルウェアの中で、パスワードを保持しているUserモデルレスポンスのパスワードを破棄して返すことをしたい(現状はUseCase各操作の中でサニタイズしており、実装がもれると危ない)

## 使用技術(バックエンド)
- Webフレームワーク
    - Gin
- DB
    - MySQL
- OrMapper
    - [Genericsを使いミスを防ぐSQL Builder「GenORM」](https://zenn.dev/mazrean/articles/c795c04e4837b4)
- KVS(Session)
    - Reds
- DBMigration
    - Golang-Migrate
- Storage Service
    - GCS(ローカルではMinioみたいに使う)
- API Definition
- DashBoard
    - Prometheus
- ロギング
    - ZapLogger
- Gomock
- Tester
    - Testify

### その他使いたいパッケージ
- https://www.jaegertracing.io/
- https://prometheus.io/
- nginx
- https://grafana.com/

### その他やりたいこと
- サーバはHTTPSで認証を必須とする
- デバッグAPIも標準で動作させる
    - 本番環境でも動作させるために、IPアドレス制限を設けるなどする
- ファイルストレージの使い道としては、CSVのバッチ処理などすればいいと思った
    - バッチは別でインフラを立てる必要がある
    - インフラに関しても定義は自分のソースコード内に置いとけばいいや

## 使用技術(フロントエンド)
- 単体テスト
- E2Eテスト

## 使用技術(インフラ)
- Terraform
- Docker-Compose

## memo
- フロントとバックエンド、共通リポジトリにしたい
    - 定数は抽出して両方で使いまわせるようにしたい
- 定数などは外出ししたい


## 提供するAPI定義

- ユーザ
  - サインアップ
  - サインイン
  - サインアウト
  - 退会
- コンフィグ
  - 各種設定変更
- 種目
  - 作成
  - 一覧
  - 単体
  - 更新
  - 削除
- セットメニュー
  - 作成
  - 一覧
  - 単体
  - 更新
  - 削除
- 記録
  - 記録
  - 一覧
    - 普通に一覧
    - 部位ごと集計
  - 単体
  - 編集
  - 削除