APP_NAME=camermap-server
PORT=8080

.PHONY: docker-build app example

docker-build:
	docker build -t $(APP_NAME) .	


app:
	@echo "Running Server.."
	@if [ -z "$$GOOGLE_MAPS_API" ]; then \
		read -p "Enter your Google maps api key: " key; \
		GOOGLE_MAPS_API=$$key; \
	fi; \
	docker run --rm -e GOOGLE_MAPS_API="$$GOOGLE_MAPS_API" -p $(PORT):$(PORT) $(APP_NAME)
