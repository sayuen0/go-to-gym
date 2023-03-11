# Go to Gym
## タスク

- [ ] 調査
    - [ ] https://github.com/fsouza/fake-gcs-server
- [ ] 適切な権限管理
  - [ ] マスタユーザ以外は、自分のデータ以外を編集することはできない

## 使用技術(バックエンド)
- Webフレームワーク
    - Gin
- DB
    - MySQL
- OrMapper
    - SQLBoiler
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
    - Zap Logger
- Gomock
- Tester
    - Testify

### その他使いたいパッケージ
- https://www.jaegertracing.io/
- https://prometheus.io/
- nginx
- https://grafana.com/

### その他やりたいこと
- ~~サーバはHTTPSで認証を必須とする~~
- デバッグAPIも標準で動作させる
    - 本番環境でも動作させるために、IPアドレス制限を設けるなどする
- ファイルストレージの使い道としては、CSVのバッチ処理などすればいい
    - バッチは別でインフラを立てる必要がある
    - インフラに関しても定義は自分のソースコード内に置いとけばいい

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
