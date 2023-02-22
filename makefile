init:
	go mod tidy
	docker-compose up
	sql-migrate up -config=dbconfig.yml -env="development"

dev:
	go run ./main.go

migration-up:
	sql-migrate up -config=dbconfig.yml -env="development"

migration-down:
	sql-migrate down -config=dbconfig.yml -env="development"