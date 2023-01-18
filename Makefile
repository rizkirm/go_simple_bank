migrateup:
	migrate -path db/migration -database "postgresql://root:password@db-postgre:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable" -verbose down
