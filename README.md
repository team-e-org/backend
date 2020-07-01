## backend
![](https://github.com/team-e-org/backend/workflows/go_test/badge.svg)
![](https://github.com/team-e-org/backend/workflows/go_integration_test/badge.svg)
[![codecov](https://codecov.io/gh/team-e-org/backend/branch/develop/graph/badge.svg)](https://codecov.io/gh/team-e-org/backend)

### 開発

`docker-compose up` でいろいろ立ち上がる。
  * localhost:3000でAPIサーバー
  * localhost:8080でOpenAPIのドキュメント
  * localhost:6379でredisサーバ
  * localhost:3306でMySQLサーバ
  
*1 起動に失敗した場合、ポート番号やローカルのMySQLと衝突している
可能性が高いです。

*2 goのコードを編集すると自動で再ビルドが走ります。

masterにpushするとstaging環境へ自動デプロイ

### テスト

`make unit_tests` で単体テスト

`make integration_tests` で結合テスト

ローカルでカバレッジを見るには

`go test -race -v -cover -coverprofile=coverage.o ./...`

した後に

`go tool cover -html=coverage.o`

するとブラウザが立ち上がって、カバレッジの詳細が見られる。

### テーブル図

https://gyazo.com/1d05718fc2d6ed4fb85cb7412ca96d5d


### タグに関連する設計

https://user-images.githubusercontent.com/24651683/85977118-e8712f00-ba16-11ea-9924-b5a5bf97d36e.png
