APP_NAME=camermap-server
PORT=8080

.PHONY: docker-build app-dev app

docker-build:
	docker build -t $(APP_NAME) .

app-dev:    docker-build
	@echo "Running Server.."
	@if [ -z "$$GOOGLE_MAPS_API" ]; then \
		read -p "Enter your Google maps api key: " key; \
		GOOGLE_MAPS_API=$$key; \
	fi; \
	docker run --rm -e GOOGLE_MAPS_API="$$GOOGLE_MAPS_API" -p $(PORT):$(PORT) -v $(abspath ${CURDIR}/frontend):/frontend $(APP_NAME)


app:    docker-build
	@echo "Running Server.."
	@if [ -z "$$GOOGLE_MAPS_API" ]; then \
		read -p "Enter your Google maps api key: " key; \
		GOOGLE_MAPS_API=$$key; \
	fi; \
	docker run --rm -e GOOGLE_MAPS_API="$$GOOGLE_MAPS_API" -p $(PORT):$(PORT) $(APP_NAME)
