## boiler_plate_go
golangによるapiサーバのボイラープレートです

DDDをある程度踏襲してます

## commands
### make genMock param=user
usecase/user.goのテストを作成する際にrepositoryのパッケージのmockを作成します

### make connDB env=local
local環境のmysqlに接続します
-prootです

### make initDB env=local
local環境のmysqlにinitdb.dのsqlファイルを流し込みます

### make testUC
usecase層のテストを実行します

### make testEntry
entry層のテストを実行します

## architecture

### overview

```
handler
↓
usecase
↓
domain
↑
infra
```

domainドメイン部分にrepositoryを設定し、infra層はダックタイピングによってrepositoryのインターフェース定義を満たしてる
usecaseはrepositoryを使用

error構造体やlogger構造体は独自パッケージ

### tree
```
❯ tree ./cmd ./pkg
./cmd
└── api
    ├── Dockerfile
    └── main.go
./pkg
├── config
│   ├── common.yaml
│   ├── config.go
│   ├── local.yaml
│   └── test.yaml
├── database
│   └── database.go
├── domain
│   ├── entity
│   │   └── user.go
│   └── repository
│       ├── repository.go
│       └── user.go
├── error
│   ├── db.go
│   ├── error.go
│   ├── sql.go
│   ├── user.go
│   └── validation.go
├── handler
│   ├── handler.go
│   └── user.go
├── infra
│   ├── dao
│   │   └── user.go
│   └── entry
│       ├── infra.go
│       ├── infra_test.go
│       ├── transact.go
│       ├── user.go
│       └── user_test.go
├── logger
│   └── logger.go
└── usecase
    ├── input
    │   └── user.go
    ├── mock
    │   └── user.go
    ├── usecase.go
    ├── usecase_test.go
    ├── user.go
    └── user_test.go

```

### cmd/api
一番最初のapiサーバのエントリー

### pkg/handler
エンドポイントの入り口

### pkg/usecase/input
handler層でリクエストを変換させる型,及びその変換する関数群

### pkg/usecase/
リポジトリ層にリストされてる関数を利用してビジネスロジックを記載する

### pkg/usecase/mock
テストでリポジトリのモックが格納されている

### pkg/infra/dao
DBを直接叩く.最小構成要素

### pkg/infra/entry
daoで記載されてる最小要素を組み合わせてトランザクションを用いることができる

リポジトリでダックタイピングしている箇所は基本ここ

### pkg/domain/entity
ドメイン知識

### pkg/domain/repository
infra/entryのインターフェース

### pkg/error
独自のエラー構造体パッケージ
DBエラーなどもラップされている

### pkg/datagase
DBとのコネクションの管理

### pkg/logger
独自のloggerパッケージ

環境ごとに出力するログレベルを変えてくれる
