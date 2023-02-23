db.init:
	docker run -d --name auth -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=lemmein -e POSTGRES_DB=auth -p 5432:5432 postgres:14-alpine


wire.install:
	go install github.com/google/wire/cmd/wire@latest

wire.gen:
	wire ./cmd/

ent.install:
	go get -d entgo.io/ent/cmd/ent

ent.init:
	go run -mod=mod entgo.io/ent/cmd/ent new --target internal/ent/schema User

ent.gen:
	go generate ./internal/ent/...

swag.install:
	go get github.com/swaggo/swag/cmd/swag
	go install github.com/swaggo/swag/cmd/swag
	go get github.com/swaggo/gin-swagger
	go get github.com/swaggo/files

swag.gen:
	swag init -d ./cmd,./internal/service,./internal/model,./pkg/errors -g main.go --output docs/swagger

go.install:
	wire.install
	ent.install
	swag.install
