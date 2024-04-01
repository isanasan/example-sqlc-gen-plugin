# sqlc-gen-hello

ビルド

```bash
cd plugin
tinygo build -o ../bin/sqlc-gen-hello.wasm -gc=leaking -scheduler=none -target=wasi -no-debug
```

sha256生成

```bash
openssl sha256 bin/sqlc-gen-hello.wasm
```
