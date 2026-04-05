.PHONY: up-web up-swagger up-all down

up-web:
	docker compose up -d web

up-swagger:
	docker compose up -d swagger-ui

up-all:
	docker compose up -d

down:
	docker compose down
