# Fullstack examination 2024

## What is Fullstack examination 2024?

ZUUのフルスタックエンジニアを採用する際に、技術的なスキルを測るための試験のためのリポジトリです。

## Dev

asdfを使って開発環境を構築します。asdfは複数の言語のバージョン管理を行うツールです。
もしasdfを利用できない環境の場合は、環境に合わせたインストールを行ってください。

### Install

#### Install asdf

[Getting Started | asdf](https://asdf-vm.com/guide/getting-started.html)

#### Install asdf plugins

```bash
asdf plugin add air
asdf plugin add golang
asdf plugin add gotestsum
asdf plugin add nodejs
asdf plugin add swag
```

#### Install asdf versions

.tool-versions に記載されているバージョンをインストールします。

```bash
asdf install
```

### Start Development Environment

### backend

```bash
make dep-backend-local
```

```bash
make serve-backend
```

### ui

```bash
make dep-ui-local
```

```bash
make serve-ui
```

[http://localhost:3000/](http://localhost:3000/) にアクセスすると、UIの画面が表示されます。

### Format

コードに統一性を持たせるために、フォーマットを行います。開発が終わったら実行しましょう。

```bash
make fmt
```

### Swagger

Swaggerを利用してAPIのドキュメントを作成します。

```bash
make swagger
```

ドキュメント生成のために、Goファイルへのコメント記述が必要です。
記法はこちらのドキュメントを参考にしてください。

[swaggo/swag: Automatically generate RESTful API documentation with Swagger 2.0 for Go.](https://github.com/swaggo/swag?tab=readme-ov-file#declarative-comments-format)

`make serve-backend`を実行している状態で、[http://localhost:1314/](http://localhost:1314/) にアクセスするとSwagger UIが表示されます。

もしくは`docs/swagger.yaml`を開いてください。

### Test

テストを実行します。

```bash
make test-backend
```
