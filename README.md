# 並行処理

## ゴルーチン

- Goのランタイムによって管理される「軽い」スレッド
- 関数名の前に`go`キーワードを置くとゴルーチンになる

## チャネル

```go
ch := make(chan int)
```

# 標準ライブラリ

## 入出力
入出力インターフェースは`io.Reader`と`io.Writer`には唯一のメソッドが定義されている

```go
package io

type Reader interface {
    Read(p []byte) (n int, err error)
}
type Writer interface {
    Write(p []byte) (n int, err error)
}
```