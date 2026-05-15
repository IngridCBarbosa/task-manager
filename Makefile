.PHONY: redis-up redis-down redis-logs redis-cli

# Starts Redis container
redis-up:
	docker run -d \
		--name task-manager-redis \
		-p 6379:6379 \
		redis

# Stops and removes Redis container
redis-down:
	docker stop task-manager-redis && docker rm task-manager-redis

# Shows Redis container logs
redis-logs:
	docker logs -f task-manager-redis

# Opens Redis CLI inside the container
redis-cli:
	docker exec -it task-manager-redis redis-cli