# sqlc-gen-hello

ビルド

```bash
cd plugin
GOOS=wasip1 GOARCH=wasm go build -o ../bin/sqlc-gen-hello.wasm main.go
```

sha256生成

```bash
openssl sha256 bin/sqlc-gen-hello.wasm
```
