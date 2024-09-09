redis-up:
	@docker run -d -p 6379:6379 --name cache redis
redis-down:
	@docker stop cache && docker rm cache

postgres-up:
	@docker run -d -p 5432:5432 -e POSTGRES_PASSWORD="postgres" --name db postgres && \
	docker cp data.sql db:/  
postgres-dump: 
	@docker exec -ti db psql -U postgres -f data.sql
postgres-down:
	@docker stop db && docker rm db
