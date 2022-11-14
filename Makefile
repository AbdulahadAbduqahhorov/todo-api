migrateup:
	migrate -path ./migrations -database 'postgres://abdulahad:passwd123@localhost:5432/todo?sslmode=disable' up
migratedown:
	migrate -path ./migrations -database 'postgres://abdulahad:passwd123@localhost:5432/todo?sslmode=disable' down
run :
	go run cmd/main.go 

	