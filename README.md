# GO-GRAPHQL-SERVER-TEST

GraphQL Server Side の実装を Go でお試しする。

## 準備

### gqlgen

`tools/tools.go` を作成し、以下のコードを貼り付ける。

```go
package tools

import (
	_ "github.com/99designs/gqlgen"
)
```

`go mod tidy` を実行する。
`go run github.com/99designs/gqlgen init` を実行する。

成功すると以下のメッセージが出力される。

```sh
Creating gqlgen.yml
Creating graph/schema.graphqls
Creating server.go
Generating...

Exec "go run ./server.go" to start GraphQL server
```

`go run ./server.go` でサーバーを起動すると、 [GraphiQL](http://localhost:8080/) へとアクセス可能となる。

### air

Live reload tool.

- `go install github.com/cosmtrek/air@v1.44.0` を実行する。
- `air init` を実行し、 `.air.toml` ファイルを生成する。
- `.air.toml` を編集し、 `air` コマンドを実行する。

## 使い方

### DB の起動など

`make` コマンドに実行コマンドをまとめている。

| DB    | 操作 | コマンド          |
| ----- | ---- | ----------------- |
| MySQL | 起動 | `make mysql/up`   |
| MySQL | ログ | `make mysql/log`  |
| MySQL | 接続 | `make mysql/exec` |
| MySQL | 終了 | `make mysql/rm`   |

### Go と MySQL の接続確認

`./connectiontest/main.go` に動作確認用のコードを作成している。
`make connectiontest` を実行することで接続確認を行うことができる。
