.SILENT:

run: db-deploy
	go mod download && go run ./cmd/main.go

#before "make test", do "make run" in another terminal
test:
	go clean -testcache && go test ./tests/ -v

mock-test:
	go clean -testcache && go test ./internal/grpcServer/ -v

db-deploy:
	docker-compose up -d database

#example: make db-insert AUTHOR="A.S. Pushkin" BOOK="Dubrovskiy"
db-insert: db-deploy
	go mod download && go run ./tests/mysql_insert -author "$(AUTHOR)" -book "$(BOOK)"

db-select: db-deploy
	go mod download && go run ./tests/mysql_select