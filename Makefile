.PHONY: reset sonarscanner build push

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

profile_service_version = v0.1.1
project_service_version = v0.1.1
assessment_service_version = v0.1.0
scheduler_worker_version = v0.1.0
gateway_version = v0.1.1
api_gateway_version = v0.1.1

build: 
	docker build -t gcr.io/${PROJECT_ID}/profile-service:$(profile_service_version) -f ./services/profile/Dockerfile.prod .
	docker build -t gcr.io/${PROJECT_ID}/project-service:$(project_service_version) -f ./services/project/Dockerfile.prod .
	docker build -t gcr.io/${PROJECT_ID}/assessment-service:$(assessment_service_version) -f ./services/assessment/Dockerfile.prod .
	docker build -t gcr.io/${PROJECT_ID}/scheduler-worker:$(scheduler_worker_version) -f ./scheduler/worker/Dockerfile .
	docker build -t gcr.io/${PROJECT_ID}/gateway:$(gateway_version) -f ./gateway/Dockerfile .
	docker build -t gcr.io/${PROJECT_ID}/api-gateway:$(api_gateway_version) -f ./gateway/krakend/Dockerfile .

push:
	docker push gcr.io/${PROJECT_ID}/profile-service:$(profile_service_version)
	docker push gcr.io/${PROJECT_ID}/project-service:$(project_service_version)
	docker push gcr.io/${PROJECT_ID}/assessment-service:$(assessment_service_version)
	docker push gcr.io/${PROJECT_ID}/scheduler-worker:$(scheduler_worker_version)
	docker push gcr.io/${PROJECT_ID}/gateway:$(gateway_version)
	docker push gcr.io/${PROJECT_ID}/api-gateway:$(api_gateway_version)