# Build the Pocketbase binary
build:
	go build -o pocketbase cmd/main/main.go

# Format/prettify code
format:
	gofmt -w *.go

# Serve compiled Pocketbase client
serve:
	./pocketbase serve

# Serve source/dev Pocketbase client
serve-dev:
	go run cmd/main/main.go serve
