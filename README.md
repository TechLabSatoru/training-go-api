## 環境構築

```shell
$ docker compose up -d

$ docker compose exec -it golang bash

# モジュールの初期化（初回実行のみ）
/home/TechLabSatoru# go mod init training-go-api

/home/TechLabSatoru# go run main.go
```

## サーバ構築

```shell
$ go run main.go
```

- tar.gzファイル作成

```shell
$ tar zcvf sample.tar.gz sample/
```

- tar.gzファイルをサーバへPostする処理を記載する.

```shell
$ curl -i -f -H "Authorization:Bearer `cat token`" -H "content-type:application/x-binary" -X POST "localhost:18080" --data-binary "@$BINARY_DATA=PATH"

$ curl -i -f -H "content-type:application/x-binary" -X POST "localhost:18080" --data-binary "@./input/sample.tar.gz=PATH"

$ curl -i -f -H "content-type:application/x-binary" -X POST "localhost:18080" --data-binary "@input/sample.tar.gz"
```

- GET用

```shell
$ curl -i -f -X GET "localhost:18080/"
```

- POST用（サンプル）

```shell
$ curl -i -f -H "content-type:application/x-binary" -X POST "localhost:18080/v1/upload" --data-binary "@input/sample.tar.gz"
```
