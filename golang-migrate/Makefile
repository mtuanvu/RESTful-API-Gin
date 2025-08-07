include .env
export

CONN_STRING = postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SLLMODE)
MIGRATION_DIRS	= internal/db/migrations 

#import database


#export database


#Run server


# create new migration

# run all pending migration


migrate-down:
	migrate -path $(MIGRATION_DIRS) -database "$(CONN_STRING)" down 

.PHONY: importdb exportdb server migrate-create migrate-up migrate-down