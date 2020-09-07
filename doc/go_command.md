 ```go
// verify version
go version
// execution go file
go run hello.go
// create module
got mod init hello

go install

go get module_name
```
## unpublishなモジュールをインポートする方法
import したいgo.modファイルに以下をかく
```go
replace example.com/greetings => ../greetings

// go buildコマンドを叩く
```