wire:
	@cd internal/bootstrap && wire
migrate-up:
	migrate -path migrations -database "postgres://sportgroup:${db_password}@localhost:${db_port}/oss?sslmode=disable" -verbose up
