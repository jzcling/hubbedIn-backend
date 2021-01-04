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

profile_service_version = v0.1.2
project_service_version = v0.1.1
assessment_service_version = v0.1.0
joblisting_service_version = v0.1.0
scheduler_worker_version = v0.1.0
gateway_version = v0.1.2
api_gateway_version = v0.1.2

build: 
	docker build -t gcr.io/${PROJECT_ID}/profile-service:$(profile_service_version) -f ./services/profile/Dockerfile.prod .
	docker build -t gcr.io/${PROJECT_ID}/project-service:$(project_service_version) -f ./services/project/Dockerfile.prod .
	docker build -t gcr.io/${PROJECT_ID}/assessment-service:$(assessment_service_version) -f ./services/assessment/Dockerfile.prod .
	docker build -t gcr.io/${PROJECT_ID}/joblisting-service:$(joblisting_service_version) -f ./services/joblisting/Dockerfile.prod .
	docker build -t gcr.io/${PROJECT_ID}/scheduler-worker:$(scheduler_worker_version) -f ./scheduler/worker/Dockerfile .
	docker build -t gcr.io/${PROJECT_ID}/gateway:$(gateway_version) -f ./gateway/Dockerfile .
	docker build -t gcr.io/${PROJECT_ID}/api-gateway:$(api_gateway_version) -f ./gateway/krakend/Dockerfile .

push:
	docker push gcr.io/${PROJECT_ID}/profile-service:$(profile_service_version)
	docker push gcr.io/${PROJECT_ID}/project-service:$(project_service_version)
	docker push gcr.io/${PROJECT_ID}/assessment-service:$(assessment_service_version)
	docker push gcr.io/${PROJECT_ID}/joblisting-service:$(joblisting_service_version)
	docker push gcr.io/${PROJECT_ID}/scheduler-worker:$(scheduler_worker_version)
	docker push gcr.io/${PROJECT_ID}/gateway:$(gateway_version)
	docker push gcr.io/${PROJECT_ID}/api-gateway:$(api_gateway_version)

update:
	kubectl set image deployment profile-service profile-service-sha256-1=profile-service:$(profile_service_version)
	kubectl set image deployment project-service project-service-sha256-1=project-service:$(project_service_version)
	kubectl set image deployment assessment-service assessment-service-sha256-1=assessment-service:$(assessment_service_version)
	kubectl set image deployment joblisting-service joblisting-service-sha256-1=joblisting-service:$(joblisting_service_version)
	kubectl set image deployment scheduler-worker scheduler-worker-sha256-1=scheduler-worker:$(scheduler_worker_version)
	kubectl set image deployment gateway gateway-sha256-1=gateway:$(gateway_version)
	kubectl set image deployment api-gateway api-gateway-sha256-1=api-gateway:$(api_gateway_version)

delete-evicted:
	kubectl get pods --all-namespaces | grep Evicted | awk '{print $2 " --namespace=" $1}' | xargs kubectl delete pod
