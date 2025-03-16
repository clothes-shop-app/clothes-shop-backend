.PHONY: run build clean

run:
	go run ./cmd/app

build:
	go build -o ./bin/app ./cmd/app

clean:
	rm -rf ./bin/app

migration:
	@if [ -z "$(name)" ]; then \
		echo "need to set migration name: make migration name=create_users_table"; \
		exit 1; \
	fi
	migrate create -ext sql -dir migrations -seq $(name)
	@echo "migration crated"
