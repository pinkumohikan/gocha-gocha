run:
	docker run --rm -d -p "6379:6379" redis:latest
	sbt run

down:
	docker kill $(shell docker ps -q --filter "ancestor=redis")

