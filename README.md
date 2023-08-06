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
