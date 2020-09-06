 ```go
// verify version
go version
// execution go file
go run hello.go
// create module
got mod inig hello
```

## APの起動（Docker）
./start.sh
curl localhost:8081

## create module
```go
go mod init hogehoge/hoge

```
## Spec
* Goでは大文字から始まる関数は、外のモジュールからよべる（public）go ではexported nameとよぶ

## unpublishなモジュールをインポートする方法
import したいgo.modファイルに以下をかく
```go
replace example.com/greetings => ../greetings

// go buildコマンドを叩く
```