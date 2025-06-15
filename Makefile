APP_NAME=job-queue-go

build:
	go build -o $(APP_NAME) main.go

run:
	go run main.go

test:
	go test ./... -v

clean:
	rm -f $(APP_NAME)

.PHONY: build run test clean
