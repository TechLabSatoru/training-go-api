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
    - 通常のHTTPリクエスト
```shell
$ curl -i -f -X GET "localhost:18080/"
```
    - TLS（HTTPS）リクエスト
```shell
$ curl -i -f -X GET "https://localhost:18080/"
```
    - mTLS用リクエスト
```shell
$ curl -i -f --cert ./input/client.crt --key ./input/client.key https://localhost:18080/
or
$ curl -i -f --cert ./input/client.crt --key ./input/client.key --cacert ./input/ca.pem https://localhost:18080/
```

- POST用（サンプル）

```shell
$ curl -i -f -H "content-type:application/x-binary" -X POST "localhost:18080/v1/upload" --data-binary "@input/sample.tar.gz"
```

## 証明書の発行関係

```shell
# サーバ側
## Create CSR (Certificate Signing Request) and private key
$ openssl req -new -newkey rsa:2048 -nodes -keyout server.key -out server.csr

## Check the information
$ openssl req -text -noout -verify -in server.csr

# 認証局の役割
## Root証明書の作成
$ openssl req -x509 -new -nodes -keyout ca.key -sha256 -days 1024 -out ca.pem

## 証明書の発行
$ openssl x509 -req -in server.csr -CA ca.pem -CAkey ca.key -CAcreateserial -out server.crt -days 500 -sha256
```

```shell
# クライアント側
## 秘密鍵の作成
$ openssl genpkey -algorithm RSA -out client.key

## クライアント証明書の生成
$ openssl req -new -key client.key -out client.csr

## クライアント証明書の発行（認証局）
$ openssl x509 -req -in client.csr -CA ca.pem -CAkey ca.key -CAcreateserial -out client.crt -days 500 -sha256

$ openssl req -text -noout -verify -in client.csr
```

## 証明書の内容確認

```shell
$ openssl x509 -in ca.pem -text -noout

$ openssl x509 -in server.crt -text -noout

$ openssl x509 -in client.crt -text -noout
```

- CA認証局のCNとServer側のCNは異なる必要がある.

```shell
$ openssl verify -CAfile ca.pem server.crt

$ openssl verify -CAfile ca.pem client.crt
```

```shell
$ openssl s_client -connect localhost:18080 -CAfile ca.pem
```
