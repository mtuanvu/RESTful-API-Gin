include .env
export

CONN_STRING = postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SLLMODE)
MIGRATION_DIRS	= internal/db/migrations 

#import database
importdb:
	docker exec -i postgres-db psql -U root -d tuanvu1234 < ./namefile.sql

#export database
exportdb: 
	docker exec -i postgres-db pg_dump -U root -d tuanvu1234 > ./namefile.sql

#Run server
server:
	go run .

# create new migration
migrate-create:
	migrate create -ext sql -dir $(MIGRATION_DIRS) -seq $(NAME)

# run all pending migration
migrate-up:
	migrate -path $(MIGRATION_DIRS) -database "$(CONN_STRING)" up

migrate-down:
	migrate -path $(MIGRATION_DIRS) -database "$(CONN_STRING)" down 

.PHONY: importdb exportdb server migrate-create migrate-up migrate-down