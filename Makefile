################################################################################
## API
################################################################################

# Build the Pocketbase binary
api-build:
	cd ./pocketbase && Make api-build && cd -

# Serve compiled Pocketbase client
api-serve:
	cd ./pocketbase && Make api-serve && cd -

# Serve source/dev Pocketbase client
api-serve-dev:
	cd ./pocketbase && Make api-serve-dev && cd -

################################################################################
## App
################################################################################

app-serve:
	cd ./web && npm run dev && cd -