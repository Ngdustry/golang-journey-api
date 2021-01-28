.PHONY: docker-up \
docker-down

docker-up:
	docker-compose -f .docker/docker-compose.yml up \
	--remove-orphans

docker-down:
	docker-compose -f .docker/docker-compose.yml down \
	--remove-orphans
