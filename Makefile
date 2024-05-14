# Binary name for the build output
BINARY_NAME=snake

run:
	@echo "Running application.."
	go run .

test:
	@echo "Running tests.."
	go test ./... -v
