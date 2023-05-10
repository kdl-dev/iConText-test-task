test:
	go test ./...

db_init:
	chmod +x schema/pg_init_db.sh && ./schema/pg_init_db.sh;

build:
	go build -o icontext_test_task cmd/main.go

run:
	 ./icontext_test_task -host localhost -port 6379