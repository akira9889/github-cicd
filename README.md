# github-cicd

## コンテナに入って `go` コマンドを実行する

イメージをビルドします。

```bash
docker build -t github-cicd-app .
```

対話シェルでコンテナに入ります（終了は `exit`）。

```bash
docker run -it --rm --entrypoint sh github-cicd-app
```

作業ディレクトリは `/app` です。例:

```bash
cd /app
go run ./src
go test ./...
go build -o myapp ./src
./myapp
```

Docker Compose を使う場合の例:

```bash
docker compose run --rm --entrypoint sh app
```

コンテナ内のパスは上記と同じく `/app` です。
