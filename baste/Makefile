
.PHONY: revendor
revendor: 
	go mod tidy
	go mod vendor

.PHONY: update-protos
update-protos:
	go get -u github.com/ssargent/apis

.PHONY: build-linux
build-linux:
	rm -rf ./build/dist
	mkdir -p ./build/dist/amd/linux
	mkdir -p ./build/dist/arm/linux
	GOOS=linux GOARCH=amd64 go build -o ./build/dist/amd/linux/bbqctl ./cmd/bbqctl
	GOOS=linux GOARCH=arm64 go build -o ./build/dist/arm/linux/bbqctl ./cmd/bbqctl

.PHONY: push-gimli
push-gimli:
	scp ./build/dist/amd/linux/bbqctl scott@gimli:/home/scott/bbqctl

.PHONY: run-local
run-local: build-bbq
	./bin/bbq serve -e .env/development.env 

.PHONY: build-bbq
build-bbq:
	go build -o ./bin/bbq ./cmd/bbq

PHONY: export-schema
export-schema:
	@pg_dump --schema-only -d bbqd > ./sql/schema.sql

.PHONY: generate
generate: export-schema
	@sqlc generate