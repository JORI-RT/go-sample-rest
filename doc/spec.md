* Goでは大文字から始まる関数は、外のモジュールからよべる（public）go ではexported nameとよぶ

* レシーバを伴うメソッドの宣言はレシーバ型が同じパッケージにある必要がある、他人が作った型にメソッド追加はできない

## GOでのUT
UTは言語に組み込まれている
組み込みのtestingパッケージを使う
関数名がtest名になる
```go
func test01(t *testing.T) {
    ...

}
```
実行のコマンドは
```go
go test
or 
go test -v 

```
