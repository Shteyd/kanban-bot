# local env develop
local-migrate-up:
	migrate -path ./migrations -database postgres://postgres:postgres@localhost:5432/kanban_bot?sslmode=disable up

local-migrate-down:
	migrate -path ./migrations -database postgres://postgres:postgres@localhost:5432/kanban_bot?sslmode=disable down

local-docker-up:
	docker compose -f ./deploy/local/docker-compose.yml up

git-sync:
	git fetch --all -Pp