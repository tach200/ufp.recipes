start:
	docker compose up -d

stop:
	docker compose down -v

build-local-arm64:
	GOOS=linux GOARCH=arm64 go build -o out/ && \
	docker build . -t recipes

db-seed:
	docker exec -i ufprecipes-db-1 psql -U myuser -d mydb < internal/db/seed/tables.sql && \
	docker exec -i ufprecipes-db-1 psql -U myuser -d mydb < internal/db/seed/recipes.sql && \
	docker exec -i ufprecipes-db-1 psql -U myuser -d mydb < internal/db/seed/products.sql

db-query:
	docker exec -it ufprecipes-db-1 psql -h localhost -U myuser -d mydb -p 5432