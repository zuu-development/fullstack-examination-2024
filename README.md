# Fullstack examination 2024

## What is Fullstack examination 2024?

This is a repository for exams used to assess technical skills for hiring full-stack engineers (Golang).

## Dev

You will set up the development environment using asdf, a tool for managing multiple language versions.
If you are in an environment where asdf cannot be used, please install the necessary tools according to your environment.

### Install

#### Install asdf

[Getting Started | asdf](https://asdf-vm.com/guide/getting-started.html)

#### Install asdf plugins

```bash
asdf plugin add air
asdf plugin add golang
asdf plugin add golangci-lint
asdf plugin add gotestsum
asdf plugin add nodejs
asdf plugin add swag
```

#### Install asdf versions

Please install the versions listed in the `.tool-versions` file.

```bash
asdf install
```

### Start Development Environment

#### backend

```bash
make dep-backend-local
```

```bash
make migrate
```

```bash
make serve-backend
```

#### ui

```bash
make dep-ui-local
```

```bash
make serve-ui
```

When you access [http://localhost:3000/](http://localhost:3000/), the UI screen will be displayed.

### Migration

Please run the migration process.

```bash
make migrate
```

If the migration fails due to the current state of the schema, please delete the database and run the migration again.

```bash
make reset-local-db migrate
```

### Format

To maintain consistency in the code, formatting should be applied. Be sure to run it once development is complete.

```bash
make fmt
```

### Swagger

Please create API documentation using Swagger.

```bash
make swagger
```

To generate documentation, you need to add comments to your Go files. Please refer to this documentation for the correct syntax and guidelines.

[swaggo/swag: Automatically generate RESTful API documentation with Swagger 2.0 for Go.](https://github.com/swaggo/swag#declarative-comments-format)

While running `make serve-backend`, you can access [http://localhost:1314/](http://localhost:1314/) to view the Swagger UI. Or alternatively, you can open `docs/swagger.yaml` to view the API documentation.

### Test

Please run the tests.

```bash
make test-backend
```

## examination

Please solve the tasks outlined in [examination.md](./examination.md).
