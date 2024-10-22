export POSTGRES_URL = 'postgres://root:root@localhost:5433/music-catalog?sslmode=disable'

migrate-create migrate_create:
	@ migrate create -ext sql -dir scripts/migrations -seq $(name)

migrate-up migrate_up:
	@ migrate -database ${POSTGRES_URL} -path scripts/migrations up 1

migrate-down migrate_down:
	@ migrate -database ${POSTGRES_URL} -path scripts/migrations down 1

d-restart d_restart:
	@ docker-compose down && docker-compose build && docker-compose up -d

d-postgres d_postgres:
	@ docker-compose down && docker-compose up local-postgres-music-catalog-db -d

d-down d_down:
	@ docker-compose down