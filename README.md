## backend
![](https://github.com/team-e-org/backend/workflows/go_test/badge.svg)
![](https://github.com/team-e-org/backend/workflows/go_integration_test/badge.svg)
[![codecov](https://codecov.io/gh/team-e-org/backend/branch/develop/graph/badge.svg)](https://codecov.io/gh/team-e-org/backend)

### 開発

`docker-compose up` でいろいろ立ち上がる。
  * localhost:3000でAPIサーバー
  * localhost:8080でOpenAPIのドキュメント
  * localhost:6379でredisサーバ
  
*1 起動に失敗した場合、ポート番号やローカルのMySQLと衝突している
可能性が高いです。

*2 goのコードを編集すると自動で再ビルドが走ります。
