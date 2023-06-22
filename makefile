run:
	sh runproject.sh
build:
	docker-compose up --build
makemigrations:
	migrate create -ext sql -dir ./migration -seq init
migrate:
	migrate -path ./migration -database 'postgres://$(USER):$(PASS)@localhost:5432/postgres?sslmode=disable' up
rollback:
	migrate -path ./migration -database 'postgres://$(USER):$(PASS)@localhost:5432/postgres?sslmode=disable' up