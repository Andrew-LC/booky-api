APP_NAME=bookmark-api
DB_HOST=localhost
DB_PORT=5432
DB_USER=myuser
DB_DBNAME=mydb
DB_PASSWORD=mypass
DB_SSLMODE=disable

.PHONY: run
run:
	@echo "Running app locally..."
	@DATABASE_URL="host=$(DB_HOST) user=$(DB_USER) password=$(DB_PASSWORD) dbname=$(DB_DBNAME) port=$(DB_PORT) sslmode=$(DB_SSLMODE)" \
	go run main.go
