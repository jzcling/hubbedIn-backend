.PHONY: reset sonarscanner build

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

build: 
	docker build -t gcr.io/${PROJECT_ID}/profile-service:v0.1.0 -f ./services/profile/Dockerfile.prod .
	docker build -t gcr.io/${PROJECT_ID}/project-service:v0.1.0 -f ./services/project/Dockerfile.prod .
	docker build -t gcr.io/${PROJECT_ID}/gateway:v0.1.0 -f ./gateway/Dockerfile .
	docker build -t gcr.io/${PROJECT_ID}/api-gateway:v0.1.0 -f ./gateway/krakend/Dockerfile .
	