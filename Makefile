APP_NAME=job-queue-go

build:
	go build -o $(APP_NAME) main.go

run:
	go run main.go $(ARGS)

shell : 
	@echo "Starting the CLI"
	@while true; do \
		read -p "jobqueuecli >" LINE; \
		if ["$$CMD" = "exit"]; then break; fi; \
		eval "go run main.go $$LINE"; \
	done

test:
	go test ./... -v

clean:
	rm -f $(APP_NAME)

.PHONY: build run test clean
