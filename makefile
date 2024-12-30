# Docker Composeファイルの指定
DOCKER_COMPOSE_FILE=docker-compose.yml

# コンテナを起動
up:
	docker compose -f $(DOCKER_COMPOSE_FILE) up -d --build

# コンテナを停止
down:
	docker compose -f $(DOCKER_COMPOSE_FILE) down

# ログを表示
logs:
	docker compose -f $(DOCKER_COMPOSE_FILE) logs -f

# 再起動
restart:
	docker compose -f $(DOCKER_COMPOSE_FILE) down
	docker compose -f $(DOCKER_COMPOSE_FILE) up -d --build
