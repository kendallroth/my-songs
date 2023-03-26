################################################################################
## API
################################################################################

# Build the Pocketbase binary
api-build:
	cd ./pocketbase && go build -o pocketbase main.go && cd -

# Serve compiled Pocketbase client
api-serve:
	./pocketbase/main.go serve

# Serve source/dev Pocketbase client
api-serve-dev:
	cd ./pocketbase && go run main.go serve && cd -

################################################################################
## App
################################################################################

app-serve:
	cd ./web && npm run dev && cd -