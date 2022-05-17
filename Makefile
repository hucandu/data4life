create-pg-user:
	@psql -U postgres -c "CREATE USER one with encrypted password 'one';"

create-db:
	@psql -U postgres -c 'create database alpha with owner one;'

run:
	@go run cmd/main.go
