migrateup:
	migrate -path db/migration -database "postgresql://root:123456@192.168.123.99:25433/bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:123456@192.168.123.99:25433/bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	gp test -v -cover ./...

.PHONY: migrateup migratedown sqlc test