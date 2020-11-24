.PHONY: reset sonarscanner

reset:
	docker-compose down -v
	docker-compose build
	docker-compose up -d

sonarscanner:
	docker run \
		--rm \
		-e SONAR_HOST_URL="http://sonarqube:9000" \
		-v "$(PWD):/usr/src" \
		--network=in-backend_backend \
		sonarsource/sonar-scanner-cli