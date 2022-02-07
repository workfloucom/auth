.PHONY: test
test:
	go test ./test

.PHONY: dev
dev:
	go run cmd/server/main.go \
	--driver="sqlite" \
	--dsn="file:db.sqlite?cache=shared" \
	--jwt-secret="secret"

.PHONY: migrate
migrate:
	go run cmd/migrate/main.go \
	--driver="sqlite" \
	--dsn="file:db.sqlite?cache=shared"

.PHONY: seed
seed:
	go run cmd/seed/main.go \
	--driver="sqlite" \
	--dsn="file:db.sqlite?cache=shared"