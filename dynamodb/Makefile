.PHONY: local-db
local-db:
	@docker-compose down
	@docker-compose up -d
	@while ! curl -s http://localhost:8000/shell/ > /dev/null; do \
		sleep 1; \
	done
	@./migrator/setup.sh