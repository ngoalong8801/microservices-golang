migrate-new:
	migrate create -ext sql -dir db/migration -tz UTC $(name)
migrate-up:
	migrate -source file://db/migration -database postgres://postgres:123456@localhost:5432/grpc -verbose up $(v)

migrate-down:
	migrate -source file://db/migration -database postgres://postgres:123456@localhost:5432/grpc -verbose down $(v)

migrate-force:
	migrate -source file://db/migration -database postgres://postgres:123456@localhost:5432/grpc -verbose force $(v)

.PHONY: migrate-down migrate-down migrate-new migrate-force