.PHONY: local
local: start-localdb

.PHONY: clean
clean: clean-localdb

.PHONY: start-localdb
start-localdb:
	docker compose -f "localdb/docker-compose.yaml" up -d

.PHONY: clean-localdb
clean-localdb:
	docker compose -f "localdb/docker-compose.yaml" down
