include .env

stop_containers:
	@echo "Stopping other docker containers.."
	if [ $$(docker ps -q) ]; then \
		echo "found and stopped containers."; \
		docker stop $$(docker ps -q); \
	else \
		echo "no continers currently running..;" \
	fi

create_container:
	docker run --name ${DB_DOCKER_CONTAINER} -p 5432:5432 -e POSTGRES_USER=${USER} -e POSTGRES_PASSWORD=${PASSWORD} -d postgres:12-alpine

create_db:
	docker exec -it ${DB_DOCKER_CONTAINER} createdb --username=${USER} --owner=${USER} ${DB_NAME}

start_container:
	docker start ${DB_DOCKER_CONTAINER}

create_migrations:
	sqlx migrate add -r init

migrate_up:
	sqlx migrate run --database-url "postgres://${USER}:${PASSWORD}@${HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

migrate_down: 
	sqlx migrate revert --database-url "postgres:${USER}:{$PASSWORD}@${HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"